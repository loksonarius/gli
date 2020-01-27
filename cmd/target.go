package cmd

import (
	"strings"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

var (
	targetCmd = &cobra.Command{
		Use:   "target [name]",
		Short: "Manage saved GitLab targets",
		Long:  `Perform set, list, delete, and rename operations on targets.`,
	}

	targetSetCmd = &cobra.Command{
		Use:   "set [name]",
		Short: "Set a saved GitLab target as currently active",
		Long:  `Set a locally-saved target as the current active default.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			verifyTargetName(name)

			// Set the current active target to the one given if needed
			if name != Config.CurrentTarget {
				Config.CurrentTarget = name

				saveConfig()
			}

			logger.Printf("Set target %s as active default\n", name)
		},
	}

	targetListCmd = &cobra.Command{
		Use:   "list",
		Short: "List saved GitLab targets",
		Long:  `Print out a list of all saved GitLab targets.`,
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			t.AppendHeader(table.Row{
				"Active", "Name", "Endpoint", "Auth Type",
			})

			targets := Config.Targets
			for name, target := range targets {
				auth := target.Auth
				selected := "no"
				if name == Config.CurrentTarget {
					selected = "yes"
				}

				t.AppendRow(table.Row{
					selected, name, auth.Endpoint, auth.Type,
				})
			}

			t.SortBy([]table.SortBy{
				table.SortBy{Name: "Name", Mode: table.Asc},
			})
			logger.Println(t.Render())
		},
	}

	targetDeleteCmd = &cobra.Command{
		Use:   "delete [name]",
		Short: "Delete a saved GitLab target",
		Long: `Delete a currently saved GitLab target. If no another target is
still saved and the deleted target is the currently active target, then
the next alphabetically named target will be activated and printed out.`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			verifyTargetName(name)

			// Delete target and switch active target if necessary
			delete(Config.Targets, name)
			if name == Config.CurrentTarget {
				Config.CurrentTarget = ""
				for other := range Config.Targets {
					if Config.CurrentTarget == "" ||
						strings.Compare(Config.CurrentTarget, other) > 1 {
						Config.CurrentTarget = other
					}
				}

				logger.Println(
					"Deleting currently active target, need to switch",
				)
				logger.Printf(
					"Setting current target to %s\n",
					Config.CurrentTarget,
				)
			}

			saveConfig()

			logger.Printf("Deleted locally-saved target %s\n", name)
		},
	}

	targetRenameCmd = &cobra.Command{
		Use:   "rename [name] [new-name]",
		Short: "Rename a saved GitLab target",
		Long:  `Change the name of a currently saved GitLab target.`,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			name, newName := args[0], args[1]
			verifyTargetName(name)
			verifyTargetName(newName)
		},
	}
)

func init() {
	targetCmd.AddCommand(targetSetCmd)
	targetCmd.AddCommand(targetListCmd)
	targetCmd.AddCommand(targetDeleteCmd)
	targetCmd.AddCommand(targetRenameCmd)

	rootCmd.AddCommand(targetCmd)
}
