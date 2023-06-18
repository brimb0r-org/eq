resource "aws_vpc" "mock_vpc" {
  cidr_block = "10.98.188.0/22"
}

module "s3" {
  source             = "./mod/s3"
  aws_account_number = var.aws_account_number
  aws_region         = var.aws_region
  environment        = var.environment
}

module "secrets_manager" {
  source             = "./mod/secrets_manager"
  aws_account_number = var.aws_account_number
  environment        = var.environment
}
