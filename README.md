# Drush Version Manager 2.x

**Purpose**:

Drush version control system, to manage a variety of Drush versions and provide the ability to switch to specific versions at any time.

It has been rewritten in Google's Golang as a pilot for the core maintainer to learn a new language and see what it would hold in professional development. It wasn't intended to stick or be released however the pilot program was extremely stable and efficient and it didn't make sense not to release it.

**DVM usage**

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

````

The flag system is undergoing active development, it's far from final.

Examples:

**Install**

dvm -version=7.2.0 -install

**Uninstall**

dvm -version=7.2.0 -uninstall

**Reinstall**

dvm -version=7.2.0 -uninstall

**Reinstall**

dvm -version=7.2.0 -default

**List available versions**

dvm -list=available

**List installed versions**

dvm -list=available

**List installed drush modules** - Needs work!

dvm -list=installed -package="reg"

````

**Usage in Go programming**:

Examples of each module can be found in the test functions of each test file of each package as tests have been written and passed for every function written.

**Package rundown**:

* Drush Version Management
    * plugin
        ````
        Drush package management, provides a way to install, reinstall, uninstall and list package objects.
        ````
    * version
        ````
        Drush version management, where a version can be installed, reinstalled, uninstalled and a variety of other utility methods.
        ````
    * versionlist
        ````
        Provides Drush Version Management information, by displaying desired information on particular versions - installed or not.
        ````
* Drush Execution
    * aliases
        ````
        Execute 'drush sa' and grab all aliases, and provides convenient ways to filter and store those results.
        ````
    * sites
        ````
        Execute 'drush sa' and grab all aliases and filter them as non-alias items, with a wide array of filter options.
        ````