/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
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
)

var taskResourceService = client.TaskResourceApiService{
	APIClient: client.NewAPIClient(nil, settings.NewHttpDefaultSettings()),
}

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:     "task",
	Aliases: []string{"t"},
	Short:   "Task related operations",
	Long:    `This sub command is scoped by task resource operations of the workflow orchestrator. `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("task called")
	},
}
var showTaskCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"sh"},
	Short:   "shows task executions by their id",
	Long: `For example:
		
		cnd task show <<taskId>>.

		Please visit api/tasks/{taskId} in swagger documentation of the project. This command's scoped by task-resource api`,
	Run: func(cmd *cobra.Command, args []string) {
		isLog, _ := cmd.Flags().GetBool("logs")
		if isLog {
			var taskLog, resp, err = taskResourceService.GetTaskLogs(context.Background(), args[0])
			if err == nil {
				if resp.StatusCode >= 200 && resp.StatusCode <= 400 {
					util.PrintJSON(taskLog)
				}
			} else {
				fmt.Println("Operation could not be handled by error", err)
			}
		} else {
			var task, resp, err = taskResourceService.GetTask(context.Background(), args[0])
			if err == nil {
				if resp.StatusCode >= 200 && resp.StatusCode <= 400 {
					util.PrintJSON(task)
				}
			} else {
				fmt.Println("Operation could not be handled by error", err)
			}
		}
	},
}
var queueAllCmd = &cobra.Command{
	Use:     "queue",
	Aliases: []string{"q"},
	Short:   "get all task queues",
	Long: `For example:
		
		cnd task queue.

		Please visit api/tasks/queue/all in swagger documentation of the project. This command's scoped by task-resource api`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		if verbose {
			var queueData, resp, err = taskResourceService.AllVerbose(context.Background())
			if err == nil {
				if resp.StatusCode >= 200 && resp.StatusCode <= 400 {
					util.PrintJSON(queueData)
				}
			} else {
				fmt.Println("Operation could not be handled by error", err)
			}
		} else {
			var queueDataVerbose, resp, err = taskResourceService.All(context.Background())
			if err == nil {
				if resp.StatusCode >= 200 && resp.StatusCode <= 400 {
					util.PrintJSON(queueDataVerbose)
				}
			} else {
				fmt.Println("Operation could not be handled by error", err)
			}
		}
	},
}
var pollDataCmd = &cobra.Command{
	Use:     "poll-data",
	Aliases: []string{"pd"},
	Short:   "get polling data for  task queues",
	Long: `For example:
		
		cnd task poll-data <<taskType>>.

		Please visit api/tasks/queue/polldata in swagger documentation of the project. This command's scoped by task-resource api`,
	Run: func(cmd *cobra.Command, args []string) {
		isAll, _ := cmd.Flags().GetBool("all")
		if isAll {
			var allPollData, resp, err = taskResourceService.GetAllPollData(context.Background())
			if err == nil {
				if resp.StatusCode >= 200 && resp.StatusCode <= 400 {
					util.PrintJSON(allPollData)
				}
			} else {
				fmt.Println("Operation could not be handled by error", err)
			}
		} else {
			var pollData, resp, err = taskResourceService.GetPollData(context.Background(), args[0])
			if err == nil {
				if resp.StatusCode >= 200 && resp.StatusCode <= 400 {
					util.PrintJSON(pollData)
				}
			} else {
				fmt.Println("Operation could not be handled by error", err)
			}
		}
	},
}

var queueSizeCmd = &cobra.Command{
	Use:     "queue-size",
	Aliases: []string{"qs"},
	Short:   "get queue size for  task queues",
	Long: `For example:
		
		cnd task queue-size <<taskType>>.

		Please visit api/tasks/queue/size in swagger documentation of the project. This command's scoped by task-resource api`,
	Run: func(cmd *cobra.Command, args []string) {
		taskOpts := client.TaskResourceApiSizeOpts{
			TaskType: optional.NewInterface(args[0]),
		}
		var queueSizes, resp, err = taskResourceService.Size(context.Background(), &taskOpts)
		if err == nil {
			if resp.StatusCode >= 200 && resp.StatusCode <= 400 {
				util.PrintJSON(queueSizes)
			}
		} else {
			fmt.Println("Operation could not be handled by error", err)
		}
	},
}

func init() {
	showTaskCmd.Flags().BoolP("logs", "l", false, "Show execution logs of the fgiven task")
	queueAllCmd.Flags().StringP("worker-id", "w", "", "Worker id of the task to be polled")
	queueAllCmd.Flags().StringP("domain", "d", "", "domain of the task to be polled")
	queueAllCmd.Flags().BoolP("verbose", "v", false, "Get task queues verbose flag")
	pollDataCmd.Flags().BoolP("all", "a", false, "Flag to get all poll data")
	taskCmd.AddCommand(showTaskCmd)
	taskCmd.AddCommand(queueAllCmd)
	taskCmd.AddCommand(pollDataCmd)
	taskCmd.AddCommand(queueSizeCmd)
	rootCmd.AddCommand(taskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// taskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
