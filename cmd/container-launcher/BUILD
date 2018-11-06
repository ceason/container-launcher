load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//container:layer.bzl", "container_layer")

container_layer(
    name = "layer",
    directory = "/",
    files = [
        ":container-launcher",
    ],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/ceason/container-launcher/cmd/container-launcher",
    visibility = ["//visibility:private"],
    deps = [
        "//:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
    ],
)

go_binary(
    name = "container-launcher",
    embed = [":go_default_library"],
    goarch = "amd64",
    goos = "linux",
    pure = "on",
    visibility = ["//visibility:public"],
)
