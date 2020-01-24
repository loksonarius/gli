package cfg

// Config represents expected structure configuring the CLI
type Config struct {
	CurrentTarget string
	Targets       map[string]TargetConfig
}

// TargetConfig represents an authenticated GitLab endpoint and the current
// group being focused on
type TargetConfig struct {
	CurrentGroup string
	Auth         AuthConfig
}
