# KBase SDK 2 - Command Line Interface

This is a command line interface for writing, managing, and testing KBase modules.

Also see the [kbase_module](https://github.com/jayrbolton/kbase_module) package, which is imported inside the actual modules.

## Install

Install via a quick shell command:

```sh
curl -L https://github.com/jayrbolton/kbase_sdk_cli/releases/download/v0.0.1/install.sh | sh
```

The above script downloads a binary from the Github releases page to `/usr/local/bin/kbase-sdk`

You can install the CLI manually by going to the releases page and selecting the binary for your OS and architecture. Download the file to a directory in your executable `$PATH`, such as `/usr/local/bin` or `~/.local/bin`.

## Setup

Set the following environment variables:

* `KBASE_USERNAME` - required - your KBase developer username
* `KBASE_DEV_TOKEN` - optional - your KBase developer token
* `KBASE_CLI_PATH` - optional - the installation directory of the CLI (defaults to `~/.kbase`)

## Usage

```sh
$ kbase-sdk help
```

You can also use the syntax `kbase-sdk -h` or `kbase-sdk --help`.

To get more detailed information about a command, run any of

```sh
kbase-sdk help <command>
kbase-sdk <command> --help
kbase-sdk <command> -h
```

### Initialize a module

```sh
kbase-sdk init [module name]
```

### Run tests

In a module's directory, run

```sh
kbase-sdk test
```

On first run, your Docker container will be built. On each subsequent run, it will use the previously built container. To force a new build, do

```sh
kbase-sdk test --build
```

You can also build without using any caching with `kbase-sdk test --build-no-cache`.

### Upgrade the CLI

Check for any updates on the SDK with

```sh
kbase-sdk upgrade
```

## Development

This section has information about development on the CLI itself (not on KBase modules).

This project uses Go, which can be installed [with these instructions](https://golang.org/doc/install). Clone this project under `src/kbase/kbase_sdk_cli` [inside your Go workspace](https://golang.org/doc/code.html).

One way to install dependencies is to run `go get ./...` while inside the project directory.

### Publishing updates

To publish new CLI code, first run:

```sh
bash build.sh
```

This will generate binaries in the `dist/` folder. Then, create a new Github release on the `kbase/kbase_sdk_cli` repository. Upload all the binaries along with the `install.sh` script.

Users of the CLI can automatically download your new release by running `kbase-sdk upgrade`.
