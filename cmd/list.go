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
	"github.com/spf13/cobra"

	"github.com/fubarhouse/dvm/data/versions"
)

var flagAvailable bool
var flagInstalled bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available or installed Drush versions.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		Drushes := versions.NewDrushVersionList()
		if flagAvailable == true {
			Drushes.PrintRemote()
		} else if flagInstalled == true {
			Drushes.PrintInstalled()
		} else {
			cmd.Help()
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&flagAvailable, "available", "a", false, "List available versions")
	listCmd.Flags().BoolVarP(&flagInstalled, "installed", "i", false, "List installed versions")
}
