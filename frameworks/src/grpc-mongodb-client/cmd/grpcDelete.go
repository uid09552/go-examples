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

// grpcDeleteCmd represents the grpcDelete command
var grpcDeleteCmd = &cobra.Command{
	Use:   "grpcDelete",
	Short: "grpc-delete",
	Long:  `delete grpc command`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		req := &blogpg.DeleteBlogReq{
			Id: id,
		}
		// We only return true upon success for other cases an error is thrown
		// We can thus omit the response variable for now and just print something to console
		_, err = client.DeleteBlog(context.Background(), req)
		if err != nil {
			return err
		}
		fmt.Printf("Succesfully deleted the blog with id %s\n", id)
		return nil

	},
}

func init() {
	grpcDeleteCmd.Flags().StringP("id", "i", "", "The id of the blog")
	grpcDeleteCmd.MarkFlagRequired("id")

	grpcCmd.AddCommand(grpcDeleteCmd)
}
