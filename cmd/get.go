/*
Copyright © 2022 Sam Kim <netpple@gmail.com>
This file is part of Sam's CLI myapp.
*/
package cmd

import (
	"net/url"
	"errors"
	"github.com/spf13/cobra"
	"github.com/netpple/iget/downloader"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("enter the URL")
		}

		_, err := url.ParseRequestURI(args[0])
		if err != nil {
			return errors.New("invalid URL")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		urlString := args[0]

		dl := downloader.New(urlString)
		err := dl.Get()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
