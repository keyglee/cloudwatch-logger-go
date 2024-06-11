# Cloudwatch Logger

Holds clients for publishing custom metrics to cloudwatch as an extension of logger

# Before you begin
Metrics in cloudwatch get created as you put new metric data in
so running ```aws cloudwatch put-metric-data --metric-name PageViewCount --namespace MyService --value 2 --timestamp 2016-10-20T12:00:00.000Z``` will automatically create the metric to filter by in cloudwatch

# Adding metrics
For each language we utilize (python, golang, typescript, cobol, etc) we will need to create a client with the operations and metric names to hit cloudwatch with