# Cloudwatch Logger

Holds clients for publishing custom metrics to cloudwatch as an extension of logger

# Before you begin
Metrics in cloudwatch get created as you put new metric data in
so running ```aws cloudwatch put-metric-data --metric-name PageViewCount --namespace MyService --value 2 --timestamp 2016-10-20T12:00:00.000Z``` will automatically create the metric to filter by in cloudwatch

# Adding metrics
For each language we utilize (python, golang, typescript, cobol, etc) we will need to create a client with the operations and metric names to hit cloudwatch with

# Permissions

in AWS CDK, a policy is required to be created that allows cloudwatch to publish metrics for the namespace and metric name you are using.  This policy is created in the cdk.json file under `app: 'aws-cdk-lib/aws-iam'`.

Your files may look something similar to this:
```ts 
// Access Role
const iam = require('aws-cdk-lib/aws-iam');

const accessRole = new iam.Role(stack, 'AccessRole', {
    assumedBy: new iam.ServicePrincipal('lambda.amazonaws.com'),
    managedPolicies: [
        iam.ManagedPolicy.fromAwsManagedPolicyName('service-role/AWSLambdaBasicExecutionRole')
    ]
});

accessRole.addToPolicy(new iam.PolicyStatement({
    effect: iam.Effect.ALLOW,
    actions: ['cloudwatch:PutMetricData'],
    resources: ['*']
}));
```

```ts 
// Lambda
const sqsProcessorLambda = createCompiledFunction(stack, {
    name: "SomeImportantLambda",
    location: "bin/functions/bigLambda.go",
    role: accessRole,
    props: {
        ...baseProps,
    },
    environment: {
        ...env
    }
});
```