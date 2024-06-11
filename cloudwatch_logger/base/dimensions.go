package base

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func CreateDimension(name string, value string) *cloudwatch.Dimension {
	return &cloudwatch.Dimension{
		Name:  aws.String(name),
		Value: aws.String(value),
	}
}
