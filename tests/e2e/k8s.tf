data kubectl_namespace current {}

resource kubernetes_config_map hello_world_server {
  metadata {
    name      = "hello-world-server"
    namespace = "${data.kubectl_namespace.current.id}"
  }

  data {
    CONTAINERLAUNCHER_ENVIRONMENT = "${local.containerlauncher_environment}"
    CONTAINERLAUNCHER_FILES       = "${local.containerlauncher_files}"
  }
}

locals {
  service_url = "http://hello-world-server.${data.kubectl_namespace.current.id}.svc.cluster.local"
}

resource local_file test_vars {
  filename = "test_vars.sh"
  content = <<EOF
NAMESPACE=${data.kubectl_namespace.current.id}
EXPECTED_OUTPUT="${local.test_message}"
SERVICE_URL=${local.service_url}
EOF
}

output service_url {
  value = "${local.service_url}"
}

