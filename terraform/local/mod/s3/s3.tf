resource "aws_s3_bucket" "core_app_file_storage" {
  bucket = "${var.environment}-${var.aws_region}-core"
  acl    = "private"

  tags = {
    Name        = "${var.environment}-${var.aws_region}-core"
    Environment = var.environment
  }

  versioning {
    enabled = true
  }

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm = "AES256"
      }
    }
  }
}

resource "aws_s3_bucket_policy" "primary_app_file_storage_bucket_policy" {
  bucket = aws_s3_bucket.core_app_file_storage.id

  policy = <<END_POLICY
{
  "Version": "2012-10-17",
  "Statement":[
    {
      "Sid": "AppAccess",
      "Effect": "Allow",
      "Principal":{
        "AWS": "arn:aws:iam::${var.aws_account_number}:root"
      },
       "Action":[
         "s3:DeleteObject",
         "s3:GetObject",
         "s3:PutObject"
       ],
       "Resource": "arn:aws:s3:::${var.environment}-${var.aws_region}-app/*"
    },
    {
      "Effect": "Deny",
      "Principal": "*",
      "Action": "*",
      "Resource": [
        "arn:aws:s3:::${var.environment}-${var.aws_region}-app",
        "arn:aws:s3:::${var.environment}-${var.aws_region}-app/*"
      ],
      "Condition": {
        "Bool": {
          "aws:SecureTransport": "false"
        }
      }
    }
  ]
}
END_POLICY
}