
resource kubernetes_config_map hello_world_server {
  metadata {
    name      = "hello-world-server"
    namespace = "${data.kubectl_namespace.current.id}"
  }

  data {
    CONTAINERLAUNCHER_ENVIRONMENT = <<EOF
CUSTOM_SERVER_MESSAGE_FILE=content:${local.test_message_file}
EOF
    CONTAINERLAUNCHER_FILES       = <<EOF
${local.test_message_file}=content:${local.test_message}
EOF
  }
}