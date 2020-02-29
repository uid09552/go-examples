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

	"github.com/spf13/cobra"
)

// grpcEmptyCmd represents the grpcEmpty command
var grpcEmptyCmd = &cobra.Command{
	Use:   "grpcEmpty",
	Short: "grpc-do nothing call",
	Long:  `grpc empty call, respond nothing, pass nothing, just for testing the connection`,
	Run:   RunEmpty,
}

func RunEmpty(cmd *cobra.Command, args []string) {
	empty := &blogpg.Empty{}
	res, err := client2.DoNothing(
		context.TODO(),
		// wrap the blog message in a CreateBlog request message
		empty,
	)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	fmt.Println("run empty finished")
	fmt.Printf("%v", res)
}

func init() {
	grpcCmd.AddCommand(grpcEmptyCmd)
}
