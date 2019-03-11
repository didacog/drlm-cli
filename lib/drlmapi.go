package lib

import (
	"context"
	"time"

	pb "github.com/brainupdaters/drlm-common/comms"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type DrlmcoreConfig struct {
	Server   string
	Port     string
	Tls      bool
	Cert     string
	User     string
	Password string
}

type Authentication struct {
	Login    string
	Password string
}

func SetDrlmcoreConfigDefaults() {
	viper.SetDefault("drlmcore.server", "localhost")
	viper.SetDefault("drlmcore.port", "50051")
	viper.SetDefault("drlmcore.tls", true)
	viper.SetDefault("drlmcore.cert", "cert/server.crt")
	viper.SetDefault("drlmcore.user", "drlmadmin")
	viper.SetDefault("drlmcore.password", "drlm3api")
}
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"login":    a.Login,
		"password": a.Password,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires transport security
func (a *Authentication) RequireTransportSecurity() bool {
	return Config.Drlmcore.Tls
}

var conn *grpc.ClientConn
var ctx context.Context
var client pb.DrlmApiClient
var cancel context.CancelFunc

func InitConnexion() {
	var err error
	var creds credentials.TransportCredentials

	if Config.Drlmcore.Tls {
		// Create the client TLS credentials
		creds, err = credentials.NewClientTLSFromFile(Config.Drlmcore.Cert, "")
		if err != nil {
			log.Fatalf("could not load tls cert: %s", err)
		}
	}

	// Setup the login/pass
	auth := Authentication{
		Login:    Config.Drlmcore.User,
		Password: Config.Drlmcore.Password,
	}

	// Initiate a connection with the server
	if Config.Drlmcore.Tls {
		conn, err = grpc.Dial(Config.Drlmcore.Server+":"+Config.Drlmcore.Port, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&auth))
	} else {
		conn, err = grpc.Dial(Config.Drlmcore.Server+":"+Config.Drlmcore.Port, grpc.WithPerRPCCredentials(&auth), grpc.WithInsecure())
	}
	if err != nil {
		log.Fatal("did not connect: " + err.Error())
	}

	client = pb.NewDrlmApiClient(conn)
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	
}

func APIUserAdd(usr *User) {
	InitConnexion()
	defer conn.Close()
	defer cancel()

	r, err := client.AddUser(ctx, &pb.UserRequest{User: usr.User, Pass: usr.Password})
	if err != nil {
		log.Fatal("could not add user: " + err.Error())
	}

	log.Info("Response DRLM-Core Server: " + r.Message)
}

func APIUserDelete(usr *User) {
	InitConnexion()
	defer conn.Close()
	defer cancel()

	r, err := client.DelUser(ctx, &pb.UserRequest{User: usr.User, Pass: usr.Password})
	if err != nil {
		log.Fatalf("could not delete user: %v", err)
	}

	log.Printf("Response DRLM-Core Server: %s", r.Message)
}

func APIUserList(){
	InitConnexion()
	defer conn.Close()
	defer cancel()

	r, err := client.ListUser(ctx, &pb.UserRequest{User: "", Pass: ""})
	if err != nil {
		log.Fatalf("could not lists users: %v", err)
	}

	log.Printf("Response DRLM-Core Server: %s", r.Message)
}