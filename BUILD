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
rm -rf .ijwb/.gopath # dep will break if these intellij project files are present
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