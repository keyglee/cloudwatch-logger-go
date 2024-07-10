package base

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func (t *CloudwatchMetric) PutMetric(cloudwatch_dimensions []*cloudwatch.Dimension) (*cloudwatch.PutMetricDataOutput, error) {

	sess := session.Must(session.NewSession())
	svc := cloudwatch.New(sess)

	metricData := &cloudwatch.MetricDatum{
		MetricName: aws.String(t.MetricName),
		Dimensions: cloudwatch_dimensions,
		Timestamp:  aws.Time(time.Now()),
		Unit:       aws.String(cloudwatch.StandardUnitCount),
		Value:      aws.Float64(1.0),
	}

	input := &cloudwatch.PutMetricDataInput{
		Namespace:  aws.String(t.Namespace),
		MetricData: []*cloudwatch.MetricDatum{metricData},
	}

	// Push the metric to CloudWatch
	out, err := svc.PutMetricData(input)
	if err != nil {
		log.Fatalf("Failed to put metric data: %s", err)
		return nil, err
	}

	fmt.Println("Metric pushed to CloudWatch successfully")

	return out, nil
}
