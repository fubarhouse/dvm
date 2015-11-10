# Drush Version Manager

DVM was inspired from NVM, RVM and PyEnv in that where so many versions of Drush and there's no convenient way to simply switch version in the way NVM, RVM and PyEnv have endeavored to achieve. With DVM joining the pack, the release of Drupal 8 and the popularity of a server configuration management tool Ansible I've essentially made this happen where there's no alternative program publicly available.

It's usage was intended for myself, but with a tool this convenient and cool there was no way I could keep it to myself.

## Installation

This script uses UNIX commands via Bash for ultimate compatibility and such requires a UNIX-based system.
It was made, tested and used on a OSX El Capitan system for intended use on a Debian environment for use with Vagrant.

### Ansible role

  An Ansible role has been made available for ease of implementation of Debian based systems: https://github.com/fubarhouse/fubarhouse.dvm

  This role can be found on the Galaxy at https://galaxy.ansible.com/detail#/role/5868

  This role was produced to work along-side [Jeff Geerling](https://twitter.com/geerlingguy)s' [DrupalVM](http://www.drupalvm.com/) and it works well for that purpose.

### Install script

To install or update dvm, you can use Wget:

    wget -O /usr/local/bin/dvm https://raw.githubusercontent.com/fubarhouse/dvm/master/dvm

And then you can ensure all the dependencies (composer, unzip etc) are installed using:

    dvm initialize

Dependencies include a default Drush installation (installed at ~/drush/) using composer, but it won't take long.

Dependencies not installed via the initialization argument (currently) include:

* pear
* wget
* unzip

### Manual install

For manual install create a folder somewhere in your filesystem with the `dvm` file inside it.

Once you have this file, move it somewhere available to the `$PATH` variable.

Run `dvm initialize` to get started.

## Configuring

To escape the dependencies of the sudo, there's one variable you may need to configure at the beginning of the `_SETVARS` function. Change the value to a directory which the default account has full read, write and executable access, otherwise the script may fail when changing to use the default versions. DVM can be used with sudo correctly even if this isn't set up properly.

    LINKDIR="/usr/local/bin";

## Usage

To get a list of available versions:

    dvm ls-remote

Or to search for a version, type an argument to ls-remote to do a basic query string on the same set of commands:

    dvm ls-remote 8.0.0

To download and install version v7.0.0 and v8.0.0-rc3, do this:

    dvm install 7.0.0
    dvm install 8.0.0-rc3

You can switch between installed versions:

    dvm use 7.0.0
    dvm use 8.0.0-rc3

Or you can just run a command using a specific version of Drush using:

    dvm exec 7.0.0 --version

## Compatibility

### Drush version Manager

Only OSX and Ubuntu (CLI) systems are supported at this time.

### Drush

So far, Drush versions `2.0.0` and later are successfully working without errors.

There's a desire to make Drush v1 work, but there's no practical reasoning to it - so this may come later.

For use with Drush v1,2,3,4 & 5, the script will use wget based on an array with all the information and all other versions will use composer to get the live versions available.

## License

dvm is released under the MIT license.

Copyright (C) 2015-2020 Karl Hepworth

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

## Compatibility Issues

`dvm` was made for use by OSX, Ubuntu and UNIX systems, and supports no other platform at this time.

## Problems

There are no known errors.

Still early days for this one, but there're questions about what we could do with this micro-app so time will tell.
