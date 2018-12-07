load("@rules_terraform//experimental:publishing.bzl", "ghrelease_assets", "ghrelease_publisher", "ghrelease_test_suite")
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

VERSION = "0.1"

ghrelease_publisher(
    name = "release",
    args = ["--draft"],
    version = VERSION,
    deps = [
        ":assets",
        ":prerelease-tests",
    ],
)

ghrelease_test_suite(
    name = "prerelease-tests",
    tests = [
        "//...",
        "//tests/e2e:content-e2e_integration_test",
        "//tests/e2e:s3-e2e_integration_test",
    ],
)

ghrelease_assets(
    name = "assets",
    data = ["//cmd/container-launcher:container-launcher-linux-amd64"],
)
