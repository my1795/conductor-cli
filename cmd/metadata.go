/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cnd/cmd/util"
	_ "cnd/cmd/util"
	"context"
	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var MetadataClient = client.MetadataResourceApiService{
	APIClient: client.NewAPIClient(util.Authsettings, settings.NewHttpSettings(viper.GetString("baseurl"))),
}
var metadataCMD = &cobra.Command{
	Use:     "metadata",
	Aliases: []string{"mt"},
	Short:   "It is related metadata resource related operations",
	Long:    `Please visit /api/metadata in swagger documentation of the project. This command's scoped by metadata api'`,
	Run: func(cmd *cobra.Command, args []string) {
		isWorkflow, _ := cmd.Flags().GetBool("workflow")
		isTask, _ := cmd.Flags().GetBool("task")
		isAll, _ := cmd.Flags().GetBool("all")
		version, _ := cmd.Flags().GetInt32("version")

		if isWorkflow {
			if isAll {
				printWorkflowDefs()
			} else {
				printWorkflowDef(args[0], optional.NewInt32(version))
			}
		}
		if isTask {
			if isAll {
				printTaskDefs()
			} else {
				printTaskDef(args[0])
			}
		}
	},
}

func printWorkflowDefs() {
	var callRes, _, _ = MetadataClient.GetAll(context.Background())
	var summaries []model.WorkflowDef

	// Iterate through the callRes and generate summaries
	for _, item := range callRes {
		//summary, _ := util.SummarizeWorkflowDef(&item)
		summaries = append(summaries, item)
	}

	util.PrintJSON(summaries)
}
func printWorkflowDef(name string, version optional.Int32) {
	ver := client.MetadataResourceApiGetOpts{
		Version: version,
	}
	var callRes, _, _ = MetadataClient.Get(context.Background(), name, &ver)
	util.PrintJSON(callRes)
}
func printTaskDefs() {
	var callRes, _, _ = MetadataClient.GetTaskDefs(context.Background())
	var summaries []model.TaskDef

	// Iterate through the callRes and generate summaries
	for _, item := range callRes {
		//summary, _ := util.SummarizeWorkflowDef(&item)
		summaries = append(summaries, item)
	}

	util.PrintJSON(summaries)
}
func printTaskDef(name string) {
	var callRes, _, _ = MetadataClient.GetTaskDef(context.Background(), name)
	util.PrintJSON(callRes)
}

func init() {
	metadataCMD.Flags().BoolP("workflow", "w", false, "workflow flag for metadata resource")
	metadataCMD.Flags().BoolP("task", "t", false, "task flag for metadata resource")
	metadataCMD.Flags().BoolP("all", "a", false, "flag to get single element or all")
	metadataCMD.Flags().Int32P("version", "v", 1, "version for workflow metadata resource")

	rootCmd.AddCommand(metadataCMD)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// metadataCMD.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// metadataCMD.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
