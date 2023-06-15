The Triton Packer builder is able to create new images for use with Triton.
These images can be used with both the [Joyent public cloud](https://www.joyent.com/) (which is powered by Triton) as well as with private [Triton](https://github.com/joyent/triton) installations.

This builder uses the Triton Cloud API to create these images. Triton also
supports the Docker API however this builder does not. If you want to create
Docker images on Triton you should use the Packer Docker builder.

### Installation
To install this plugin add this code into your Packer configuration and run [packer init](/packer/docs/commands/init)

```hcl
packer {
  required_plugins {
    triton = {
      version = "~> 1"
      source  = "github.com/hashicorp/triton"
    }
  }
}
```

Alternatively, you can use `packer plugins install` to manage installation of this plugin.

```sh
packer plugins install github.com/hashicorp/triton
```

### Components

#### Builders

- [triton](/packer/integrations/hashicorp/triton/latest/components/builder/triton) - The Triton Packer builder is
  able to create new images for use with Triton.
