resource random_string uniqifier {
  length  = 5
  special = false
  upper   = false
}

locals {
  test_message                  = "This is a custom message, configured in the test workspace! Our unique token for this run is '${random_string.uniqifier.result}'"
  test_message_file             = "/tmp/test-message-${random_string.uniqifier.result}"

  containerlauncher_environment = <<EOF
CUSTOM_SERVER_MESSAGE_FILE=content:${local.test_message_file}
EOF
  containerlauncher_files       = <<EOF
${local.test_message_file}=content:${local.test_message}
EOF
}





