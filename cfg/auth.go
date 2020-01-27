package cfg

import (
	"fmt"

	"github.com/xanzy/go-gitlab"
)

// AuthConfig is any kind of configuration that can authenticate to GitLab
type AuthConfig struct {
	Endpoint string
	Type     string
	Token    string
	Username string
	Password string
}

func (a AuthConfig) Client() (*gitlab.Client, error) {
	var client *gitlab.Client
	var err error

	switch a.Type {
	case "basic":
		client, err = gitlab.NewBasicAuthClient(nil, a.Endpoint, a.Username, a.Password)
		if err != nil {
			err = client.SetBaseURL(a.Endpoint)
		}
	case "oauth":
		client = gitlab.NewOAuthClient(nil, a.Token)
		err = client.SetBaseURL(a.Endpoint)
	case "token":
		client = gitlab.NewClient(nil, a.Token)
		err = client.SetBaseURL(a.Endpoint)
	default:
		client, err = nil, fmt.Errorf("%v is not a vaild AuthType", a.Type)
	}

	return client, err
}
