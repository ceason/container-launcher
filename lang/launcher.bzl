load(
    "@io_bazel_rules_docker//container:image.bzl",
    _container_image = "container_image",
    _image = "image",
)
load(
    "@io_bazel_rules_docker//container:providers.bzl",
    _ImageInfo = "ImageInfo",
)
load(
    "@bazel_tools//tools/build_defs/hash:hash.bzl",
    _hash_tools = "tools",
    _sha256 = "sha256",
)
load(
    "@io_bazel_rules_docker//container:layer_tools.bzl",
    _assemble_image = "assemble",
    _get_layers = "get_from_target",
    _incr_load = "incremental_load",
    _layer_tools = "tools",
)
load(
    "@io_bazel_rules_docker//skylib:path.bzl",
    "dirname",
    "strip_prefix",
    _canonicalize_path = "canonicalize",
    _join_path = "join",
)

# - add launcher to base layer (defaults specified in respective <lang>.bzl wrappers)
# - create lang image using new base
# - prefix entrypoint with launcher
#

def launcher_image(name, image_rule, base = None, **kwargs):
    base_image = name + ".launcher-layer"
    lang_image = name + ".lang-image"
    _container_image(
        name = base_image,
        base = base,
        layers = ["//cmd/container-launcher:layer"],
    )
    image_rule(
        name = lang_image,
        base = ":" + base_image,
        **kwargs
    )
    _launcher_image(
        name = name,
        lang_image = ":" + lang_image,
    )

def _repository_name(ctx):
    """Compute the repository name for the current rule."""
    if ctx.attr.legacy_repository_naming:
        # Legacy behavior, off by default.
        return _join_path(ctx.attr.repository, ctx.label.package.replace("/", "_"))

    # Newer Docker clients support multi-level names, which are a part of
    # the v2 registry specification.

    return _join_path(ctx.attr.repository, ctx.label.package)

def _impl(ctx):
    runfiles = []
    transitive_runfiles = []

    image = ctx.attr.lang_image[_ImageInfo]
    container_parts = {}
    container_parts.update(image.container_parts)

    # prefix the image config's entrypoint array with our launcher
    config = ctx.actions.declare_file(ctx.attr.name + ".config")
    container_parts["config"] = config
    container_parts["config_digest"] = _sha256(ctx, config)
    runfiles.append(config)
    runfiles.append(container_parts["config_digest"])
    ctx.actions.run(
        inputs = [image.container_parts["config"]],
        outputs = [config],
        executable = ctx.executable._launcher_injector,
        tools = ctx.attr._launcher_injector.default_runfiles.files,
        arguments = [
            "--input=%s" % image.container_parts["config"].path,
            "--output=%s" % config.path,
            # TODO(ceason): parameterize the launcher
            "--entrypoint_prefix=/container-launcher",
            #"--entrypoint_prefix=--",
        ],
    )

    tag_name = _repository_name(ctx) + ":" + ctx.attr.name
    images = {
        tag_name: container_parts,
    }
    transitive_runfiles.append(ctx.attr.lang_image.default_runfiles.files)
    transitive_runfiles.append(ctx.attr.lang_image.data_runfiles.files)

    _incr_load(
        ctx,
        images,
        ctx.outputs.executable,
        run = True,
        run_flags = image.docker_run_flags,
    )

    return [
        DefaultInfo(
            files = depset(direct = runfiles),
            runfiles = ctx.runfiles(
                files = runfiles,
                transitive_files = depset(transitive = transitive_runfiles),
            ),
        ),
        _ImageInfo(
            container_parts = container_parts,
            docker_run_flags = image.docker_run_flags,
            legacy_run_behavior = image.legacy_run_behavior,
        ),
    ]

_launcher_image = rule(
    implementation = _impl,
    executable = True,
    toolchains = ["@io_bazel_rules_docker//toolchains/docker:toolchain_type"],
    attrs = _hash_tools + _image.attrs + {
        "lang_image": attr.label(
            providers = [_ImageInfo],
        ),
        "_launcher_injector": attr.label(
            executable = True,
            cfg = "host",
            default = Label("//lang:launcher_injector"),
        ),
    },
)
