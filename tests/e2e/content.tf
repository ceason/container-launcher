
resource kubernetes_config_map hello_world_server {
  metadata {
    name      = "hello-world-server"
    namespace = "${data.kubectl_namespace.current.id}"
  }

  data {
    CUSTOM_SERVER_MESSAGE_FILE = "containerlauncher:${local.test_message_file}:content:${local.test_message}"
  }
}