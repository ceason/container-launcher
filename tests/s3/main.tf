resource random_string uniqifier {
  length  = 5
  special = false
  upper   = false
}

data aws_region current {}

resource aws_kms_key key {}

resource aws_s3_bucket config {
  bucket_prefix = "config"
  region        = "${data.aws_region.current.name}"
  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm     = "aws:kms"
        kms_master_key_id = "${aws_kms_key.key.arn}"
      }
    }
  }
}

resource aws_s3_bucket_object custom_message {
  bucket  = "${aws_s3_bucket.config.bucket}"
  key     = "custom-message"
  content = "${local.test_message}"
}

locals {
  test_message            = "This is a custom message, configured in the test workspace! Our unique token for this run is '${random_string.uniqifier.result}'"
  containerlauncher_files = <<EOF
/tmp/custom-message=s3://${aws_s3_bucket.config.bucket}/${aws_s3_bucket_object.custom_message.id}
EOF
  # language=json
  iam_policy              = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [{
    "Resource": "${aws_s3_bucket.config.arn}/*",
    "Effect": "Allow",
    "Action": ["s3:GetObject"]
  },{
    "Resource": ["${aws_kms_key.key.arn}"],
    "Effect": "Allow",
    "Action": ["kms:Decrypt"]
}]
}
EOF
}





