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

// membersCmd represents the members command
var membersCmd = &cobra.Command{
	Use:   "members",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("members called")
		cfg := &engine.Config{}
		err := env.Parse(cfg)
		if err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(1)
		}
		cfg.BaseURL = baseURL
		engine.GetGroupMembers(cfg, groupID)
	},
}

func init() {
	groupsCmd.AddCommand(membersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// membersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// membersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
