load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/ceason/container-launcher

gazelle(name = "gazelle")

genrule(
    name = "dep-ensure",
    outs = ["dep-ensure.sh"],
    cmd = """cat <<'EOF' > $@
#!/usr/bin/env bash
set -euo pipefail
cd "$$BUILD_WORKSPACE_DIRECTORY"
$(location @com_github_golang_dep//cmd/dep) ensure
$(location @bazel_gazelle//cmd/gazelle) update
git add Gopkg.{lock,toml} vendor/. $$(find . -name BUILD -o -name BUILD.bazel)
EOF""",
    executable = True,
    tools = [
        "@bazel_gazelle//cmd/gazelle",
        "@com_github_golang_dep//cmd/dep",
    ],
)

go_library(
    name = "go_default_library",
    srcs = glob(
        ["*.go"],
        exclude = ["*_test.go"],
    ),
    importpath = "github.com/ceason/container-launcher",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/aws/aws-sdk-go/aws:go_default_library",
        "//vendor/github.com/aws/aws-sdk-go/aws/session:go_default_library",
        "//vendor/github.com/aws/aws-sdk-go/service/s3:go_default_library",
        "//vendor/github.com/aws/aws-sdk-go/service/s3/s3manager:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
    ],
)
