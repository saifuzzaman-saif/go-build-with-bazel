load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "greetings",
    srcs = ["greetings.go"],
    importpath = "GO-BUILD-WITH-BAZEL/com/saif/hello-world/greetings",
    visibility = ["//visibility:public"],
)

go_test(
    name = "greetings_test",
    srcs = ["greetings_test.go"],
    embed = [":greetings"],
)
