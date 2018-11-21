# KBase SDK 2 - Command Line Interface

This is a command line interface for writing, managing, and testing KBase modules.

Also see the [kbase_module]() Python package, which is used inside the modules themselves.

## Install

Install via a quick shell script:

```sh
curl https://github.com/kbase/kbase_sdk_cli/archive/install.sh | sh
```

[Read the above shell script to see what it does](). The steps it takes are:
- Prompts for you to choose an installation directory (defaults to ~/.kbase)
- Copies the CLI binary and sets up some storage directories
- Prompts you to add the binary to your $PATH

If you installed to `~/.kbase`, then add `~/.kbase/bin/kbase-sdk` to your $PATH.

## Setup

Set the following environment variables:

* `KBASE_USERNAME` - required - your KBase developer username
* `KBASE_DEV_TOKEN` - required - your KBase developer token
* `KBASE_CLI_PATH` - optional - the installation directory of the CLI (defaults to `~/.kbase`)

## Usage

```sh
$ kbase-sdk help
```

You can also use the syntax `kbase-sdk -h` or `kbase-sdk --help`.

For any of the commands, you can run

```sh
kbase-sdk <command> --help
```

to get quick usage details (equivalently `kbase-sdk <command> -h`.

To get more detailed information about a command, run

```sh
kbase-sdk help <command>
```

### Initialize a module

```sh
$ kbase-sdk init project_name [directory]
```

Directory is optional, and defaults to the project name.

### Validate a module

Inside a module's directory:

```sh
$ kbase-sdk validate
```

### Run integration tests

```sh
$ kbase-sdk test
```

### Publish or view registration status

View the registration status of your module:

```sh
$ kbase-sdk status
```

Publish the current version of your module to a kbase catalog:

```sh
$ kbase-sdk publish
```

### Upgrade the CLI

Check for any updates on the SDK

```sh
$ kbase-sdk upgrade
```

## Development

This section has information about development on the CLI itself (not on KBase modules).
