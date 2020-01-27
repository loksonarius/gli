package cmd

import (
	"github.com/spf13/cobra"
)

var (
	targetCmd = &cobra.Command{
		Use:   "target [name]",
		Short: "Manage saved GitLab targets",
		Long:  `Perform list, delete, and rename operations on targets.`,
	}

	targetListCmd = &cobra.Command{
		Use:   "list",
		Short: "List saved GitLab targets",
		Long:  `Print out a list of all saved GitLab targets.`,
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			logger.Println("Unimplemented")
		},
	}

	targetDeleteCmd = &cobra.Command{
		Use:   "delete [name]",
		Short: "Delete a saved GitLab target",
		Long:  `Delete a currently saved GitLab target.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			logger.Println("Unimplemented")
		},
	}

	targetRenameCmd = &cobra.Command{
		Use:   "rename [name] [new-name]",
		Short: "Rename a saved GitLab target",
		Long:  `Change the name of a currently saved GitLab target.`,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			logger.Println("Unimplemented")
		},
	}
)

func init() {
	targetCmd.AddCommand(targetListCmd)
	targetCmd.AddCommand(targetDeleteCmd)
	targetCmd.AddCommand(targetRenameCmd)

	rootCmd.AddCommand(targetCmd)
}
