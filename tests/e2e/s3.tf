resource kubernetes_config_map hello_world_server {
  metadata {
    name      = "hello-world-server"
    namespace = "${data.kubectl_namespace.current.id}"
  }
  data {
    CUSTOM_SERVER_MESSAGE_FILE = "containerlauncher:${local.test_message_file}:s3://${aws_s3_bucket.config.bucket}/${aws_s3_bucket_object.custom_message.id}"
    AWS_ACCESS_KEY_ID          = "${aws_iam_access_key.testuser.id}"
    AWS_SECRET_ACCESS_KEY      = "${aws_iam_access_key.testuser.secret}"
    AWS_REGION                 = "${data.aws_region.current.name}"
  }
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

resource aws_iam_user testuser {
  name          = "testuser-${random_string.uniqifier.result}"
  force_destroy = true
}
resource aws_iam_access_key testuser {
  user = "${aws_iam_user.testuser.name}"
}
resource aws_iam_user_policy testuser {
  user   = "${aws_iam_user.testuser.name}"
  # language=json
  policy = <<EOF
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



