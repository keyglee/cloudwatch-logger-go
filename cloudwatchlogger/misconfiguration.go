package cloudwatchlogger

import (
	"os"

	"github.com/keyglee/cloudwatch-logger-go/cloudwatchlogger/base"
	"github.com/keyglee/cloudwatch-logger-go/cloudwatchlogger/errors"
	"github.com/keyglee/cloudwatch-logger-go/cloudwatchlogger/metrics"

	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func LogMetric(metric_name metrics.Metrics, namespace string, accessed_resource string, config_error errors.ConfigError) (*cloudwatch.PutMetricDataOutput, error) {

	client := base.CloudwatchMetric{MetricName: string(metric_name), Namespace: namespace}

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
