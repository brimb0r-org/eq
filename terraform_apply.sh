#!/usr/bin/env sh
terraform -chdir=/src/terraform init
terraform -chdir=/src/terraform plan -var aws_endpoint=localstack
terraform -chdir=/src/terraform apply -var aws_endpoint=localstack -auto-approve

# Remove all the terraform resources after the fact
rm -rf /src/terraform/.terraform/modules
rm -f /src/terraform/.terraform/terraform.tfstate
rm -f /src/terraform/.terraform.lock.hcl
rm -f /src/terraform/terraform.tfstate
rm -f /src/terraform/terraform.tfstate.backup