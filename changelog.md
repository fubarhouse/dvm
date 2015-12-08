# Drush Version Manager

## Change log

### 1.4

    #48: Update change log
    #47: add documentation regarding campaign page
    #46: Update changelog and release 1.4
    #45: improve performance of composer
    #44: create campaign website
    #43: prevent use of fs commands when no version exists
    #42: add details on tested environments
    #41: Fix call to initialisation function
    #40: segregate versions into individual arrays
    #39: add 8.0.1 to master
    #38: functionalize new get_versions method and tell other functions to use this
    #37: update get_versions to latest version output method
    #36: reorder the output of "ls remote"
    #35: simplify update function
    #34: add regex to update function
    #33: address the accurateness of the current command
    #32: Add 8.0.0
    #30: revise readme for public image
    #29: encapsulate complete change history in changelog
    #28: reduce functions available to those who install with wget

### 1.3

    #31: sort latest version before updating
    #27: add a basic changelog
    #26: adjust output spacing
    #25: add ability to simply get and install the latest version
    #24: fix duplicated output on new ls function
    #23: fix up "current" command
    #22: remove regex *[.dev]* from installable versions
    #21: convert arguments to array for very strict validation
    #20: associate ls and ll to fetching
    #19: remove char 'v' from $CLEANARG output
    #18: explicit validation of arguments
    #17: use getopts
    #16: Change regex integer values to ranges.
    #15: add 8.0.0-rc4

### 1.2

    #14: add validation of symlink status in case dvm is not in use
    #13: dvm isn't changing versions on vagrant.
    #12: document upgrade path better
    #11: clean up case statement selectors with regex
    #9: symlinks are not being corrected
    #8: make git usage compatible for 1.7.0
    #4: add logic to prevent upgrades to non-git installations of dvm

### 1.1

    #5: expose upgrade as function with it's own unique logic, vars etc...
    #3: remove duplicates of arrays
    #2: proper upgrades to dvm via version numbers associated to git
    #1: only use composer when version not found

### 1.0

    #7: allow upgrade path for version 1.0
    #6: adjust outdated readme information
