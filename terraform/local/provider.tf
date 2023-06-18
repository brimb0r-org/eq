# this config struct is for local only

terraform {
  backend "local" {}
  required_providers {
    null = {
      source  = "hashicorp/null"
      version = "2.1"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.29.1"
    }
  }
}

provider "aws" {
  access_key                  = "mock_access_key"
  region                      = var.aws_region
  s3_force_path_style         = true
  secret_key                  = "mock_secret_key"
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true

  endpoints {
    apigateway       = "http://${var.aws_endpoint}:4566"
    cloudwatch       = "http://${var.aws_endpoint}:4566"
    cloudwatchlogs   = "http://${var.aws_endpoint}:4566"
    cloudwatchevents = "http://${var.aws_endpoint}:4566"
    dynamodb         = "http://${var.aws_endpoint}:4566"
    es               = "http://${var.aws_endpoint}:4566"
    ec2              = "http://${var.aws_endpoint}:4566"
    iam              = "http://${var.aws_endpoint}:4566"
    route53          = "http://${var.aws_endpoint}:4566"
    s3               = "http://${var.aws_endpoint}:4566"
    sqs              = "http://${var.aws_endpoint}:4566"
    secretsmanager   = "http://${var.aws_endpoint}:4566"
    sts              = "http://${var.aws_endpoint}:4566"
  }
}