package aws

var USGovWest = Region{
	"us-gov-west-1",
	ServiceInfo{"https://ec2.us-gov-west-1.amazonaws.com", V2Signature},
	"https://s3-fips-us-gov-west-1.amazonaws.com",
	"",
	true,
	true,
	"",
	"https://sns.us-gov-west-1.amazonaws.com",
	"https://sqs.us-gov-west-1.amazonaws.com",
	"",
	"https://iam.us-gov.amazonaws.com",
	"https://elasticloadbalancing.us-gov-west-1.amazonaws.com",
	"",
	"https://dynamodb.us-gov-west-1.amazonaws.com",
	ServiceInfo{"https://monitoring.us-gov-west-1.amazonaws.com", V2Signature},
	"https://autoscaling.us-gov-west-1.amazonaws.com",
	ServiceInfo{"https://rds.us-gov-west-1.amazonaws.com", V2Signature},
	"",
	"https://sts.amazonaws.com",
	"https://cloudformation.us-gov-west-1.amazonaws.com",
	"",
}

var USEast = Region{
	"us-east-1",
	ServiceInfo{"https://ec2.us-east-1.amazonaws.com", V2Signature},
	"https://s3-external-1.amazonaws.com",
	"",
	false,
	false,
	"https://sdb.amazonaws.com",
	"https://sns.us-east-1.amazonaws.com",
	"https://sqs.us-east-1.amazonaws.com",
	"https://email.us-east-1.amazonaws.com",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.us-east-1.amazonaws.com",
	"https://kms.us-east-1.amazonaws.com",
	"https://dynamodb.us-east-1.amazonaws.com",
	ServiceInfo{"https://monitoring.us-east-1.amazonaws.com", V2Signature},
	"https://autoscaling.us-east-1.amazonaws.com",
	ServiceInfo{"https://rds.us-east-1.amazonaws.com", V2Signature},
	"https://kinesis.us-east-1.amazonaws.com",
	"https://sts.amazonaws.com",
	"https://cloudformation.us-east-1.amazonaws.com",
	"https://elasticache.us-east-1.amazonaws.com",
}

var USWest = Region{
	"us-west-1",
	ServiceInfo{"https://ec2.us-west-1.amazonaws.com", V2Signature},
	"https://s3-us-west-1.amazonaws.com",
	"",
	true,
	true,
	"https://sdb.us-west-1.amazonaws.com",
	"https://sns.us-west-1.amazonaws.com",
	"https://sqs.us-west-1.amazonaws.com",
	"",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.us-west-1.amazonaws.com",
	"https://kms.us-west-1.amazonaws.com",
	"https://dynamodb.us-west-1.amazonaws.com",
	ServiceInfo{"https://monitoring.us-west-1.amazonaws.com", V2Signature},
	"https://autoscaling.us-west-1.amazonaws.com",
	ServiceInfo{"https://rds.us-west-1.amazonaws.com", V2Signature},
	"https://kinesis.us-west-1.amazonaws.com",
	"https://sts.amazonaws.com",
	"https://cloudformation.us-west-1.amazonaws.com",
	"https://elasticache.us-west-1.amazonaws.com",
}

var USWest2 = Region{
	"us-west-2",
	ServiceInfo{"https://ec2.us-west-2.amazonaws.com", V2Signature},
	"https://s3-us-west-2.amazonaws.com",
	"",
	true,
	true,
	"https://sdb.us-west-2.amazonaws.com",
	"https://sns.us-west-2.amazonaws.com",
	"https://sqs.us-west-2.amazonaws.com",
	"https://email.us-west-2.amazonaws.com",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.us-west-2.amazonaws.com",
	"https://kms.us-west-2.amazonaws.com",
	"https://dynamodb.us-west-2.amazonaws.com",
	ServiceInfo{"https://monitoring.us-west-2.amazonaws.com", V2Signature},
	"https://autoscaling.us-west-2.amazonaws.com",
	ServiceInfo{"https://rds.us-west-2.amazonaws.com", V2Signature},
	"https://kinesis.us-west-2.amazonaws.com",
	"https://sts.amazonaws.com",
	"https://cloudformation.us-west-2.amazonaws.com",
	"https://elasticache.us-west-2.amazonaws.com",
}

var EUWest = Region{
	"eu-west-1",
	ServiceInfo{"https://ec2.eu-west-1.amazonaws.com", V2Signature},
	"https://s3-eu-west-1.amazonaws.com",
	"",
	true,
	true,
	"https://sdb.eu-west-1.amazonaws.com",
	"https://sns.eu-west-1.amazonaws.com",
	"https://sqs.eu-west-1.amazonaws.com",
	"https://email.eu-west-1.amazonaws.com",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.eu-west-1.amazonaws.com",
	"https://kms.eu-west-1.amazonaws.com",
	"https://dynamodb.eu-west-1.amazonaws.com",
	ServiceInfo{"https://monitoring.eu-west-1.amazonaws.com", V2Signature},
	"https://autoscaling.eu-west-1.amazonaws.com",
	ServiceInfo{"https://rds.eu-west-1.amazonaws.com", V2Signature},
	"https://kinesis.eu-west-1.amazonaws.com",
	"https://sts.amazonaws.com",
	"https://cloudformation.eu-west-1.amazonaws.com",
	"https://elasticache.eu-west-1.amazonaws.com",
}

var EUCentral = Region{
	"eu-central-1",
	ServiceInfo{"https://ec2.eu-central-1.amazonaws.com", V4Signature},
	"https://s3-eu-central-1.amazonaws.com",
	"",
	true,
	true,
	"https://sdb.eu-central-1.amazonaws.com",
	"https://sns.eu-central-1.amazonaws.com",
	"https://sqs.eu-central-1.amazonaws.com",
	"",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.eu-central-1.amazonaws.com",
	"https://kms.eu-central-1.amazonaws.com",
	"https://dynamodb.eu-central-1.amazonaws.com",
	ServiceInfo{"https://monitoring.eu-central-1.amazonaws.com", V2Signature},
	"https://autoscaling.eu-central-1.amazonaws.com",
	ServiceInfo{"https://rds.eu-central-1.amazonaws.com", V2Signature},
	"https://kinesis.eu-central-1.amazonaws.com",
	"https://sts.amazonaws.com",
	"https://cloudformation.eu-central-1.amazonaws.com",
	"",
}

var APSoutheast = Region{
	"ap-southeast-1",
	ServiceInfo{"https://ec2.ap-southeast-1.amazonaws.com", V2Signature},
	"https://s3-ap-southeast-1.amazonaws.com",
	"",
	true,
	true,
	"https://sdb.ap-southeast-1.amazonaws.com",
	"https://sns.ap-southeast-1.amazonaws.com",
	"https://sqs.ap-southeast-1.amazonaws.com",
	"",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.ap-southeast-1.amazonaws.com",
	"https://kms.ap-southeast-1.amazonaws.com",
	"https://dynamodb.ap-southeast-1.amazonaws.com",
	ServiceInfo{"https://monitoring.ap-southeast-1.amazonaws.com", V2Signature},
	"https://autoscaling.ap-southeast-1.amazonaws.com",
	ServiceInfo{"https://rds.ap-southeast-1.amazonaws.com", V2Signature},
	"https://kinesis.ap-southeast-1.amazonaws.com",
	"https://sts.amazonaws.com",
	"https://cloudformation.ap-southeast-1.amazonaws.com",
	"https://elasticache.ap-southeast-1.amazonaws.com",
}

var APSoutheast2 = Region{
	"ap-southeast-2",
	ServiceInfo{"https://ec2.ap-southeast-2.amazonaws.com", V2Signature},
	"https://s3-ap-southeast-2.amazonaws.com",
	"",
	true,
	true,
	"https://sdb.ap-southeast-2.amazonaws.com",
	"https://sns.ap-southeast-2.amazonaws.com",
	"https://sqs.ap-southeast-2.amazonaws.com",
	"",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.ap-southeast-2.amazonaws.com",
	"https://kms.ap-southeast-2.amazonaws.com",
	"https://dynamodb.ap-southeast-2.amazonaws.com",
	ServiceInfo{"https://monitoring.ap-southeast-2.amazonaws.com", V2Signature},
	"https://autoscaling.ap-southeast-2.amazonaws.com",
	ServiceInfo{"https://rds.ap-southeast-2.amazonaws.com", V2Signature},
	"https://kinesis.ap-southeast-2.amazonaws.com",
	"https://sts.amazonaws.com",
	"https://cloudformation.ap-southeast-2.amazonaws.com",
	"https://elasticache.ap-southeast-2.amazonaws.com",
}

var APNortheast = Region{
	"ap-northeast-1",
	ServiceInfo{"https://ec2.ap-northeast-1.amazonaws.com", V2Signature},
	"https://s3-ap-northeast-1.amazonaws.com",
	"",
	true,
	true,
	"https://sdb.ap-northeast-1.amazonaws.com",
	"https://sns.ap-northeast-1.amazonaws.com",
	"https://sqs.ap-northeast-1.amazonaws.com",
	"",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.ap-northeast-1.amazonaws.com",
	"https://kms.ap-northeast-1.amazonaws.com",
	"https://dynamodb.ap-northeast-1.amazonaws.com",
	ServiceInfo{"https://monitoring.ap-northeast-1.amazonaws.com", V2Signature},
	"https://autoscaling.ap-northeast-1.amazonaws.com",
	ServiceInfo{"https://rds.ap-northeast-1.amazonaws.com", V2Signature},
	"https://kinesis.ap-northeast-1.amazonaws.com",
	"https://sts.amazonaws.com",
	"https://cloudformation.ap-northeast-1.amazonaws.com",
	"https://elasticache.ap-northeast-1.amazonaws.com",
}

var SAEast = Region{
	"sa-east-1",
	ServiceInfo{"https://ec2.sa-east-1.amazonaws.com", V2Signature},
	"https://s3-sa-east-1.amazonaws.com",
	"",
	true,
	true,
	"https://sdb.sa-east-1.amazonaws.com",
	"https://sns.sa-east-1.amazonaws.com",
	"https://sqs.sa-east-1.amazonaws.com",
	"",
	"https://iam.amazonaws.com",
	"https://elasticloadbalancing.sa-east-1.amazonaws.com",
	"https://kms.sa-east-1.amazonaws.com",
	"https://dynamodb.sa-east-1.amazonaws.com",
	ServiceInfo{"https://monitoring.sa-east-1.amazonaws.com", V2Signature},
	"https://autoscaling.sa-east-1.amazonaws.com",
	ServiceInfo{"https://rds.sa-east-1.amazonaws.com", V2Signature},
	"",
	"https://sts.amazonaws.com",
	"https://cloudformation.sa-east-1.amazonaws.com",
	"https://elasticache.sa-east-1.amazonaws.com",
}

var CNNorth1 = Region{
	"cn-north-1",
	ServiceInfo{"https://ec2.cn-north-1.amazonaws.com.cn", V2Signature},
	"https://s3.cn-north-1.amazonaws.com.cn",
	"",
	true,
	true,
	"",
	"https://sns.cn-north-1.amazonaws.com.cn",
	"https://sqs.cn-north-1.amazonaws.com.cn",
	"",
	"https://iam.cn-north-1.amazonaws.com.cn",
	"https://elasticloadbalancing.cn-north-1.amazonaws.com.cn",
	"",
	"https://dynamodb.cn-north-1.amazonaws.com.cn",
	ServiceInfo{"https://monitoring.cn-north-1.amazonaws.com.cn", V4Signature},
	"https://autoscaling.cn-north-1.amazonaws.com.cn",
	ServiceInfo{"https://rds.cn-north-1.amazonaws.com.cn", V4Signature},
	"",
	"https://sts.cn-north-1.amazonaws.com.cn",
	"",
	"",
}
var Exoscale = Region{
	"exoscale",
	ServiceInfo{"https://sos.exo.io", V2Signature},
	"https://sos.exo.io",
	"",
	true,
	true,
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	"",
	ServiceInfo{"https://monitoring.cn-north-1.amazonaws.com.cn", V4Signature},
	"",
	ServiceInfo{"https://rds.cn-north-1.amazonaws.com.cn", V4Signature},
	"",
	"",
	"",
	"",
}
