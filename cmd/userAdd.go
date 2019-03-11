// Copyright © 2019 Pau Roura <pau@brainupdaters.net>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"github.com/brainupdaters/drlm-cli/lib"
	"github.com/spf13/cobra"
)

// userAddCmd represents the userAdd command
var userAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new user to DRLM server",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runUserAdd,
}

func init() {
	// This is a command related to user and is added to userCmd
	userCmd.AddCommand(userAddCmd)

	// Here you will define your flags and configuration settings.
	// add User name flag and mark as required
	userAddCmd.Flags().StringP("user", "u", "", "User name")
	userAddCmd.MarkFlagRequired("user")
	userAddCmd.Flags().StringP("pass", "p", "", "Password")
	userAddCmd.MarkFlagRequired("pass")
}

func runUserAdd(cmd *cobra.Command, args []string) {
	usr := lib.User{User: cmd.Flag("user").Value.String(), Password: cmd.Flag("pass").Value.String()}
	lib.APIUserAdd(&usr)
}
