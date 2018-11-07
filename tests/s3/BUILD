load("@io_bazel_rules_docker//container:image.bzl", "container_image")
load("@rules_terraform//terraform:def.bzl", "terraform_integration_test", "terraform_module", "terraform_workspace")
load("@rules_terraform//terraform:container.bzl", "image_embedder", "terraform_k8s_manifest")
load("@io_bazel_rules_docker//python:image.bzl", "py_image")

# e2e will test:
# - populating ENV vars
# - populating files on disk

py_image(
    name = "server",
    srcs = ["server.py"],
    launcher = "//cmd/container-launcher",
)

terraform_k8s_manifest(
    name = "k8s-deployment",
    srcs = ["server.yaml"],
    image_chroot = "$(IMAGE_CHROOT)",
    images = {
        "hello-world-server:dev": ":server",
    },
)

terraform_workspace(
    name = "k8s-workspace",
    srcs = [
        "main.tf",
        "k8s.tf",
    ],
    embed = [":k8s-deployment"],
)

sh_test(
    name = "e2e_test",
    srcs = ["test.sh"],
    tags = ["manual"],
)

terraform_integration_test(
    name = "k8s-e2e_integration_test",
    timeout = "short",
    srctest = ":e2e_test",
    #tags = ["manual"],
    terraform_workspace = ":k8s-workspace",
)
