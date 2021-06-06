# Baisc Go

### :peach: Go CLI
- `go build` = compile go code file eg: `go build main.go` will produce main.exe, to execute Mac: ./main Win: click on main.exe
- `go run` = compile andexecute 1-2 files eg: go run main.go
- `go fmt` = format all the code in each file in the current dir
- `go install` = compile and install a package
- `go get` = download the raw source code of package
- `go test` = run any tests associated with the current project

### :peach: Note

####  Package statement
1. Executable (package main) = generate a file that we can run, contains '*.go' files. (another name than main, when run go build it will not generate executable file. It requires func main() in the package, which is run auto when exec.
2. Reusable (package xxx) = code used as 'helpers where we put reusable logic

####  Import statement
- to import other package to be used in this package. ['fmt'](https://golang.org/pkg/fmt/) is a standard package of go.