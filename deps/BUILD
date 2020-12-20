package(default_visibility = ["PUBLIC"])

go_get(
    name = "ansi",
    get = "github.com/mgutz/ansi",
    revision = "d51e80ef957dba7f19388ca64afefbd5a096af30",
    deps = [
        ":colorable",
    ],
)

go_get(
    name = "colorable",
    get = "github.com/mattn/go-colorable",
    revision = "v0.1.8",
    deps = [
        ":isatty",
    ],
)

go_get(
    name = "isatty",
    get = "github.com/mattn/go-isatty",
    revision = "v0.0.12",
    deps = [
        "//deps/x:sys",
    ],
)
