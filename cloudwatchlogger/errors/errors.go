package base

type ConfigError string

const (
	UserConfig         ConfigError = "UserConfig"
	OrganizationConfig ConfigError = "OrganizationConfig"
	NetworkConfig      ConfigError = "NetworkConfig"
	Other              ConfigError = "Other"
)
