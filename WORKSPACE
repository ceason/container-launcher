load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

git_repository(
    name = "rules_terraform",
    commit = "840e35e04c2c5be4f733b9aec518181fd8bd8374",
    remote = "git@github.com:ceason/rules_terraform.git",
)

git_repository(
    name = "io_bazel_rules_docker",
    commit = "82b38c6845fd764315eb6fe53d02e85251fd1ed8",
    remote = "git@github.com:ceason/rules_docker.git",
)

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.16.1",
)

git_repository(
    name = "bazel_gazelle",
    remote = "https://github.com/bazelbuild/bazel-gazelle.git",
    tag = "0.15.0",
)

load(
    "@io_bazel_rules_docker//python:image.bzl",
    _py_image_repos = "repositories",
)

_py_image_repos()

load("@io_bazel_rules_go//go:def.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

load("@rules_terraform//terraform:dependencies.bzl", "terraform_repositories")

terraform_repositories()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "com_github_golang_dep",
    importpath = "github.com/golang/dep",
    tag = "v0.5.0",
)
