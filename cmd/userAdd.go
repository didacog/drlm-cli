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
	"time"

	"github.com/brainupdaters/drlm-cli/lib"
	pb "github.com/brainupdaters/drlm-common/comms"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Authentication struct {
	Login    string
	Password string
}

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
	// add Password flag
	userAddCmd.Flags().StringP("pass", "p", "", "Password")
}

// GetRequestMetadata gets the current request metadata
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"login":    a.Login,
		"password": a.Password,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires transport security
func (a *Authentication) RequireTransportSecurity() bool {
	return lib.Config.Drlmcore.Tls
}

func runUserAdd(cmd *cobra.Command, args []string) {

	var conn *grpc.ClientConn
	var err error
	var creds credentials.TransportCredentials

	if lib.Config.Drlmcore.Tls {
		// Create the client TLS credentials
		creds, err = credentials.NewClientTLSFromFile(lib.Config.Drlmcore.Cert, "")
		if err != nil {
			log.Fatalf("could not load tls cert: %s", err)
		}
	}

	// Setup the login/pass
	auth := Authentication{
		Login:    lib.Config.Drlmcore.User,
		Password: lib.Config.Drlmcore.Password,
	}

	// Initiate a connection with the server
	if lib.Config.Drlmcore.Tls {
		conn, err = grpc.Dial(lib.Config.Drlmcore.Server+":"+lib.Config.Drlmcore.Port, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&auth))
	} else {
		conn, err = grpc.Dial(lib.Config.Drlmcore.Server+":"+lib.Config.Drlmcore.Port, grpc.WithPerRPCCredentials(&auth), grpc.WithInsecure())
	}
	if err != nil {
		log.Fatal("did not connect: " + err.Error())
	}
	defer conn.Close()

	client := pb.NewDrlmApiClient(conn)

	// Data to pass the message
	user := cmd.Flag("user").Value.String()
	pass := cmd.Flag("pass").Value.String()

	// Screen information
	fmt.Println("drlm-cli user add called")
	fmt.Println("User: " + user)
	fmt.Println("Password: " + pass)

	// What does context??
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.AddUser(ctx, &pb.UserRequest{User: user, Pass: pass})
	if err != nil {
		log.Fatal("could not add user: " + err.Error())
	}

	log.Info("Response DRLM-Core Server: " + r.Message)
}
