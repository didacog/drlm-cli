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
	"context"
	"fmt"
	"log"
	"time"

	"github.com/brainupdaters/drlm-cli/lib"
	pb "github.com/brainupdaters/drlm-comm/drlmcomm"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// userListCmd represents the userList command
var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "List DRLM server users",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		user := ""
		pass := ""
		fmt.Println("drlm-cli user list called")

		conn, err := grpc.Dial(lib.Config.Drlmcore.Server+":"+lib.Config.Drlmcore.Port, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		client := pb.NewDrlmApiClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := client.ListUser(ctx, &pb.UserRequest{User: user, Pass: pass})
		if err != nil {
			log.Fatalf("could not lists users: %v", err)
		}

		log.Printf("Response DRLM-Core Server: %s", r.Message)
	},
}

func init() {
	// This is a command related to user and is added to userCmd
	userCmd.AddCommand(userListCmd)

	// Here you will define your flags and configuration settings.
}
