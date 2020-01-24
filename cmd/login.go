package cmd

import (
	"github.com/spf13/cobra"
)

const (
	defaultEndpoint = "https://gitlab.com/"
)

var (
	endpoint string
	authType string
	token    string
	user     string
	password string

	loginCmd = &cobra.Command{
		Use:   "login [name]",
		Short: "Log in to a GitLab instance and save it as a target",
		Long: `Login creates a new 'name' target to run commands against. Login
supports Basic Auth, Token, and OAuth Token authentication against a
given endpoint. If using a token-based authentication, use the -e and -t
flags. If using Basic Auth, then please use the -u, -p, and -e flags.
Upon a successful login, the new 'name' target will be stored internally
along with the given credentials. List these using 'gli target list'.`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			logger.Println("Unimplemented")
		},
	}
)

func init() {
	loginCmd.Flags().StringVarP(
		&endpoint,
		"endpoint",
		"s",
		defaultEndpoint,
		"gitLab endpoint to auth against",
	)

	loginCmd.Flags().StringVarP(
		&authType,
		"auth-type",
		"a",
		"token",
		"auth-type to authenticate with; one of 'token', 'basic', or 'oauth'",
	)

	loginCmd.Flags().StringVarP(
		&token,
		"token",
		"t",
		"",
		"token to use when auth-type is 'token' or 'oauth'",
	)

	loginCmd.Flags().StringVarP(
		&user,
		"user",
		"u",
		"",
		"username to use when auth-type is 'basic'",
	)

	loginCmd.Flags().StringVarP(
		&user,
		"pass",
		"p",
		"",
		"password to use when auth-type is 'basic'",
	)

	rootCmd.AddCommand(loginCmd)
}
