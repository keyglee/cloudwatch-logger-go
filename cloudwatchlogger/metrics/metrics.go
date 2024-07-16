package metrics

type Metrics string

const (
	Misconfiguration       Metrics = "Misconfiguration"
	UndefinedConfiguration Metrics = "MissingConfiguration"
	Unauthorized           Metrics = "Unauthorized"
	Network                Metrics = "Network"
	Billing                Metrics = "Billing"
)
