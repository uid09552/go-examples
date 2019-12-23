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

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: Run,
}

func Run(cmd *cobra.Command, args []string) {
	author, err := cmd.Flags().GetString("author")
	title, err := cmd.Flags().GetString("title")
	content, err := cmd.Flags().GetString("content")
	if err != nil {
		fmt.Errorf("%v", err)
	}
	// Create a blog protobuffer message
	blog := &blogpg.Blog{
		AuthorId: author,
		Title:    title,
		Content:  content,
	}
	// RPC call
	res, err := client.CreateBlog(
		context.TODO(),
		// wrap the blog message in a CreateBlog request message
		&blogpg.CreateBlogReq{
			Blog: blog,
		},
	)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	fmt.Printf("Blog created: %s\n", res.Blog.Id)

}

func init() {
	createCmd.Flags().StringP("author", "a", "", "Add an author")
	createCmd.Flags().StringP("title", "t", "", "A title for the blog")
	createCmd.Flags().StringP("content", "c", "", "The content for the blog")
	createCmd.MarkFlagRequired("author")
	createCmd.MarkFlagRequired("title")
	createCmd.MarkFlagRequired("content")
	rootCmd.AddCommand(createCmd)
}
