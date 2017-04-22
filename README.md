# Drush Version Manager 2.x

## Purpose

Drush version control system, to manage a variety of Drush versions and provide the ability to switch to specific versions at any time.

It has been rewritten in Google's Golang as a pilot for the core maintainer to learn a new language and see what it would hold in professional development. It wasn't intended to stick or be released however the pilot program was extremely stable and efficient and it didn't make sense not to release it.

## Installation

There are three ways to install DVM, but if you're at all familiar with Golang these are very standardized shipping methods.

1. run `go get -u github.com/fubarhouse/dvm/...`
2. run `go get -u github.com/fubarhouse/dvm/...` and compile your own build the program.
3. Download the pre-compiled binaries.

Either method chosen, you can optionally copy or create a symlink to the binary in a location discoverable by $PATH.

## DVM Usage

```text
  -default
    Version of Command set as system default
  -install
    Version of Command to install
  -list string
    List to print (installed|available)
  -package string
    Use package flag to tell DVM to target package.
  -reinstall
    Version of Command to reinstall
  -uninstall
    Version of Command to uninstall
  -version string
    Version to perform action on.
```

The flag system is undergoing active development, it's far from final.

## Examples

### Install

`dvm -version=7.2.0 -install`

### Uninstall

`dvm -version=7.2.0 -uninstall`

### Reinstall

`dvm -version=7.2.0 -default`

### List Available Versions

`dvm -list=available`

### List Installed Versions

`dvm -list=available`

### List Installed Drush Modules (needs work)

`dvm -list=installed -package="reg"`

## Usage in Go Programming

Examples of each module can be found in the test functions of each test file of each package as tests have been written and passed for every function written.

### Package Rundown

* plugin
  * Drush package management, provides a way to install, reinstall, uninstall and list package objects.
* version
  * Drush version management, where a version can be installed, reinstalled, uninstalled and a variety of other utility methods.
* versionlist
  * Provides Drush Version Management information, by displaying desired information on particular versions - installed or not.
