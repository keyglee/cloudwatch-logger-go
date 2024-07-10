package metrics

type Metrics string

const (
	Misconfiguration       Metrics = "MisconfigurationError"
	UndefinedConfiguration Metrics = "MissingConfigurationError"
	Unauthorized           Metrics = "UnauthorizedError"
)
