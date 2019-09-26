package version

import (
	log "github.com/Sirupsen/logrus"
	"os"
	"os/exec"
	"os/user"
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
	log.Infoln("Fixing dependency issue with module Console_Table")
	ctFileName := "Table.inc"
	ctRemotePath := "https://raw.githubusercontent.com/pear/Console_Table/master/Table.php"
	ctPath := usr.HomeDir + sep + ".dvm" + sep + "versions" + sep + "drush-" + drushVersion.fullVersion + sep + "includes" + sep
	ctFile := ctPath + ctFileName

	_, wgetErr := exec.Command("wget", ctRemotePath).Output()
	if wgetErr != nil {
		log.Infoln("wget returned error:", wgetErr)
	}
	tmpFile := "Table.php"
	exec.Command("mv", "./" + tmpFile, ctFile).Run()
}

// LegacyInstallVersion will install from a zip file which was located via git tags (manual input see ListLocal()).
//
// Deprecated: Drush version manager no longer supports legacy installs.
func (drushVersion *DrushVersion) LegacyInstallVersion() {
	usr, _ := user.Current()
	log.Infoln("Downloading and extracting legacy Drush version ", drushVersion.majorVersion)
	zipFileName := drushVersion.fullVersion + ".zip"
	remotePath := "https://github.com/drush-ops/drush/archive/" + zipFileName
	zipPath := usr.HomeDir + sep + ".dvm" + sep + "versions" + sep
	zipFile := zipPath + zipFileName
	exec.Command("sh", "-c", "mkdir -p "+zipPath).Run()
	_, wgetErr := exec.Command("wget", remotePath).Output()
	if wgetErr != nil {
		log.Warnln("wget returned error:", wgetErr)
		log.Warnln(remotePath)
	}
	exec.Command("sh", "-c", "mv "+zipFileName+" "+zipPath).Run()
	exec.Command("sh", "-c", "cd "+zipPath+" && unzip "+zipFile).Run()
	exec.Command("sh", "-c", "rm -f "+zipFile).Run()
	drushVersion.Status()
}
