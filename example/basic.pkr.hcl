source "triton" "example" {
  image_name    = "my_new_image"
  image_version = "1.0.0"
  source_machine_image_filter = {
    most_recent = "true"
    name        = "ubuntu-16.04"
    type        = "lx-dataset"
  }
  source_machine_name    = "image-builder"
  source_machine_package = "g4-highcpu-128M"
  ssh_username           = "root"
  triton_account         = "triton_username"
  triton_key_id          = "6b:95:03:3d:d3:6e:52:69:01:96:1a:46:4a:8d:c1:7e"
}

build {
  sources = ["source.triton.example"]
}
