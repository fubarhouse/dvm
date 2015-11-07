# Drush Version Manager

## Installation

This script uses UNIX commands and such requires a UNIX-based system.
It was made and tested on a OSX El Capitan system for a Debian environment for use with Vagrant

### Ansible role

  An Ansible role has been made available for ease of implementation of Debian based systems: https://github.com/fubarhouse/fubarhouse.dvm

  This role can be found on the Galaxy at https://galaxy.ansible.com/detail#/role/5868

### Install script

To install or update dvm, you can use Wget:

  wget -O /usr/local/bin/dvm https://raw.githubusercontent.com/fubarhouse/dvm/master/dvm

### Manual install

For manual install create a folder somewhere in your filesystem with the `dvm` file inside it.

Once you have this file, move it somewhere available to the `$PATH` variable.

## Usage

To download, install, and set the default version to v7.0.0 release of node, do this:

    dvm install 7.0.0

You can switch between installed versions:

    dvm use 8.0.0-rc3

Or you can just run a command using a specific version of drush using:

    dvm exec 8.0.0-rc3 --version

## Compatibility

### Drush version Manager

Only OSX and Ubuntu (CLI) systems are supported at this time.

### Drush

So far, Drush versions `3.3.0` and later are successfully working without errors.

For use with Drush v1,2,3,4 & 5, the script will use wget based on an array with all the information and all other versions will use composer to get the live versions available.

I am planning on building in compatibility for prior versions but there is no practical application for that feature.

## License

nvm is released under the MIT license.


Copyright (C) 2015-2020 Karl Hepworth

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

## Compatibility Issues

`dvm` was made for use by OSX, Ubuntu and UNIX systems, and supports no other platform at this time.

## Problems

There are no known errors.

Still early days for this one, but there're questions about what we could do with this micro-app so time will tell.
