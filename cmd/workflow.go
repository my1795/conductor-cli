/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cnd/cmd/util"
	"context"
	"encoding/json"
	"fmt"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/tidwall/pretty"

	_ "cnd/cmd/util"
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
var workflowListCmd = &cobra.Command{
	Use:   "list",
	Short: "lists workflows as summary",
	Long:  `To get full workflow definition please use workflow <<workflowName>>`,
	Run: func(cmd *cobra.Command, args []string) {
		var MetadataClient = client.MetadataResourceApiService{
			APIClient: client.NewAPIClient(nil, settings.NewHttpDefaultSettings()),
		}
		var callRes, _, _ = MetadataClient.GetAll(context.Background())
		var summaries []util.WorkflowDefSummary

		// Iterate through the callRes and generate summaries
		for _, item := range callRes {
			summary, _ := util.SummarizeWorkflowDef(&item)
			summaries = append(summaries, summary)
		}

		// Marshal the entire slice into JSON
		jsonData, err := json.MarshalIndent(summaries, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling summaries:", err)
			return
		}

		// Beautify the JSON using pretty package
		beautifiedSummary := pretty.Pretty(jsonData)

		// Print the JSON array
		fmt.Println(string(beautifiedSummary))

	},
}

func init() {
	workflowCmd.AddCommand(workflowListCmd)
	rootCmd.AddCommand(workflowCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workflowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workflowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
