/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/alwindoss/ergatis/internal/engine"
	"github.com/caarlos0/env/v6"
	"github.com/spf13/cobra"
)

// groupsCmd represents the groups command
var groupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("groups called")
		cfg := &engine.Config{}
		err := env.Parse(cfg)
		if err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(1)
		}
		cfg.BaseURL = baseURL
		engine.GetGroups(cfg, groupID)
	},
}

var (
	groupID string
	baseURL string
)

func init() {
	getCmd.AddCommand(groupsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// groupsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// groupsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	groupsCmd.Flags().StringVar(&groupID, "group-id", "", "--group-id 717337")
	groupsCmd.Flags().StringVar(&baseURL, "base-url", "", "--base-url \"https://git.rockylinux.org/api/v4\"")
	groupsCmd.MarkFlagRequired("group-id")
}
