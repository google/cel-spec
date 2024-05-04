load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def _non_module_dependencies_impl(_ctx):
    http_archive(
        name = "com_google_googleapis",
        sha256 = "bd8e735d881fb829751ecb1a77038dda4a8d274c45490cb9fcf004583ee10571",
        strip_prefix = "googleapis-07c27163ac591955d736f3057b1619ece66f5b99",
        urls = [
            "https://github.com/googleapis/googleapis/archive/07c27163ac591955d736f3057b1619ece66f5b99.tar.gz",
        ],
    )

non_module_dependencies = module_extension(
    implementation = _non_module_dependencies_impl,
)
