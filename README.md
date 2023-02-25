# NetForemost backend core

API for NetForemost project.

![Go](https://img.shields.io/badge/Golang-1.18-blue.svg?logo=go&longCache=true&style=flat)

## Getting Started

This project uses the **Go** programming language (Golang).

### Prerequisites

[Go](https://golang.org/) at least in version 1.18

### Installing


#### Using GOPATH

```bash
go mod tidy

```

## Running the tests

```bash
go test ./...
```

## Deployment

Clone the repository

```bash
git clone git@github.com:alvaradopcesar/netforemost.git
```

Enter the repository folder

```bash
cd netforemost
```

Build the binary

```bash
go build cmd/main.go
```

Run the program

```bash
# In Unix-like OS
./main

# In Windows
main.exe
```

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://gitlab.com/prettytechnical/oryx-backend-core/-/tags).
