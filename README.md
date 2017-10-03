# Drush Version Manager 2.x

## Purpose

Drush version control system, to manage a variety of Drush versions and provide the ability to switch to specific versions at any time.

It has been rewritten in Google's Golang as a pilot for the core maintainer to learn a new language and see what it would hold in professional development. It wasn't intended to stick or be released however the pilot program was extremely stable and efficient and it didn't make sense not to release it.

## Requirements

DVM will require the following to be available in the system depending on what you're trying to do:
* composer
* unzip
* wget

## Installation

There are three ways to install DVM, but if you're at all familiar with Golang these are very standardized shipping methods.

### Option 1
1. run `go get -u github.com/fubarhouse/dvm/...`
2. Use like any other Go binary.

### Option 2
1. run `go get -u github.com/fubarhouse/dvm/...`
2. Build your own Go binary using the API's from the packages downloaded.

### Option 3
1. Download the [pre-compiled binaries](https://github.com/fubarhouse/dvm/releases).
2. Copy to location in $PATH environment variable.

## Configuration
Configurations are loaded via Viper, an example is below.

The default values are in this example and should be overriden in `~/.dvm/config.toml` to suit your system.

```
[config]
path = "/usr/local/bin/drush"
```

## DVM Usage

### Install

`dvm install 7.2.0`

### Uninstall

`dvm uninstall 7.2.0`

### Reinstall

`dvm use 7.2.0`

### List Available Versions

`dvm list available`

### List Installed Versions

`dvm list installed`

### Install Drush Modules

`dvm package install registry_rebuild`

### List Installed Drush Modules (needs work)

`dvm package list`

## Usage in Go Programming

Examples of each module can be found in the test functions of each test file of each package as tests have been written and passed for every function written.

### Package Rundown

#### plugin
* Drush package management, provides a way to install, reinstall, uninstall and list package objects.
#### version
* Drush version management, where a version can be installed, reinstalled, uninstalled and a variety of other utility methods.
#### versionlist
* Provides Drush Version Management information, by displaying desired information on particular versions - installed or not.
