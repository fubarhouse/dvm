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
	usr, _ := user.Current()
	pwd, _ := os.Getwd()
	log.Infoln("Fixing dependency issue with module Console_Table")
	ctFileName := "Table.inc"
	ctRemotePath := "https://raw.githubusercontent.com/pear/Console_Table/master/Table.php"
	ctPath := usr.HomeDir + sep + ".dvm" + sep + "versions" + sep + "drush-" + drushVersion.fullVersion + sep + "includes" + sep
	ctFile := ctPath + ctFileName

	_, wgetErr := wget.Run(ctRemotePath)
	if wgetErr != nil {
		log.Infoln("wget returned error:", wgetErr)
	}
	tmpFile := fmt.Sprintf("%v%vTable.php", pwd, sep)
	move(tmpFile, ctFile)
}

// LegacyInstallVersion will install from a zip file which was located via git tags (manual input see ListLocal()).
//
// Deprecated: Drush version manager no longer supports legacy installs.
func (drushVersion *DrushVersion) LegacyInstallVersion() {
	usr, _ := user.Current()
	log.Infoln(fmt.Sprintf("Downloading and extracting legacy Drush version %v.", drushVersion.fullVersion))
	zipFileName := drushVersion.fullVersion + ".zip"
	remotePath := "https://github.com/drush-ops/drush/archive/" + zipFileName
	zipPath := usr.HomeDir + sep + ".dvm" + sep + "versions" + sep
	zipFile := zipPath + zipFileName
	zipPathFull := fmt.Sprintf("%v%v.dvm%vversions%v%v", usr.HomeDir, sep, sep, sep, zipFileName)
	mkdir(zipPath, 0755)
	_, wgetErr := wget.Run(remotePath)
	if wgetErr != nil {
		log.Warnln("wget returned error:", wgetErr)
		log.Warnln(remotePath)
	}
	move(zipFileName, zipPathFull)
	exec.Command("sh", "-c", "cd "+zipPath+" && unzip "+zipFileName).Run()
	remove(zipFile)
	drushVersion.Status()
}
