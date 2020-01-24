package cmd

import (
	"github.com/spf13/cobra"
)

var (
	project string
	global  = false

	searchCmd = &cobra.Command{
		Use:   "search [kind] [query]",
		Short: "Search for GitLab for resources",
		Long: `Search through GitLab for projects, issues, merge requests,
groups, snippets, and other kinds of resources. The default limiting
context is your currently active group, but this can be changed to be a
global search or a project-specific search. Scope must be one of:
  - projects (invalid if using --project flag)
  - issues
  - mrs
  - milestones
  - snippets (only valid if using --global flag)
  - notes (only valid if using --project flag)
  - wikis
  - blobs
  - snippets
  - commits
  - users`,
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			logger.Println("Unimplemented")
		},
	}
)

func init() {
	searchCmd.Flags().BoolVarP(
		&global,
		"global",
		"g",
		false,
		"whether to search GitLab globally or not",
	)

	searchCmd.Flags().StringVarP(
		&project,
		"project",
		"p",
		"",
		"name of project to limit search to",
	)

	rootCmd.AddCommand(searchCmd)
}
