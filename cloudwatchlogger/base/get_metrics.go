package base

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func getMetricInfo(metricName string, namespace string) (*cloudwatch.GetMetricStatisticsOutput, error) {
	// Load the Shared AWS Configuration (~/.aws/config)

	sess := session.Must(session.NewSession())
	svc := cloudwatch.New(sess)

	// Set the time period for the metric data
	endTime := time.Now()
	startTime := endTime.Add(-24 * time.Hour) // Last 24 hours

	statistics := []*string{
		aws.String("SampleCount"),
		aws.String("Sum"),
	}

	// Define the input parameters for GetMetricStatistics
	input := &cloudwatch.GetMetricStatisticsInput{
		Namespace:  aws.String(namespace),
		MetricName: aws.String(metricName),
		StartTime:  aws.Time(startTime),
		EndTime:    aws.Time(endTime),
		Period:     aws.Int64(60), // 1 minute intervals
		Statistics: statistics,
	}

	// Get metric statistics
	result, err := svc.GetMetricStatistics(input)
	if err != nil {
		return nil, fmt.Errorf("unable to get metric statistics, %v", err)
	}

	return result, nil
}

func (t *CloudwatchMetric) GetMetrics() (*cloudwatch.GetMetricStatisticsOutput, error) {

	input, err := getMetricInfo(t.MetricName, t.Namespace)

	// Push the metric to CloudWatch
	if err != nil {
		log.Fatalf("Failed to get metric data: %s", err)
		return nil, err
	}

	return input, nil
}
