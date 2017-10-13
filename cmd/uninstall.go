// Copyright © 2017 Karl Hepworth Karl.Hepworth@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/fubarhouse/dvm/version"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall a specific version of Drush",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			this := version.NewDrushVersion(args[0])
			this.Uninstall()
		} else {
			cmd.Help()
			logrus.Fatal("expected 1 argument, 0 found")
		}
	},
}

func init() {
	RootCmd.AddCommand(uninstallCmd)
}
