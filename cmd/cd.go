package cmd

import (
	"path"

	"github.com/spf13/cobra"
)

var (
	cdCmd = &cobra.Command{
		Use:   "cd [path]",
		Short: "Navigate group paths",
		Long: `Change active group to given path. The group path can be one of:
  - /:        deactivate current group to set scope to that of user
  - ../:      navigate one level up the group path per repetition of '../'
  - foo/bar:  navigate to the relative path 'foo/bar' from current path
  - /foo/bar: navigate to absolute group path /foo/bar/car`,
		Run: func(cmd *cobra.Command, args []string) {
			target := Config.Targets[Config.CurrentTarget]
			groupPath := target.CurrentGroup

			// Resolve and clean up group path to list
			if len(args) > 0 && len(args[0]) > 0 {
				if args[0][0] == '/' {
					groupPath = args[0]
				} else {
					groupPath = path.Join(groupPath, args[0])
				}
			}
			groupPath = path.Clean(groupPath)

			client := getClient(target)
			if groupPath == "/" {
				// Navigate to root, works always
				target.CurrentGroup = "/"
			} else {
				// Navigate to group, needs checks
				_, _, err := client.Groups.GetGroup(groupPath[1:])
				if err != nil {
					logger.Fatalf(
						"Received error checking group path %s: %v",
						groupPath,
						err,
					)
				}

				target.CurrentGroup = groupPath
			}

			// Update path as necessary
			Config.Targets[Config.CurrentTarget] = target
			saveConfig()

			logger.Printf("Changed current group path to %s\n", groupPath)
		},
	}
)

func init() {
	rootCmd.AddCommand(cdCmd)
}
