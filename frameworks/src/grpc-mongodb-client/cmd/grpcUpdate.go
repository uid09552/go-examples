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

// grpcUpdateCmd represents the grpcUpdate command
var grpcUpdateCmd = &cobra.Command{
	Use:   "grpcUpdate",
	Short: "grpc update",
	Long:  `update an existing record via grpc`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get the flags from CLI
		id, err := cmd.Flags().GetString("id")
		author, err := cmd.Flags().GetString("author")
		title, err := cmd.Flags().GetString("title")
		content, err := cmd.Flags().GetString("content")
		blog := &blogpg.Blog{
			Id:       id,
			AuthorId: author,
			Title:    title,
			Content:  content,
		}
		// Create an UpdateBlogRequest
		req := &blogpg.UpdateBlogReq{
			Blog: blog,
		}

		res, err := client.UpdateBlog(context.Background(), req)
		if err != nil {
			fmt.Errorf("%v", err)
		}

		fmt.Println(res)

	},
}

func init() {
	grpcUpdateCmd.Flags().StringP("id", "i", "", "The id of the blog")
	grpcUpdateCmd.Flags().StringP("author", "a", "", "Add an author")
	grpcUpdateCmd.Flags().StringP("title", "t", "", "A title for the blog")
	grpcUpdateCmd.Flags().StringP("content", "c", "", "The content for the blog")
	grpcUpdateCmd.MarkFlagRequired("id")

	grpcCmd.AddCommand(grpcUpdateCmd)
}
