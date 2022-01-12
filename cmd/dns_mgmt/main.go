/*
Copyright © 2021-2022 Nikita Ivanovski info@slnt-opp.xyz

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"net"

	"github.com/go-redis/redis/v8"
	"github.com/slntopp/nocloud/pkg/dns"
	pb "github.com/slntopp/nocloud/pkg/dns/proto"
	"github.com/slntopp/nocloud/pkg/nocloud"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	port string
	log *zap.Logger

	redisHost string
)

func init() {
	viper.AutomaticEnv()
	log = nocloud.NewLogger()

	viper.SetDefault("PORT", "8080")
	viper.SetDefault("REDIS_HOST", "redis:6379")

	port = viper.GetString("PORT")
	redisHost = viper.GetString("REDIS_HOST")
}

func main() {
	defer func() {
		_ = log.Sync()
	}()

	log.Info("Setting up RedisDB Connection")
	rdb := redis.NewClient(&redis.Options{
        Addr:     redisHost,
        DB:       0,  // use default DB
    })
	log.Info("RedisDB connection established")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal("Failed to listen", zap.String("address", port), zap.Error(err))
	}
	s := grpc.NewServer()

	server := dns.NewDNSServer(log, rdb)
	pb.RegisterDNSServer(s, server)

	log.Info(fmt.Sprintf("Serving gRPC on 0.0.0.0:%v", port), zap.Skip())
	log.Fatal("Failed to serve gRPC", zap.Error(s.Serve(lis)))
}