package main

import (
	"fmt"
	"net"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	inflog "github.com/infinimesh/infinimesh/pkg/log"
	"github.com/slntopp/nocloud/pkg/accounts"
	"github.com/slntopp/nocloud/pkg/accounts/accountspb"
	"github.com/slntopp/nocloud/pkg/graph"
)

var (
	port 			string
	log 			*zap.Logger

	arangodbHost 	string
	arangodbCred 	string
	nocloudRootPass string

	SIGNING_KEY 	[]byte
)

func init() {
	logger, err := inflog.NewProdOrDev()
	if err != nil {
		panic(err)
	}
	log = logger

	viper.AutomaticEnv()
	viper.SetDefault("PORT", "8080")

	viper.SetDefault("DB_HOST", "db:8529")
	viper.SetDefault("DB_CRED", "root:openSesame")
	viper.SetDefault("NOCLOUD_ROOT_PASSWORD", "secret")

	viper.SetDefault("SIGNING_KEY", "seeeecreet")

	port = viper.GetString("PORT")

	arangodbHost 	= viper.GetString("DB_HOST")
	arangodbCred 	= viper.GetString("DB_CRED")
	nocloudRootPass = viper.GetString("NOCLOUD_ROOT_PASSWORD")

	SIGNING_KEY 	= []byte(viper.GetString("SIGNING_KEY"))
}

func main() {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://" + arangodbCred + "@" + arangodbHost},
	})
	if err != nil {
		log.Fatal("Error creating connection to DB", zap.Error(err))
	}

	c, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
	})
	if err != nil {
		log.Fatal("Error creating driver instance for DB", zap.Error(err))
	}
	db, err := c.Database(nil, graph.DB_NAME)
	if err != nil {
		log.Fatal("Error getting instance of DB", zap.Error(err))
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal("Failed to listen", zap.String("address", port), zap.Error(err))
	}

	server := accounts.NewServer(log, db)
	server.SIGNING_KEY = SIGNING_KEY
	server.EnsureRootExists(nocloudRootPass)

	s := grpc.NewServer()
	
	accountspb.RegisterAccountsServiceServer(s, server)
	log.Info(fmt.Sprintf("Serving gRPC on 0.0.0.0:%v", port), zap.Skip())
	log.Fatal("Failed to serve gRPC", zap.Error(s.Serve(lis)))
}