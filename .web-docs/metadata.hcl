# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

# For full specification on the configuration of this file visit:
# https://github.com/hashicorp/integration-template#metadata-configuration
integration {
  name = "Triton"
  description = "The Triton Packer builder is able to create new images for use with Triton."
  identifier = "packer/hashicorp/triton"
  flags = ["archived"]
  component {
    type = "builder"
    name = "Triton"
    slug = "triton"
  }
}
