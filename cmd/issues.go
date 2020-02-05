package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/text"
	"github.com/mitchellh/go-wordwrap"
	"github.com/spf13/cobra"
)

var (
	issueCmd = &cobra.Command{
		Use:   "issue",
		Short: "Interact with GitLab issues",
		Long: `Create, find, view, and update issues from any project under the
current group. Project arguments for subcommands accept Project ID's.
Project ID's are one of:
  - ID: the ID of the project (such as the one printed by the ls subcommand)
  - relative path: path to the project from the current group including
  - absolute path: the path to the project starting with / and group name`,
	}

	issueCloseCmd = &cobra.Command{
		Use:   "close [project] [issue]",
		Short: "Closes an open issue",
		Long: `Closes the given issue. Can be undone using the issue reopen
command.`,
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			logger.Println("command unimplemented")
		},
	}

	issueCommentCmd = &cobra.Command{
		Use:   "update [project] [issue]",
		Short: "Adds an update to an issue",
		Long: `Opens your local editor (lookup $EDITOR) to let you add a single
comment to the given issue or edit any of the fields in an issue. GitLab
enables users to edit almost any field in an issue using comments by
using '/'-prefixed commands. For example, a comment that has a line
starting with '/label security' will have that line removed from the
comment body and cause the issue to be updated with the 'security'
label. Other available edit methods include /close, /assign, /move.
GitLab doesn't seem to publicly document this feature, but to play with
it interactively, go to any ticket in your browser and start a blank
line in a comment with '/' to see available options.`,
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			logger.Println("command unimplemented")
		},
	}

	issueHistoryCmd = &cobra.Command{
		Use:   "history [project] [issue]",
		Short: "Print events for issues",
		Long: `Prints out all events (comments, changes, updates, etc) for
issues. If a project ID is specified, all events will be filtered to
just that project. If an issue ID is specified on top of a project ID,
then results will be further filtered to events on just that issue.`,
		Args: cobra.MaximumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			logger.Println("command unimplemented")
		},
	}

	issueListCmd = &cobra.Command{
		Use:   "list [project]",
		Short: "List issues under current group",
		Long: `Lists all issues under the current group. If a project ID is
specified, results will be limited to issues under that project.`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			logger.Println("command unimplemented")
		},
	}

	issueViewCmd = &cobra.Command{
		Use:   "view [project] [issue]",
		Short: "View details for an issue",
		Long: `Prints out a summary of the given issue including status,
labels, assignee, milestone, time tracking, and comments. For a full
history of an issue, consider the 'issue history' command.`,
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			target := Config.Targets[Config.CurrentTarget]
			path := resolvePath(target.CurrentGroup, args[0])
			issueID, err := strconv.Atoi(args[1])
			if _, err := strconv.Atoi(args[1]); err != nil {
				logger.Fatalln("Issue arg must be an int")
			}
			if _, err := strconv.Atoi(args[0]); err == nil {
				// project seems ID'd by an int
				// we can use the ID directly instead of a full path
				path = args[0]
			} else {
				// Remove leading '/' to get namespaced project ID
				path = path[1:]
			}

			client := getClient(target)
			_, _, err = client.Projects.GetProject(path, nil)
			if err != nil {
				logger.Fatalf(
					"Received error getting project: %v",
					err,
				)
			}

			issue, _, err := client.Issues.GetIssue(path, issueID)
			if err != nil {
				logger.Fatalf(
					"Received error getting issue: %v",
					err,
				)
			}

			status := text.Faint.Sprint(issue.State)
			id := fmt.Sprintf("%d", issueID)
			title := text.Bold.Sprint(issue.Title)
			milestone := "No Milestone"
			if issue.Milestone != nil {
				milestone = issue.Milestone.Title
			}
			labels := strings.Join(issue.Labels, text.Faint.Sprint(", "))
			assigneeText := "currently unassigned"
			if issue.Assignee != nil {
				assigneeText = fmt.Sprintf(
					"and assigned to %s",
					issue.Assignee.Name,
				)
			}
			description := "| " + strings.ReplaceAll(
				wordwrap.WrapString(issue.Description, 80),
				"\n",
				"\n| ",
			)

			logger.Printf(
				"[%s] %s\nStatus: %s %s\n%s\nMilestone: %s\nLabels: %s\n",
				id,
				title,
				status,
				assigneeText,
				description,
				milestone,
				labels,
			)
		},
	}
)

func init() {
	issueCmd.AddCommand(issueCommentCmd)
	issueCmd.AddCommand(issueHistoryCmd)
	issueCmd.AddCommand(issueListCmd)
	issueCmd.AddCommand(issueViewCmd)

	rootCmd.AddCommand(issueCmd)
}
