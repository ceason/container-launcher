load("//lang:launcher.bzl", "launcher_image")
load("@io_bazel_rules_docker//python:image.bzl", _py_image = "py_image", _DEFAULT_BASE = "DEFAULT_BASE")

def py_image(name, **kwargs):
    if "base" not in kwargs:
        kwargs["base"] = _DEFAULT_BASE
    launcher_image(name, _py_image, **kwargs)


