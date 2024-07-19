package cloudwatchlogger

import (
	"os"

	"github.com/keyglee/cloudwatch-logger-go/cloudwatchlogger/base"
	"github.com/keyglee/cloudwatch-logger-go/cloudwatchlogger/dimensions"
	"github.com/keyglee/cloudwatch-logger-go/cloudwatchlogger/metrics"
	"github.com/keyglee/cloudwatch-logger-go/cloudwatchlogger/namespaces"

	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

// LogMetric - Logs a metric to cloudwatch
//
// @param metric_name - Metric name to be logged
//
// @param namespace - Namespace of the metric being logged
//
// @param accessed_resource - Resource that is accessed by the metric being logged
//
// @param config_error - Error that occurred while accessing the resource. This is used for logging.
//
// @param extra_dimensions - Extra dimensions to be added to the metric being logged
//
// returns - cloudwatch.PutMetricDataOutput, error
func LogMetric(metric_name metrics.Metrics, namespace namespaces.Namespace, extra_dimensions []*cloudwatch.Dimension) (*cloudwatch.PutMetricDataOutput, error) {

	client := base.CloudwatchMetric{MetricName: string(metric_name), Namespace: string(namespace)}

	// AWS Reserved environment variables

	out, err := client.PutMetric(extra_dimensions)

	return out, err
}

func AddFunctionName(dimensionsList []*cloudwatch.Dimension) []*cloudwatch.Dimension {
	returnDimensions := dimensionsList
	if dimensionsList == nil {
		returnDimensions = []*cloudwatch.Dimension{}
	}

	returnDimensions = append(returnDimensions, dimensions.CreateDimension("FunctionName", os.Getenv("AWS_LAMBDA_FUNCTION_NAME")))

	return returnDimensions
}
