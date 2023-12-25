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

var metadataCMD = &cobra.Command{
	Use:     "metadata",
	Aliases: []string{"mt"},
	Short:   "It is related metadata resource related operations",
	Long:    `Please visit /api/metadata in swagger documentation of the project. This command's scoped by metadata api'`,
	Run: func(cmd *cobra.Command, args []string) {
		isWorkflow, _ := cmd.Flags().GetBool("workflow")
		isTask, _ := cmd.Flags().GetBool("task")
		isAll, _ := cmd.Flags().GetBool("all")

		if isWorkflow {
			if isAll {
				printWorkflowDefs()
			}
		}
		if isTask {
			if isAll {

			}
		}
	},
}

func printWorkflowDefs() {
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
}

func init() {
	metadataCMD.Flags().BoolP("workflow", "w", false, "workflow flag for metadata resource")
	metadataCMD.Flags().BoolP("task", "t", false, "task flag for metadata resource")
	metadataCMD.Flags().BoolP("all", "a", false, "flag to get single element or all")

	rootCmd.AddCommand(metadataCMD)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// metadataCMD.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// metadataCMD.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
