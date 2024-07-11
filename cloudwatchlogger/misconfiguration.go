package cloudwatchlogger

import (
	"os"

	"github.com/keyglee/cloudwatch-logger-go/cloudwatchlogger/base"
	"github.com/keyglee/cloudwatch-logger-go/cloudwatchlogger/dimensions"
	"github.com/keyglee/cloudwatch-logger-go/cloudwatchlogger/errors"
	"github.com/keyglee/cloudwatch-logger-go/cloudwatchlogger/metrics"

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
func LogMetric(metric_name metrics.Metrics, namespace string, accessed_resource string, config_error errors.ConfigError, extra_dimensions []*cloudwatch.Dimension) (*cloudwatch.PutMetricDataOutput, error) {

	client := base.CloudwatchMetric{MetricName: string(metric_name), Namespace: namespace}

	var dimensionList []*cloudwatch.Dimension

	dimensionList = append(dimensionList, dimensions.CreateDimension("Error", string(config_error)))
	dimensionList = append(dimensionList, dimensions.CreateDimension("Resource", accessed_resource))

	// AWS Reserved environment variables
	dimensionList = append(dimensionList, dimensions.CreateDimension("FunctionName", os.Getenv("AWS_LAMBDA_FUNCTION_NAME")))
	dimensionList = append(dimensionList, dimensions.CreateDimension("AWS_LAMBDA_LOG_GROUP_NAME", os.Getenv("AWS_LAMBDA_LOG_GROUP_NAME")))
	dimensionList = append(dimensionList, dimensions.CreateDimension("AWS_LAMBDA_LOG_STREAM_NAME", os.Getenv("AWS_LAMBDA_LOG_STREAM_NAME")))

	dimensionList = append(dimensionList, extra_dimensions...)

	out, err := client.PutMetric(dimensionList)

	return out, err
}
