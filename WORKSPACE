workspace(name = "riegeli")

load(
    "@bazel_tools//tools/build_defs/repo:http.bzl",
    "http_archive",
)

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "a82a352bffae6bee4e95f68a8d80a70e87f42c4741e6a448bec11998fcc82329",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.18.5/rules_go-0.18.5.tar.gz"],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
)

load(
    "@io_bazel_rules_go//go:deps.bzl",
    "go_register_toolchains",
    "go_rules_dependencies",
)

go_register_toolchains()

go_rules_dependencies()

load(
    "@bazel_gazelle//:deps.bzl",
    "gazelle_dependencies",
    "go_repository",
)

gazelle_dependencies()

http_archive(
    name = "build_stack_rules_proto",
    sha256 = "5474d1b83e24ec1a6db371033a27ff7aff412f2b23abba86fedd902330b61ee6",
    strip_prefix = "rules_proto-91cbae9bd71a9c51406014b8b3c931652fb6e660",
    urls = ["https://github.com/stackb/rules_proto/archive/91cbae9bd71a9c51406014b8b3c931652fb6e660.tar.gz"],
)

http_archive(
    name = "org_brotli",
    sha256 = "fb511e09ea284fcd18fe2a2632744609a77f69c345428b9f0d2cc15171215f06",
    strip_prefix = "brotli-ee2a5e1540cbd6ef883a897499d9596307f7f7f9",
    url = "https://github.com/google/brotli/archive/ee2a5e1540cbd6ef883a897499d9596307f7f7f9.zip",
)

http_archive(
    name = "org_brotli_go",
    sha256 = "fb511e09ea284fcd18fe2a2632744609a77f69c345428b9f0d2cc15171215f06",
    strip_prefix = "brotli-ee2a5e1540cbd6ef883a897499d9596307f7f7f9/go",
    url = "https://github.com/google/brotli/archive/ee2a5e1540cbd6ef883a897499d9596307f7f7f9.zip",
)

go_repository(
    name = "com_github_datadog_zstd",
    importpath = "github.com/DataDog/zstd",
    tag = "v1.3.5",
)

go_repository(
    name = "com_github_minio_highwayhash",
    importpath = "github.com/minio/highwayhash",
    tag = "v1.0.0",
)

go_repository(
    name = "com_github_google_go_cmp",
    importpath = "github.com/google/go-cmp",
    tag = "v0.2.0",
)
