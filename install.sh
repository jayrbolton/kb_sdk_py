#!/bin/sh

# Prompt for sudo
# Download the binary from github to /usr/local/bin

{
  set -e
  SUDO=''
  if [ "$(id -u)" != "0" ]; then
    SUDO='sudo'
    echo "This script requires superuser access."
    echo "You will be prompted for your password by sudo."
    # clear any previous sudo permission
    sudo -k
  fi

  # run inside sudo
  $SUDO bash <<SCRIPT
    set -e
    echoerr() { echo "\$@" 1>&2; }
    if [[ ! ":\$PATH:" == *":/usr/local/bin:"* ]]; then
      echoerr "Your path is missing /usr/local/bin, you need to add this to use this installer."
      exit 1
    fi

    # Detect the OS
    if [ "\$(uname)" == "Darwin" ]; then
      OS=darwin
    elif [ "\$(expr substr \$(uname -s) 1 5)" == "Linux" ]; then
      OS=linux
    else
      echoerr "This installer is only supported on Linux and MacOS"
      exit 1
    fi

    # Detect the architecture
    ARCH="\$(uname -m)"
    if [ "\$ARCH" == "x86_64" ]; then
      ARCH=amd64
    elif [[ "\$ARCH" == arm* ]]; then
      ARCH=arm
    elif [[ "\$ARCH" == i386 ]]; then
      ARCH=386
    elif [[ "\$ARCH" == amd64 ]]; then
      ARCH=amd64
    else
      echoerr "Unsupported arch: \$ARCH"
      exit 1
    fi

    VERSION=v0.0.1
    URL=https://github.com/jayrbolton/kbase_sdk_cli/releases/download/\$VERSION/cli_\$OS_\$ARCH

    cd /usr/local/bin
    rm -rf kbase-sdk
    curl -L --output kbase-sdk "\$URL"
    chmod +x kbase-sdk

SCRIPT

  # test the CLI
  LOCATION=$(command -v kbase-sdk)
  echo "KBase SDK CLI installed to $LOCATION"
  kbase-sdk version
}
