package cmd

import (
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
	}

	issueHistoryCmd = &cobra.Command{
		Use:   "history [project] [issue]",
		Short: "Print events for issues",
		Long: `Prints out all events (comments, changes, updates, etc) for
issues. If a project ID is specified, all events will be filtered to
just that project. If an issue ID is specified on top of a project ID,
then results will be further filtered to events on just that issue.`,
		Args: cobra.MaximumNArgs(2),
	}

	issueListCmd = &cobra.Command{
		Use:   "list [project]",
		Short: "List issues under current group",
		Long: `Lists all issues under the current group. If a project ID is
specified, results will be limited to issues under that project.`,
		Args: cobra.MaximumNArgs(1),
	}

	issueViewCmd = &cobra.Command{
		Use:   "view [project] [issue]",
		Short: "View details for an issue",
		Long: `Prints out a summary of the given issue including status,
labels, assignee, milestone, time tracking, and comments. For a full
history of an issue, consider the 'issue history' command.`,
		Args: cobra.ExactArgs(2),
	}
)

func init() {
	issueCmd.AddCommand(issueCommentCmd)
	issueCmd.AddCommand(issueHistoryCmd)
	issueCmd.AddCommand(issueListCmd)
	issueCmd.AddCommand(issueViewCmd)

	rootCmd.AddCommand(issueCmd)
}
