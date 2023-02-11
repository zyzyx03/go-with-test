# Issue

Go 1.18+

From Go 1.18 onwards there is native support for multi-module workspaces. This is done by having a go.work file present in your parent directory.

For a directory structure such as:

```bash
$ tree /my/parent/dir
/my/parent/dir
├── project-one
│   ├── go.mod
│   ├── project-one.go
│   └── project-one_test.go
└── project-two
    ├── go.mod
    ├── project-two.go
    └── project-two_test.go
```

Create and populate the file by executing go work:

```bash
cd /my/parent/dir
go work init
go work use project-one
go work use project-two
```

This will add a go.work file in your parent directory that contains a list of directories you marked for usage:

```text
go 1.18

use (
    ./project-one
    ./project-two
)
```

## reference

[stackOverflow](https://stackoverflow.com/questions/65748509/vscode-shows-an-error-when-having-multiple-go-projects-in-a-directory)
