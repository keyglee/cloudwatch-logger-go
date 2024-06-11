package cloudwatchlogger

import (
	"cloudwatch-logger/cloudwatch_logger/base"
	"os"

	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func PushAuthorizedAcess(accessed_resource string) (*cloudwatch.PutMetricDataOutput, error) {
	const METRIC_NAME = "AuthorizedAccessException"

	client := base.CloudwatchMetric{MetricName: METRIC_NAME, Namespace: "General"}

	var dimensions []*cloudwatch.Dimension

	dimensions = append(dimensions, base.CreateDimension("Service", os.Getenv("SERVICE_NAME")))
	dimensions = append(dimensions, base.CreateDimension("Resource", accessed_resource))
	dimensions = append(dimensions, base.CreateDimension("FunctionName", os.Getenv("AWS_LAMBDA_FUNCTION_NAME")))
	dimensions = append(dimensions, base.CreateDimension("StackName", os.Getenv("STACK_NAME")))
	dimensions = append(dimensions, base.CreateDimension("STAGE", os.Getenv("STAGE")))

	out, err := client.PutMetric(dimensions)

	return out, err
}
