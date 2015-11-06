#!/bin/bash

function DVM () {

  function _PRERUN () {
    if [[ ! -e "${DRUSHDIR}/drush" ]] && [[ ! -d "${DRUSHDIR}" ]]; then
      mkdir -p "${DRUSHDIR}";
      cd "${DRUSHDIR}";
      curl -sS https://getcomposer.org/installer | php;
      php composer.phar require drush/drush;
    fi
  }

  function _GETVERSIONS () {
    if [[ -z "${CLEANARG}" ]]; then
      if [[ -n $(cd "${DRUSHDIR}" && composer show drush/drush | grep versions | grep "${CLEANARG}") ]]; then VALIDVERSION=true; fi
    else
      if [[ -n $(cd "${DRUSHDIR}" && composer show drush/drush | grep versions) ]]; then VALIDVERSION=true; fi
    fi
  }
  function _GETCOMPOSER () {
    if [[ -z $(which composer) ]]; then
      curl -sS https://getcomposer.org/installer | php
      mv composer.phar /usr/local/bin/composer
    fi
  }
  function _SETVARS () {
    DRUSHDIR="${HOME}/drush";
    if [[ "${*}" == *"reinstall"* ]]; then STATE="reinstall"; fi
    if [[ "${*}" == *"uninstall"* ]]; then STATE="uninstall"; fi
    CLEANARG="${*}";
    CLEANARG=${CLEANARG/uninstall/};
    CLEANARG=${CLEANARG/reinstall/};
    CLEANARG=${CLEANARG/install/};
    CLEANARG=${CLEANARG/update/};
    CLEANARG=${CLEANARG/use/};
    CLEANARG=${CLEANARG/ls-local/};
    CLEANARG=${CLEANARG/ls-remote/};
    CLEANARG=${CLEANARG/ /};

    case "${CLEANARG}" in
      5*)  VERSION="${CLEANARG}"; ;;
      6*)  VERSION="${CLEANARG}"; ;;
      7*)  VERSION="${CLEANARG}"; ;;
      8*)  VERSION="${CLEANARG}"; ;;
    esac

    BASEVERSION=${VERSION::1};
    _GETVERSIONS;

    if [[ "${VERSION}" == *"*"* ]]; then
      DRUSHVERDIR="${DRUSHDIR}/versions/drush-${BASEVERSION}-master";
    else
      DRUSHVERDIR="${DRUSHDIR}/versions/drush-${VERSION}";
    fi

  }

  function _SWITCH () {
    sudo rm -f "/usr/local/bin/drush";
    sudo ln -s "${DRUSHVERDIR}/vendor/bin/drush" "/usr/local/bin/drush";
    INSTALLED=$(drush --version | cut -d: -f2);
    INSTALLED=${INSTALLED/  /};
    INSTALLED=${INSTALLED/ /};
    echo "Drush is now using ${INSTALLED}";
  }

  function _INSTALL () {
    # for X in "${VERSIONSAVAILABLE[@]}"; do
    #   echo "X: ${X}";
      # echo "VERSION: ${VERSION}";
      if [[ "${STATE}" == "uninstall" ]]; then
        INSTALLED=$(drush --version | cut -d: -f2);
        INSTALLED=${INSTALLED/  /};
        INSTALLED=${INSTALLED/ /};
        TOUNINSTALL=$(_FETCH_LOCAL);
        TOUNINSTALL=${TOUNINSTALL/  /};
        TOUNINSTALL=${TOUNINSTALL/ /};
        if [[ ! "${TOUNINSTALL}" == *"${INSTALLED}"* ]]; then
          rm -rf "${DRUSHDIR}/versions/drush-${CLEANARG}";
          echo "Folder has been removed."
        else
          echo "You cannot uninstall a currently installed version."
        fi
      fi
      if [[ "${VALIDVERSION}" == true ]]; then
        if [[ ! -d "${DRUSHVERDIR}" ]] && [[ "${STATE}" == "install" ]]; then
          echo "Drush v${CLEANARG} has already been installed.";
        else
          if [[ "${BASEVERSION}" == "5" ]]; then
            mkdir -p "${DRUSHDIR}/versions/${X}";
            cd "${DRUSHDIR}/versions/";
            wget "https://github.com/drush-ops/drush/archive/${CLEANARG}.zip"
            unzip "${CLEANARG}.zip"
            rm -f "${CLEANARG}.zip";
            sudo mv -f "./drush-${CLEANARG}" ./drush5
            if [[ -d "${CLEANARG}" ]]; then
              rm -rf "${CLEANARG}";
            fi
          fi
          if [[ "${BASEVERSION}" == "6" ]] || [[ "${BASEVERSION}" == "7" ]] || [[ "${BASEVERSION}" == "8" ]]; then
            _GETCOMPOSER;
            mkdir -p "${DRUSHVERDIR}";
            cd "${DRUSHVERDIR}";
            curl -sS https://getcomposer.org/installer | php;
            php composer.phar require "drush/drush:${VERSION}";
          fi
          _SWITCH;
        fi
      else
        echo "Drush version specified does not exist.";
      fi
    # done
  }

  function _FETCH_REMOTE () {
    declare -a AVAILABLEVERSIONS=($(composer show drush/drush | grep versions | tr ", " "\n"));
    if [[ -z "${CLEANARG}" ]]; then
      for X in "${AVAILABLEVERSIONS[@]}"; do echo "      ${X}" | grep "6."; done
      for X in "${AVAILABLEVERSIONS[@]}"; do echo "      ${X}" | grep "7."; done
      for X in "${AVAILABLEVERSIONS[@]}"; do echo "      ${X}" | grep "8."; done
    else
      for X in "${AVAILABLEVERSIONS[@]}"; do echo "      ${X}" | grep "${CLEANARG}"; done
    fi
  }

  function _FETCH_LOCAL () {
    INSTALLED=$(drush --version | cut -d: -f2);
    INSTALLED=${INSTALLED/  /};
    INSTALLED=${INSTALLED/ /};
    declare -a AVAILABLEVERSIONS=();
    AVAILABLEVERSIONS=($(cd ${DRUSHDIR}/versions && ls | cut -d/ -f2));
    for X in "${AVAILABLEVERSIONS[@]}"; do
      if [[ "${X}" == "drush-${INSTALLED}" ]]; then tput setaf 2; fi
      if [[ "${X}" = *"drush-${BASEVERSION}"* ]]; then
        if [[ -z "${CLEANARG}" ]]; then
          echo "      ${X}";
        else
          echo "      ${X}" | grep ${CLEANARG};
        fi
      fi
      if [[ "${X}" == "drush-${INSTALLED}" ]]; then tput sgr0; fi
    done
  }

  function _HELP () {
    echo " Drush Version Manager";
    echo " ";
    echo " Note: <version> refers to any version-like string dvm understands. This includes:";
    echo "   - full version numbers (7.0.0, 7.1.0, 8.0.0-rc3)";
    echo " ";
    echo " Usage:";
    echo "   dvm help                                  Show this message.";
    echo "   dvm install <version>                     Download and install <version>.";
    echo "   dvm reinstall <version>                   Reinstall a version.";
    echo "   dvm uninstall <version>                   Uninstall a version.";
    echo "   dvm use <version>                         Change the drush symlink to use another version.";
    echo "   dvm current                               Display currently activated version.";
    echo "   dvm status                                Display current status.";
    echo "   dvm ls-local                              List installed versions.";
    echo "   dvm ls-local <version>                    List installed versions matching a search criteria";
    echo "   dvm ls-remote                             List remote versions available for install";
    echo "   dvm ls-remote <version>                   List remote versions matching a search criteria";
    echo " ";
    echo " Example:";
    echo "   dvm install 7.0.0                         Install a specific version number";
    echo "   dvm use 7.0.0                             Tells the system to use drush 7.0.0";
  }

  if [[ "${*}" == *"help"* ]] || [[ -z "${*}" ]]; then
    _HELP;
    exit 0;
  fi

  _SETVARS "${*}";
  _PRERUN;
  case "${*}" in
    *"current"*)       drush --version; ;;
    *"status"*)        drush status; ;;
    *"install"*)       _INSTALL; ;;
    *"use"*)           _SWITCH; ;;
    *"ls-remote"*)     _FETCH_REMOTE; ;;
    *"ls-local"*)      _FETCH_LOCAL; ;;
    *"initialize"*) ;;
  esac
}

DVM "${*}";
