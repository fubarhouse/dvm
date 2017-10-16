# Drush Version Manager 2.x

## Purpose

Drush version control system, to manage a variety of Drush versions and provide the ability to switch to specific versions at any time.

This project was inspired by `NVM`, and was originally written in Bash. It was rewritten in Go in late-2016 and has been improving since.

It came about when attempting to use multiple Drush versions with the use of CI tooling. It has become more important now as Drupal has essentially abandoned the use of `drush make`.

## Requirements

Drush version manager requires [composer](https://getcomposer.org/), and nothing more.
* composer

## Installation

There are three ways to install DVM, but if you're at all familiar with Golang these are very standardised shipping methods.

**Option 1**: Like any other Go binary
1. run `go get -u github.com/fubarhouse/dvm`
2. Use like any other Go binary.

**Option 2**: Download a precompiled binary!
1. Download one of the [pre-compiled binaries](https://github.com/fubarhouse/dvm/releases).
2. Copy to location in `$PATH` environment variable.

**Option 3**: - As an API for use in your Go project
1. run `go get -u github.com/fubarhouse/dvm`
2. Build your own Go binary using the API's from the packages downloaded.

## Configuration
Configurations are loaded via [Viper](https://github.com/spf13/viper), an example is below.

The default values are in this example and should be overriden in `~/.dvm/config.toml` to suit your system.

```
[config]
path = "/usr/local/bin/drush"
```

## Usage

### General usage

````
Usage:
  dvm [command] [flags]

Available Commands:
  help        Help about any command
  install     Install a specific version of Drush
  list        List available or installed Drush versions.
  reinstall   Reinstall a specific version of Drush
  uninstall   Uninstall a specific version of Drush
  use         Initialise or replace an established symlink to the configured location, for a given version of Drush

Flags:
  --config string      config file (default is $HOME/config.toml)
  -a, --available      List available versions
  -h, --help           help for dvm
  -i, --installed      List installed versions
  -v, --version string Version to target, it does not have a default value.

````

### Examples

* **Install**:  `dvm install --version 7.2.0`
* **Uninstall**: `dvm uninstall --version 7.2.0`
* **Reinstall**: `dvm reinstall --version 7.2.0`
* **Switch**: `dvm use --version ````7.2.0`
* **List Available**: `dvm list --available`
* **List Installed**: `dvm list --installed`