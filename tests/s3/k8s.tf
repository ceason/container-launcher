data kubectl_namespace current {}

resource aws_iam_user testuser {
  name          = "testuser-${random_string.uniqifier.result}"
  force_destroy = true
}
resource aws_iam_access_key testuser {
  user = "${aws_iam_user.testuser.name}"
}
resource aws_iam_user_policy testuser {
  user   = "${aws_iam_user.testuser.name}"
  policy = "${local.iam_policy}"
}

resource kubernetes_config_map hello_world_server {
  metadata {
    name      = "hello-world-server"
    namespace = "${data.kubectl_namespace.current.id}"
  }
  data {
    CONTAINERLAUNCHER_FILES = "${local.containerlauncher_files}"
    AWS_ACCESS_KEY_ID       = "${aws_iam_access_key.testuser.id}"
    AWS_SECRET_ACCESS_KEY   = "${aws_iam_access_key.testuser.secret}"
    AWS_REGION              = "${data.aws_region.current.name}"
  }
}

locals {
  service_url = "http://hello-world-server.${data.kubectl_namespace.current.id}.svc.cluster.local"
}

resource local_file test_vars {
  filename = "test_vars.sh"
  content  = <<EOF
NAMESPACE=${data.kubectl_namespace.current.id}
EXPECTED_OUTPUT="${local.test_message}"
SERVICE_URL=${local.service_url}
EOF
}

output service_url {
  value = "${local.service_url}"
}

