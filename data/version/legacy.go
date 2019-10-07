package version

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"

	log "github.com/Sirupsen/logrus"
	"github.com/fubarhouse/dvm/commands/wget"
)

const sep = string(os.PathSeparator)

// LegacyInstall is basically the main() func for Legacy versions which encapsulates
// the code/dependencies for installing legacy Drush versions.
//
// Deprecated: Drush version manager no longer supports legacy installs.
func (drushVersion *DrushVersion) LegacyInstall() {
	drushVersion.LegacyInstallVersion()
	drushVersion.LegacyInstallTable()
}

// LegacyInstallTable is essentially always missing from older Drush versions.
// This ensures the script is available to the legacy version.
//
// Deprecated: Drush version manager no longer supports legacy installs.
func (drushVersion *DrushVersion) LegacyInstallTable() {
	// @TODO: Rewrite this
	usr, _ := user.Current()
	ctFileName := "Table.inc"
	ctRemotePath := "https://raw.githubusercontent.com/pear/Console_Table/master/Table.php"

	dest := fmt.Sprintf("%v%v.dvm%v.cache%v", usr.HomeDir, sep, sep, sep)
	destFile := fmt.Sprintf("%v%v", dest, ctFileName)

	versionPath := usr.HomeDir + sep + ".dvm" + sep + "versions" + sep + "drush-" + drushVersion.fullVersion + sep + "includes" + sep + "Table.php"

	ctPath := usr.HomeDir + sep + ".dvm" + sep + ".cache" + sep
	ctFile := ctPath + ctFileName

	if _, e := os.Stat(destFile); e == nil {
		log.Infof("Already downloaded '%v'.", ctFileName)
	} else {
		_, wgetErr := wget.Run(ctRemotePath)
		if wgetErr != nil {
			log.Infoln("wget returned error:", wgetErr)
		} else {
			copy(ctFile, destFile)
		}
	}

	if err := copy(destFile, versionPath); err == nil {
		log.Infoln("Fixed dependency issue with module Console_Table")
	}
}

// LegacyInstallVersion will install from a zip file which was located via git tags (manual input see ListLocal()).
//
// Deprecated: Drush version manager no longer supports legacy installs.
func (drushVersion *DrushVersion) LegacyInstallVersion() {
	// @TODO: Rewrite this.
	usr, _ := user.Current()
	zipFileName := drushVersion.fullVersion + ".zip"
	remotePath := "https://github.com/drush-ops/drush/archive/" + zipFileName
	zipPath := usr.HomeDir + sep + ".dvm" + sep + ".cache" + sep
	//zipFile := zipPath + zipFileName
	zipPathFull := fmt.Sprintf("%v%v.dvm%v.cache%v%v", usr.HomeDir, sep, sep, sep, zipFileName)
	if _, e := os.Stat(zipPathFull); e == nil {
		log.Infoln(fmt.Sprintf("Already downloaded Drush version v%v.", drushVersion.fullVersion))
	} else {
		log.Infoln(fmt.Sprintf("Downloading and extracting legacy Drush version v%v.", drushVersion.fullVersion))
		mkdir(zipPath, 0755)
		_, wgetErr := wget.Run(remotePath)
		if wgetErr != nil {
			log.Warnln("wget returned error:", wgetErr)
			log.Warnln(remotePath)
		}
	}
	dest := fmt.Sprintf("%v%v.dvm%vversions%v", usr.HomeDir, sep, sep, sep)
	destFile := fmt.Sprintf("%v%v", dest, zipFileName)

	copy(zipFileName, zipPathFull)
	copy(zipPathFull, destFile)
	exec.Command("sh", "-c", "cd "+dest+" && unzip "+zipFileName).Run()
	drushVersion.Status()

	remove(zipFileName)
}
