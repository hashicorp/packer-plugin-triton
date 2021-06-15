# Triton Plugins

The Triton Packer builder is able to create new images for use with Triton.
These images can be used with both the
[Joyent public cloud](https://www.joyent.com/) (which is powered by Triton) as
well as with private [Triton](https://github.com/joyent/triton) installations.

This builder uses the Triton Cloud API to create these images. Triton also
supports the Docker API however this builder does not. If you want to create
Docker images on Triton you should use the Packer Docker builder.

## Installation

### Using pre-built releases

#### Using the `packer init` command

Starting from version 1.7, Packer supports a new `packer init` command allowing
automatic installation of Packer plugins. Read the
[Packer documentation](https://www.packer.io/docs/commands/init) for more information.

To install this plugin, copy and paste this code into your Packer configuration .
Then, run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    triton = {
      version = ">= 1.0.0"
      source  = "github.com/hashicorp/triton"
    }
  }
}
```

#### Manual installation

You can find pre-built binary releases of the plugin [here](https://github.com/hashicorp/packer-plugin-triton/releases).
Once you have downloaded the latest archive corresponding to your target OS,
uncompress it to retrieve the plugin binary file corresponding to your platform.
To install the plugin, please follow the Packer documentation on
[installing a plugin](https://www.packer.io/docs/extending/plugins/#installing-plugins).


#### From Source

If you prefer to build the plugin from its source code, clone the GitHub
repository locally and run the command `go build` from the root
directory. Upon successful compilation, a `packer-plugin-triton` plugin
binary file can be found in the root directory.
To install the compiled plugin, please follow the official Packer documentation
on [installing a plugin](https://www.packer.io/docs/extending/plugins/#installing-plugins).


## Plugin Contents

The Triton plugin is intended as a starting point for creating Packer plugins, containing:

### Builders

- [builder](/docs/builders/builder-triton.mdx) - The Triton Packer builder is
  able to create new images for use with Triton.