load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

go_library(
    name = "riegeli",
    srcs = [
        "compression.go",
        "reader.go",
        "riegeli.go",
        "transpose_decoder.go",
        "transpose_encoder.go",
        "transpose_util.go",
        "util.go",
        "writer.go",
    ],
    importpath = "riegeli",
    visibility = ["//visibility:public"],
    deps = [
        ":records_metadata_go_proto",
        "@com_github_datadog_zstd//:go_default_library",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@com_github_minio_highwayhash//:go_default_library",
        "@org_brotli_go//cbrotli",
    ],
)

go_library(
    name = "compare",
    srcs = [
        "compare.go",
    ],
    importpath = "riegeli/compare",
    deps = [
        "@com_github_golang_protobuf//proto:go_default_library",
        "@com_github_google_go_cmp//cmp:go_default_library",
    ],
)

go_test(
    name = "riegeli_test",
    srcs = ["riegeli_test.go"],
    embed = [
        ":riegeli",
    ],
    shard_count = 10,
    deps = [
        ":compare",
        ":riegeli_test_go_proto",
    ],
)

go_test(
    name = "block_test",
    srcs = ["block_test.go"],
    embed = [":riegeli"],
)

go_test(
    name = "riegeli_bench_test",
    srcs = ["riegeli_bench_test.go"],
    embed = [":riegeli"],
)

go_test(
    name = "transpose_test",
    srcs = ["transpose_test.go"],
    embed = [":riegeli"],
)

proto_library(
    name = "riegeli_test_proto",
    srcs = ["riegeli_test.proto"],
)

go_proto_library(
    name = "riegeli_test_go_proto",
    importpath = "riegeli/riegeli_test_go_proto",
    proto = ":riegeli_test_proto",
)

proto_library(
    name = "records_metadata_proto",
    srcs = ["records_metadata.proto"],
    deps = ["@com_google_protobuf//:descriptor_proto"],
)

go_proto_library(
    name = "records_metadata_go_proto",
    importpath = "riegeli/records_metadata_go_proto",
    proto = ":records_metadata_proto",
)
