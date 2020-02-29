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

// grpcReadCmd represents the grpcRead command
var grpcReadCmd = &cobra.Command{
	Use:   "grpcRead",
	Short: "grpc read",
	Long:  `read via grpc by id`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			fmt.Errorf("%v", err)
		}
		req := &blogpg.ReadBlogReq{
			Id: id,
		}
		res, err := client.ReadBlog(context.Background(), req)
		if err != nil {
			fmt.Errorf("%v", err)
		}
		fmt.Println(res)
	},
}

func init() {
	grpcReadCmd.Flags().StringP("id", "i", "", "The id of the blog")
	grpcReadCmd.MarkFlagRequired("id")

	grpcCmd.AddCommand(grpcReadCmd)

}
