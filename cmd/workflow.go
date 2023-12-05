/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/tidwall/pretty"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/spf13/cobra"
)

// workflowCmd represents the workflow command
var workflowCmd = &cobra.Command{
	Use:   "workflow",
	Short: "workflow related resource operations",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("workflow called")

		var MetadataClient = client.MetadataResourceApiService{
			APIClient: client.NewAPIClient(nil, settings.NewHttpDefaultSettings()),
		}
		var callRes, _, _ = MetadataClient.Get(context.Background(), args[0], nil)
		// Marshal the struct into a JSON string
		jsonData, err := json.MarshalIndent(callRes, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}

		// Use pretty to beautify the JSON string
		beautifiedJSON := pretty.Pretty(jsonData)

		// Print the beautified JSON
		fmt.Println(string(beautifiedJSON))

	},
}

func init() {
	rootCmd.AddCommand(workflowCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workflowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workflowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
