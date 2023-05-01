# Files2Env

## Table of Contents

- [About](#about)
- [Installing](#installing)
- [Usage](#usage)

## About <a name = "about"></a>

Golang module, that imports all files in a folder, and assigns the name and value, as env variables.

Usecase is Kubernetes Service Bindings, where the secrets, are attached to the container, as files.

## Installing

Run the following to get the package

```bash
go get github.com/rhjensen79/files2env@v0.3.0
```

## Usage <a name = "usage"></a>

The following example, shows how to import secrets, from a ./secrets folder

```go
package main

import (
 "github.com/rhjensen79/files2env"
)

func main() {
 files2env.Import("./secrets")
}
```
