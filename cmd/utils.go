package cmd

import (
	"path"

	"github.com/loksonarius/gli/cfg"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

func resolvePath(currentPath, targetPath string) string {
	result := currentPath

	if len(targetPath) > 0 {
		if targetPath[0] == '/' {
			result = targetPath
		} else {
			result = path.Join(currentPath, targetPath)
		}
	}

	return path.Clean(result)
}

func saveConfig() {
	viper.Set("currenttarget", Config.CurrentTarget)
	viper.Set("targets", Config.Targets)
	if err := viper.WriteConfig(); err != nil {
		logger.Fatalf(
			"Failed to update local config with new target: %v\n",
			err,
		)
	}
}

func verifyTargetName(name string) {
	if _, ok := Config.Targets[name]; !ok {
		logger.Fatalf(
			"%s is not a saved target",
			name,
		)
	}
}

func getClient(target cfg.TargetConfig) *gitlab.Client {
	client, err := target.Auth.Client()
	if err != nil {
		logger.Fatalf(
			"Received following error creating client: %v\n",
			err,
		)
	}

	return client
}
