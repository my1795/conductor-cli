/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cnd/cmd/util"
	"context"
	"fmt"
	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var workflowResourceService = client.WorkflowResourceApiService{
	APIClient: client.NewAPIClient(util.Authsettings, settings.NewHttpSettings(viper.GetString("baseurl"))),
}
var workflowCmd = &cobra.Command{
	Use:     "workflow",
	Aliases: []string{"wf"},
	Short:   "workflow resource related operations are included",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
var workflowRunningCmd = &cobra.Command{
	Use:     "show-running",
	Aliases: []string{"sr"},
	Short:   "shows running workflows with their ids",
	Long: `Workflow resources are instances those can be runnable of workflow definitions.
		this command targets runnable instances of workflow definitions. For example:
		
		cnd workflow showRunning <<workflowname>> returns list of running workflow instance IDs with given workflow names.

		Please visit /api/workflow in swagger documentation of the project. This command's scoped by workflow-resource api`,
	Run: func(cmd *cobra.Command, args []string) {
		version, _ := cmd.Flags().GetInt32("version")
		ver := client.WorkflowResourceApiGetRunningWorkflowOpts{
			Version: optional.NewInt32(version),
		}
		var workflowIds, _, _ = workflowResourceService.GetRunningWorkflow(context.Background(), args[0], &ver)

		// Print the string representation of the response body
		for _, value := range workflowIds {
			fmt.Println(value)
		}
	},
}

var showWorkflowCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"sh"},
	Short:   "shows workflow exectuions by their id",
	Long: `For example:
		
		cnd workflow show <<workflowId>>.

		Please visit api/workflow/{workflowId} in swagger documentation of the project. This command's scoped by workflow-resource api`,
	Run: func(cmd *cobra.Command, args []string) {
		includeTasks, _ := cmd.Flags().GetBool("include-tasks")
		tasks := client.WorkflowResourceApiGetExecutionStatusOpts{
			IncludeTasks: optional.NewBool(includeTasks),
		}
		var workflow, _, _ = workflowResourceService.GetExecutionStatus(context.Background(), args[0], &tasks)

		// Print the string representation of the response body
		util.PrintJSON(workflow)
	},
}

var restartWorkflowCmd = &cobra.Command{
	Use:     "restart",
	Aliases: []string{"re"},
	Short:   "restarts workflow executions by their id",
	Long: `For example:
		
		cnd workflow restart <<workflowId>> 

		Please visit api/workflow/{workflowId}/restart in swagger documentation of the project. This command's scoped by workflow-resource api`,
	Run: func(cmd *cobra.Command, args []string) {
		useLatestDef, _ := cmd.Flags().GetBool("use-latest-definition")
		restartOpts := client.WorkflowResourceApiRestartOpts{
			UseLatestDefinitions: optional.NewBool(useLatestDef),
		}
		var resp, err = workflowResourceService.Restart(context.Background(), args[0], &restartOpts)
		if err == nil {
			if resp.StatusCode >= 200 && resp.StatusCode <= 400 {
				fmt.Println("Workflow Restarted with id ", args[0])
			}
		} else {
			fmt.Println("Operation could not be handled by error", err)
		}
	},
}
var searchWorkflowCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"se"},
	Short:   "searches workflow executions by given free text",
	Long: `For example:
		
		cnd workflow search <<freetext>> 

		Please visit api/workflow/search-v2 in swagger documentation of the project. This command's scoped by workflow-resource api`,
	Run: func(cmd *cobra.Command, args []string) {
		searchOpts := client.WorkflowResourceApiSearchV2Opts{
			FreeText: optional.NewString(args[0]),
		}
		var searchResult, resp, err = workflowResourceService.SearchV2(context.Background(), &searchOpts)
		if err == nil {
			if resp.StatusCode >= 200 && resp.StatusCode <= 400 {
				util.PrintJSON(searchResult)
			}
		} else {
			fmt.Println("Operation could not be handled by error", err)
		}
	},
}
var terminateWorkflowCmd = &cobra.Command{
	Use:     "terminate",
	Aliases: []string{"te"},
	Short:   "terminates workflow executions by workflowId",
	Long: `For example:
		
		cnd workflow search <<workflowId>> 

		Please visit DELETE api/workflow/<<workflowId>> in swagger documentation of the project. This command's scoped by workflow-resource api`,
	Run: func(cmd *cobra.Command, args []string) {
		reason, _ := cmd.Flags().GetString("reason")
		triggerWorkflow, _ := cmd.Flags().GetBool("trigger-failure-wf")
		terminateOpts := client.WorkflowResourceApiTerminateOpts{
			Reason:                 optional.NewString(reason),
			TriggerFailureWorkflow: optional.NewBool(triggerWorkflow),
		}
		var resp, err = workflowResourceService.Terminate(context.Background(), args[0], &terminateOpts)
		if err == nil {
			if resp.StatusCode >= 200 && resp.StatusCode <= 400 {
				fmt.Println("Workflow terminated with id ", args[0])
			}
		} else {
			fmt.Println("Operation could not be handled by error", err)
		}
	},
}
var pauseWorkflowCmd = &cobra.Command{
	Use:     "pause",
	Aliases: []string{"pa"},
	Short:   "pauses workflow executions by workflowId",
	Long: `For example:
		
		cnd workflow pause <<workflowId>> 

		Please visit PUT api/workflow/<<workflowId>>/pause in swagger documentation of the project. This command's scoped by workflow-resource api`,
	Run: func(cmd *cobra.Command, args []string) {
		var resp, err = workflowResourceService.PauseWorkflow(context.Background(), args[0])
		if err == nil {
			if resp.StatusCode >= 200 && resp.StatusCode <= 400 {
				fmt.Println("Workflow paused with id ", args[0])
			}
		} else {
			fmt.Println("Operation could not be handled by error", err)
		}
	},
}
var resumeWorkflowCmd = &cobra.Command{
	Use:   "resume",
	Short: "resumes workflow executions by workflowId",
	Long: `For example:
		
		cnd workflow resume <<workflowId>> 

		Please visit PUT api/workflow/<<workflowId>>/resume in swagger documentation of the project. This command's scoped by workflow-resource api`,
	Run: func(cmd *cobra.Command, args []string) {
		var resp, err = workflowResourceService.ResumeWorkflow(context.Background(), args[0])
		if err == nil {
			if resp.StatusCode >= 200 && resp.StatusCode <= 400 {
				fmt.Println("Workflow paused with id ", args[0])
			}
		} else {
			fmt.Println("Operation could not be handled by error", err)
		}
	},
}
var retryWorkflowCmd = &cobra.Command{
	Use:   "retry",
	Short: "retries workflow executions by their id",
	Long: `For example:
		
		cnd workflow retry <<workflowId>> 

		Please visit api/workflow/{workflowId}/retry in swagger documentation of the project. This command's scoped by workflow-resource api`,
	Run: func(cmd *cobra.Command, args []string) {
		retryFlag, _ := cmd.Flags().GetBool("resume-sub-workflow-tasks")
		retryOpts := client.WorkflowResourceApiRetryOpts{
			ResumeSubworkflowTasks: optional.NewBool(retryFlag),
		}
		var resp, err = workflowResourceService.Retry(context.Background(), args[0], &retryOpts)
		if err == nil {
			if resp.StatusCode >= 200 && resp.StatusCode <= 400 {
				fmt.Println("Workflow retried with id ", args[0])
			}
		} else {
			fmt.Println("Operation could not be handled by error", err)
		}
	},
}

func init() {
	workflowCmd.PersistentFlags().Int32P("version", "v", 1, "version for workflow  resource")
	showWorkflowCmd.Flags().BoolP("include-tasks", "t", false, "includes task executions in the workflow executions")
	restartWorkflowCmd.Flags().BoolP("use-latest-definition", "d", false, "decides if latest workflow definition will be used or not")
	terminateWorkflowCmd.Flags().StringP("reason", "r", "", "reason for terminating the workflow")
	terminateWorkflowCmd.Flags().BoolP("trigger-failure-wf", "t", false, "Failure workflow trigger flag. triggers related failure workflow if exists")
	retryWorkflowCmd.Flags().BoolP("resume-sub-workflow-tasks", "r", false, "Resumes from sub workflow last failed task if exists")
	workflowCmd.AddCommand(workflowRunningCmd)
	workflowCmd.AddCommand(showWorkflowCmd)
	workflowCmd.AddCommand(restartWorkflowCmd)
	workflowCmd.AddCommand(searchWorkflowCmd)
	workflowCmd.AddCommand(terminateWorkflowCmd)
	workflowCmd.AddCommand(pauseWorkflowCmd)
	workflowCmd.AddCommand(resumeWorkflowCmd)
	workflowCmd.AddCommand(retryWorkflowCmd)
	rootCmd.AddCommand(workflowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workflowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workflowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
