package cmd

import (
	"fmt"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/spf13/cobra"
)

var (
	statusCmd = &cobra.Command{
		Use:   "status",
		Short: "Get overview of things",
		Long:  `Print active target, group, connectivity, and list To Do's.`,
		Run: func(cmd *cobra.Command, args []string) {
			target := Config.Targets[Config.CurrentTarget]
			group := target.CurrentGroup
			connected := true

			client := getClient(target)
			user, _, err := client.Users.CurrentUser()
			if err != nil {
				connected = false
			}

			connection := text.Faint.Sprint("[offline]")
			name := text.Bold.Sprint("unknown")
			targetName := text.Bold.Sprint(Config.CurrentTarget)
			groupPath := text.Bold.Sprint(group)
			if connected {
				connection = text.Faint.Sprint("[online]")
				name = text.Bold.Sprint(user.Name)
			}
			brief := fmt.Sprintf("%s %s @ %s in %s", connection, name, targetName, groupPath)

			todoList := ""
			if connected {
				todos, _, err := client.Todos.ListTodos(nil)
				if err != nil {
					logger.Fatalf(
						"Received error listing user's To Do's: %v",
						err,
					)
				}

				if len(todos) > 0 {
					t.AppendHeader(table.Row{"Type", "ID", "Project", "Body"})

					for _, todo := range todos {
						t.AppendRow(table.Row{
							todo.TargetType, todo.ID, todo.Project.PathWithNamespace, todo.Body,
						})
					}

					todoList = fmt.Sprintf("%s\n%s",
						text.Bold.Sprint("To Do's"),
						t.Render(),
					)
				} else {
					todoList = "No To Do's Pending!"
				}
			}

			logger.Println(brief)
			logger.Println(todoList)
		},
	}
)

func init() {
	rootCmd.AddCommand(statusCmd)
}
