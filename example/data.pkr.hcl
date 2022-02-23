packer {
  required_plugins {
    rhel-image = {
      version = ">= 0.1.0"
      source = "github.com/ygalblum/rhel-image"
    }
  }
}

data "rhel-image" "image" {
  offline_token = "${var.offline_token}"
  image_checksum = "${var.image_checksum}"
}