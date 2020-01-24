package cfg

import (
	"github.com/xanzy/go-gitlab"
)

// AuthConfig is any kind of configuration that can authenticate to GitLab
type AuthConfig interface {
	Client() (*gitlab.Client, error)
}

// BasicAuthConfig is a Basic Auth-backed AuthConfig
type BasicAuthConfig struct {
	Endpoint string
	Username string
	Password string
}

func (a BasicAuthConfig) Client() (*gitlab.Client, error) {
	client, err := gitlab.NewBasicAuthClient(nil, a.Endpoint, a.Username, a.Password)
	if err != nil {
		err = client.SetBaseURL(a.Endpoint)
		return client, err
	}

	return client, err
}

// TokenAuthConfig is a Token-backed AuthConfig
type TokenAuthConfig struct {
	Endpoint string
	Token    string
}

func (a TokenAuthConfig) Client() (*gitlab.Client, error) {
	client := gitlab.NewClient(nil, a.Token)
	err := client.SetBaseURL(a.Endpoint)
	return client, err
}

// OAuthConfig is an OAuth Token-backed AuthConfig
type OAuthConfig struct {
	Endpoint string
	Token    string
}

func (a OAuthConfig) Client() (*gitlab.Client, error) {
	client := gitlab.NewOAuthClient(nil, a.Token)
	err := client.SetBaseURL(a.Endpoint)
	return client, err
}
