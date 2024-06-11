package cloudwatchlogger

import (
	"cloudwatch-logger/cloudwatch_logger/base"
	"errors"
	"os"

	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

type ConfigError string

const (
	UserConfig         ConfigError = "UserConfig"
	OrganizationConfig ConfigError = "OrganizationConfig"
	NetworkConfig      ConfigError = "NetworkConfig"
	Other              ConfigError = "Other"
)

func ValidateStatus(status ConfigError) error {
	switch status {
	case UserConfig, OrganizationConfig, NetworkConfig, Other:
		return nil
	default:
		return errors.New("invalid status")
	}
}

func PushMisconfigurationError(accessed_resource string, config_error ConfigError) (*cloudwatch.PutMetricDataOutput, error) {

	const METRIC_NAME = "MisconfigurationError"

	client := base.CloudwatchMetric{MetricName: METRIC_NAME, Namespace: "TechSupport"}

	var dimensions []*cloudwatch.Dimension

	dimensions = append(dimensions, base.CreateDimension("Error", string(config_error)))
	dimensions = append(dimensions, base.CreateDimension("Service", os.Getenv("SERVICE_NAME")))
	dimensions = append(dimensions, base.CreateDimension("Resource", accessed_resource))
	dimensions = append(dimensions, base.CreateDimension("FunctionName", os.Getenv("AWS_LAMBDA_FUNCTION_NAME")))
	dimensions = append(dimensions, base.CreateDimension("StackName", os.Getenv("STACK_NAME")))
	dimensions = append(dimensions, base.CreateDimension("STAGE", os.Getenv("STAGE")))

	out, err := client.PutMetric(dimensions)

	return out, err
}

func main() {
	PushMisconfigurationError("Something", "Something")
}
