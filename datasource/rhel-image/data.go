//go:generate packer-sdc mapstructure-to-hcl2 -type Config,DatasourceOutput
package rhelimage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/hcl2helper"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"
)

const (
	refreshTokenUrl = "https://sso.redhat.com/auth/realms/redhat-external/protocol/openid-connect/token"
	imageDetailsUrlFormat = "https://api.access.redhat.com/management/v1/images/%v/download"
	defaultTargetDirectory = "/tmp"
)

type Config struct {
	OfflineToken string `mapstructure:"offline_token"`
	ImageChecksum string `mapstructure:"image_checksum"`
	TargetDirectory string `mapstructure:"target_directory"`
}

type Datasource struct {
	config Config
}

type DatasourceOutput struct {
	ImagePath string `mapstructure:"image_path"`
}

func (d *Datasource) ConfigSpec() hcldec.ObjectSpec {
	return d.config.FlatMapstructure().HCL2Spec()
}

func (d *Datasource) Configure(raws ...interface{}) error {
	err := config.Decode(&d.config, nil, raws...)
	if err != nil {
		return err
	}
	return nil
}

func (d *Datasource) OutputSpec() hcldec.ObjectSpec {
	return (&DatasourceOutput{}).FlatMapstructure().HCL2Spec()
}

func getAccessToken(refreshToken string) (*string, error) {
	resp, err := http.PostForm(
		refreshTokenUrl,
		url.Values{"client_id": {"rhsm-api"}, "grant_type": {"refresh_token"}, "refresh_token": {refreshToken}},
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to refresh the token. bad status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return nil, err
	}

	ret := fmt.Sprintf("%v", data["access_token"])
	return &ret, nil

}

func getImageDetails(accessToken, checksum string) (*string, *string, error) {
	address := fmt.Sprintf(imageDetailsUrlFormat, checksum)
	req, err := http.NewRequest("get", address, http.NoBody)
	if err != nil {
		return nil, nil, err
	}

	bearer := "Bearer " + accessToken
	req.Header.Add("Authorization", bearer)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusTemporaryRedirect {
		return nil, nil, fmt.Errorf("failed to get image details. bad status: %s", resp.Status)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal([]byte(responseBody), &data)
	if err != nil {
		return nil, nil, err
	}

	body, ok := data["body"].(map[string]interface{})
	if !ok {
		return nil, nil, errors.New("body field is not a map")
	}

	imageUrl := fmt.Sprintf("%v", body["href"])
	imageFileName := fmt.Sprintf("%v", body["filename"])
	return &imageUrl, &imageFileName, nil
}

func downloadImage(targetPath, address, accessToken string) error {
	log.Printf("Downloading from address [%v] into path [%v]\n", address, targetPath)

	out, err := os.Create(targetPath)
	if err != nil  {
	  return err
	}
	defer out.Close()

	req, err := http.NewRequest("get", address, http.NoBody)
	if err != nil {
		return err
	}

	bearer := "Bearer " + accessToken
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
	  return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil  {
	  return err
	}

	return nil
}

func (d *Datasource) Execute() (cty.Value, error) {
	accessToken, err := getAccessToken(d.config.OfflineToken)
	if err != nil {
		return cty.NullVal(cty.EmptyObject), err
	}

	imageUrl, imageFilename, err := getImageDetails(*accessToken, d.config.ImageChecksum)
	if err != nil {
		return cty.NullVal(cty.EmptyObject), err
	}

	targetDirectory := d.config.TargetDirectory
	if targetDirectory == "" {
		targetDirectory = defaultTargetDirectory
	}

	targetPath := targetDirectory + "/" + *imageFilename

	err = downloadImage(targetPath, *imageUrl, *accessToken)
	if err != nil {
		return cty.NullVal(cty.EmptyObject), err
	}

	output := DatasourceOutput{
		ImagePath: targetPath,
	}
	return hcl2helper.HCL2ValueFromConfig(output, d.OutputSpec()), nil
}
