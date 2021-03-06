package cmd

import (
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

var (
	lsCmd = &cobra.Command{
		Use:   "ls [path]",
		Short: "List resources and groups for a path",
		Long: `Groups are treated like a namespacing tool throughout this tool.
An easy way to understand the metaphor employed is to think of a Group's
path as filesystem path and all the projects, issues, and such resources
as files. Most operations that involve listing or searching will also
look through sub-groups as well. The root path ('/') lines up with your
user's profile, and any subcommands run from there assume you intend to
scope queries and actions to your own personal projects, issues, and
such. Running ls from the root path will list top-level groups your
user has access to`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			target := Config.Targets[Config.CurrentTarget]
			path := target.CurrentGroup
			if len(args) > 0 {
				path = resolvePath(target.CurrentGroup, args[0])
			}

			t.AppendHeader(table.Row{
				"Type", "Name", "ID", "Description",
			})

			client := getClient(target)
			if path == "/" {
				// We're listing user resources
				user, _, err := client.Users.CurrentUser()
				if err != nil {
					logger.Fatalf(
						"Received error getting current user: %v",
						err,
					)
				}

				groups, _, err := client.Groups.ListGroups(nil)
				if err != nil {
					logger.Fatalf(
						"Received error listing groups: %v",
						err,
					)
				}

				for _, group := range groups {
					if group.ParentID == 0 {
						t.AppendRow(table.Row{
							"group", group.Path, group.ID, group.Description,
						})
					}
				}

				projects, _, err := client.Projects.ListUserProjects(user.ID, nil)
				if err != nil {
					logger.Fatalf(
						"Received error listing projects: %v",
						err,
					)
				}

				for _, project := range projects {
					t.AppendRow(table.Row{
						"project", project.Path, project.ID, project.Description,
					})
				}
			} else {
				groups, _, err := client.Groups.ListSubgroups(path[1:], nil)
				if err != nil {
					logger.Fatalf(
						"Received error listing groups: %v",
						err,
					)
				}

				for _, group := range groups {
					t.AppendRow(table.Row{
						"group", group.Path, group.ID, group.Description,
					})
				}

				projects, _, err := client.Groups.ListGroupProjects(path[1:], nil)
				if err != nil {
					logger.Fatalf(
						"Received error listing projects: %v",
						err,
					)
				}

				for _, project := range projects {
					t.AppendRow(table.Row{
						"project", project.Path, project.ID, project.Description,
					})
				}
			}

			t.SortBy([]table.SortBy{
				table.SortBy{Name: "Type", Mode: table.Asc},
				table.SortBy{Name: "Name", Mode: table.Asc},
			})
			logger.Println(t.Render())
		},
	}
)

func init() {
	rootCmd.AddCommand(lsCmd)
}
