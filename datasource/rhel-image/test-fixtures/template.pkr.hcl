data "rhel-image" "test" {
  offline_token = "offlineToken"
  image_checksum = "12345"
}

locals {
  image_path = data.rhel-image.test.image_path
}

source "null" "basic-example" {
  communicator = "none"
}

build {
  sources = [
    "source.null.basic-example"
  ]

  provisioner "shell-local" {
    inline = [
      "echo image_path: ${local.image_path}",
    ]
  }
}
