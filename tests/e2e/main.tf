resource random_string uniqifier {
  length  = 5
  special = false
  upper   = false
}

data kubectl_namespace current {}

locals {
  service_url       = "http://hello-world-server.${data.kubectl_namespace.current.id}.svc.cluster.local"
  test_message      = "This is a custom message, configured in the test workspace! Our unique token for this run is '${random_string.uniqifier.result}'"
  test_message_file = "/tmp/test-message-${random_string.uniqifier.result}"
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




