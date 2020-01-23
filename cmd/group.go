package cmd

import (
	"github.com/spf13/cobra"
)

var (
	groupCmd = &cobra.Command{
		Use:   "group",
		Short: "Navaigate and activate groups",
		Long: `Groups are treated like a namespacing tool throughout this tool.
An easy way to understand the metaphor employed is to think of a Group's
path as filesystem path and all the projects, issues, and such resources
as files. Most operations that involve listing or searching will also
look through sub-groups as well.`,
	}

	recursive  = false
	groupLsCmd = &cobra.Command{
		Use:   "ls [paths]",
		Short: "View group paths",
		Long:  `List subgroups under the currently active group.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			logger.Println("Unimplemented")
		},
	}

	groupCdCmd = &cobra.Command{
		Use:   "cd [path]",
		Short: "Navigate group paths",
		Long: `Change active group to given path. The target path can be one of:
  - /:        ls here to view accessible groups, deactivates current group
  - /{user}:  navigate to user's profile
  - ../:      navigate one level up the group path per repetition of '../'
  - foo/bar:  navigate to the relative path 'foo/bar'
  - /foo/bar: navigate to absolute group path /foo/bar/car`,
		Run: func(cmd *cobra.Command, args []string) {
			logger.Println("Unimplemented")
		},
	}
)

func init() {
	groupLsCmd.Flags().BoolVarP(
		&recursive,
		"recursive",
		"r",
		false,
		"whether to list subgroups recursively or not",
	)

	groupCmd.AddCommand(groupLsCmd)
	groupCmd.AddCommand(groupCdCmd)
}
