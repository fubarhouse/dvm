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

	"github.com/fubarhouse/dvm/versionlist"
)

var searchString = ""

// listCmd represents the list command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for available Drush versions using a substring.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if searchString != "" {
			versionlist.FindVersion(searchString)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVarP(&searchString, "substring", "s", "", "Substring to search for in available versions")

}
