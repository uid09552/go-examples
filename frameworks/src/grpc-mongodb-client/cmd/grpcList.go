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
	"io"

	"github.com/spf13/cobra"
)

// grpcListCmd represents the grpcList command
var grpcListCmd = &cobra.Command{
	Use:   "grpcList",
	Short: "list grpc blogs",
	Long:  `list grpc blogs`,
	RunE: func(cmd *cobra.Command, args []string) error {
		req := &blogpg.ListBlogReq{}
		stream, err := client.ListBlogs(context.Background(), req)
		if err != nil {
			return err
		}
		for {
			// stream.Recv returns a pointer to a ListBlogRes at the current iteration
			res, err := stream.Recv()
			// If end of stream, break the loop
			if err == io.EOF {
				break
			}
			// if err, return an error
			if err != nil {
				return err
			}
			// If everything went well use the generated getter to print the blog message
			fmt.Println(res.GetBlog())
		}
		return nil
	},
}

func init() {
	grpcCmd.AddCommand(grpcListCmd)
}
