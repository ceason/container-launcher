load("@io_bazel_rules_docker//container:image.bzl", "container_image")
load("@rules_terraform//terraform:def.bzl", "terraform_integration_test", "terraform_module", "terraform_workspace")
load("@rules_terraform//experimental:k8s.bzl", "terraform_k8s_manifest")
load("@io_bazel_rules_docker//container:container.bzl", "container_push")
load("@io_bazel_rules_docker//python:image.bzl", "py_image")


# One test file for each type of backend (we'll
# do an e2e test of each)
TEST_FILES = [
    "s3.tf",
    "content.tf",
]

# e2e will test:
# - populating ENV vars
# - populating files on disk

py_image(
    name = "server",
    srcs = ["server.py"],
    launcher = "//cmd/container-launcher:container-launcher-linux-amd64",
)

container_push(
    name = "push_server",
    format = "Docker",
    image = ":server",
    registry = "$(IMAGE_REGISTRY)",
    repository = "hello-world-server",
)

terraform_k8s_manifest(
    name = "k8s-deployment",
    srcs = ["server.yaml"],
    deps = [":push_server"],
)

sh_test(
    name = "e2e_test",
    srcs = ["e2e.sh"],
    tags = ["manual"],
)

[terraform_workspace(
    name = f[:-len(".tf")] + "-workspace",
    srcs = [
        "main.tf",
        f,
    ],
    embed = [":k8s-deployment"],
) for f in TEST_FILES]

[terraform_integration_test(
    name = "%s-e2e_integration_test" % f[:-len(".tf")],
    timeout = "short",
    srctest = ":e2e_test",
    tags = ["manual"],
    terraform_workspace = ":%s-workspace" % f[:-len(".tf")],
) for f in TEST_FILES]
