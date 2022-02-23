# Packer RHEL Image plugin

Packer plugin used to download RHEL images from access.redhat.com

<!-- For the full list of available features for this plugin see [documentation](https://www.packer.io/docs/datasources/sshkey). -->

Packer 1.7.3 or later is required.

## Usage example
```hcl
packer {
  required_plugins {
    rhel-image = {
      version = ">= 0.1.0"
      source = "github.com/ygalblum/rhel-image"
    }
  }
}

variable "image_checksum" {
  type = string
  default = "<Checksum of the image to download>"
}

data "rhel-image" "install" {
  offline_token = < The Generate Offline token for access.redhat.com >
  image_checksum = "${var.image_checksum}"
}

source "qemu" "install" {
  iso_url = data.rhel-image.install.image_path
  iso_checksum = "sha256:${var.image_checksum}"
  <...>
}

build {
  sources = ["source.qemu.install"]
}
```

Starting from version 1.7, Packer supports a new `packer init` command allowing
automatic installation of Packer plugins. Read the
[Packer documentation](https://www.packer.io/docs/commands/init) for more information.