load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "runfiles",
    srcs = ["main.go"],
    data = ["@go_sdk//:tools"],
    deps = ["@io_bazel_rules_go//go/runfiles"],
)