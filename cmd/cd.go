package cmd

import (
	"github.com/spf13/cobra"
)

var (
	cdCmd = &cobra.Command{
		Use:   "cd [path]",
		Short: "Navigate group paths",
		Long: `Change active group to given path. The group path can be one of:
  - /:        deactivate current group to set scope to that of user
  - [no arg]: same as navigating to /
  - ../:      navigate one level up the group path per repetition of '../'
  - foo/bar:  navigate to the relative path 'foo/bar' from current path
  - /foo/bar: navigate to absolute group path /foo/bar/car`,
		Run: func(cmd *cobra.Command, args []string) {
			target := Config.Targets[Config.CurrentTarget]
			path := target.CurrentGroup
			if len(args) > 0 {
				path = resolvePath(target.CurrentGroup, args[0])
			}

			client := getClient(target)
			if path == "/" {
				// Navigate to root, works always
				target.CurrentGroup = "/"
			} else {
				// Navigate to group, needs checks
				_, _, err := client.Groups.GetGroup(path[1:])
				if err != nil {
					logger.Fatalf(
						"Received error checking group path %s: %v",
						path,
						err,
					)
				}

				target.CurrentGroup = path
			}

			// Update path as necessary
			Config.Targets[Config.CurrentTarget] = target
			saveConfig()

			logger.Printf("Changed current group path to %s\n", path)
		},
	}
)

func init() {
	rootCmd.AddCommand(cdCmd)
}
