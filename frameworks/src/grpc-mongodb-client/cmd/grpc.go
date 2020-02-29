/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"blogpg"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var client blogpg.BlogServiceClient
var client2 blogpg.EmptyServiceClient

var requestCtx context.Context
var requestOpts grpc.DialOption

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "grpc commands",
	Long:  `grpc commands`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		server := "localhost:50051"
		server = viper.GetString("server")
		if server == "" {
			server = "localhost:50051"
		}
		fmt.Printf("Connecting to server %v\n", server)
		// Establish context to timeout after 10 seconds if server does not respond
		requestCtx, _ = context.WithTimeout(context.Background(), 10*time.Second)
		// Establish insecure grpc options (no TLS)
		requestOpts = grpc.WithInsecure()
		// Dial the server, returns a client connection
		conn, err := grpc.Dial(server, requestOpts)
		if err != nil {
			log.Fatalf("Unable to establish client connection to localhost:50051: %v", err)
		}
		// Instantiate the BlogServiceClient with our client connection to the server
		client = blogpg.NewBlogServiceClient(conn)
		client2 = blogpg.NewEmptyServiceClient(conn)
		viper.WriteConfig()
	},
}

func init() {
	grpcCmd.PersistentFlags().StringP("server", "s", "", "server address of grpc server")
	viper.BindPFlag("server", grpcCmd.PersistentFlags().Lookup("server"))
	viper.BindEnv("GRPC_SERVER")
	rootCmd.AddCommand(grpcCmd)
}
