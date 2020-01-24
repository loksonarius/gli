package cmd

import (
	"net/url"

	"github.com/loksonarius/gli/cfg"

	"github.com/spf13/cobra"
)

const (
	defaultEndpoint = "https://gitlab.com/"
)

var (
	endpoint string
	authType string
	token    string
	username string
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

			// Should panic out if anything's blatantly wrong
			validateFlags()

			var auth cfg.AuthConfig
			switch authType {
			case "basic":
				auth = cfg.BasicAuthConfig{
					Endpoint: endpoint,
					Username: username,
					Password: password,
				}
			case "oauth":
				auth = cfg.OAuthConfig{
					Endpoint: endpoint,
					Token:    token,
				}
			case "token":
				auth = cfg.TokenAuthConfig{
					Endpoint: endpoint,
					Token:    token,
				}
			}

			client, err := auth.Client()
			if err != nil {
				logger.Fatalf(
					"Received following error creating auth client: %v\n",
					err,
				)
			}

			// Check that credentials actually work with a basic request
			if _, _, err = client.Users.CurrentUserStatus(); err != nil {
				logger.Fatalf(
					"Request failed against user status api: %v\n",
					err,
				)
			}
		},
	}
)

func init() {
	loginCmd.Flags().StringVarP(
		&endpoint,
		"endpoint",
		"e",
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
		&username,
		"username",
		"u",
		"",
		"username to use when auth-type is 'basic'",
	)

	loginCmd.Flags().StringVarP(
		&password,
		"password",
		"p",
		"",
		"password to use when auth-type is 'basic'",
	)

	rootCmd.AddCommand(loginCmd)
}

func validateFlags() {
	_, err := url.Parse(endpoint)
	if err != nil {
		logger.Fatalf(
			"endpoint '%s' is not a valid URL\n",
			endpoint,
		)
	}

	// ensure auth type is a valid type
	// ensure flags match with auth type
	switch authType {
	case "basic":
		if username == "" {
			logger.Fatalln("username must not be empty if using Basic Auth")
		}

		if password == "" {
			logger.Fatalln("password must not be empty if using Basic Auth")
		}

		if token != "" {
			logger.Fatalln("token only used for token-based auth")
		}
	case "oauth":
	case "token":
		if token == "" {
			logger.Fatalln("token must not be empty if using token-based auth")
		}

		if username != "" || password != "" {
			logger.Fatalln("username and password only used for Basic Auth")
		}
	default:
		logger.Fatalln("auth-type must one of 'basic', 'oauth', 'token'")
	}
}
