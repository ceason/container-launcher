load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

git_repository(
    name = "rules_terraform",
    commit = "7bb4a94fa6f89493a5162511dafabbf4df64a724",
    remote = "git@github.com:ceason/rules_terraform.git",
)

git_repository(
    name = "io_bazel_rules_docker",
    commit = "f0cd1fccefaaad3b63781b08230649687c8a2261",
    remote = "git@github.com:ceason/rules_docker.git",
)

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.16.3",
)

git_repository(
    name = "bazel_gazelle",
    commit = "44ce230b3399a5d4472198740358fcd825b0c3c9",
    remote = "https://github.com/bazelbuild/bazel-gazelle.git",
)

load("@io_bazel_rules_go//go:def.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

load("@rules_terraform//terraform:dependencies.bzl", "terraform_repositories")

terraform_repositories()

load(
    "@io_bazel_rules_docker//python:image.bzl",
    _py_image_repos = "repositories",
)

_py_image_repos()

go_repository(
    name = "com_github_golang_dep",
    importpath = "github.com/golang/dep",
    tag = "v0.5.0",
)
