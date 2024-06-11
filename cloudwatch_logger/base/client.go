package base

import (
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

type CloudwatchLogger interface {
	PutMetric(cloudwatch_dimensions []*cloudwatch.Dimension) (*cloudwatch.PutMetricDataOutput, error)
	GetMetric() (*cloudwatch.GetMetricStatisticsOutput, error)
}

type CloudwatchMetric struct {
	MetricName string `json:"metric_name"`
	Namespace  string `json:"namespace"`
}
