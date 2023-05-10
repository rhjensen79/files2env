# Files2Env

## Table of Contents

- [About](#about)
- [Installing](#installing)
- [Usage](#usage)
- [Collaboration](#collaboration)

## About <a name = "about"></a>

Golang module, that imports all files in a folder, and assigns the names and value, as env variables.

Usecase is Kubernetes Service Bindings, where the secrets, are attached to the container, as files.

This allows you to easy extract the secrets, and map them as env variables, to be used in your code.

## Installing

Run the following to install the package

```bash
go get github.com/rhjensen79/files2env@v0.7.5
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

## Collaboration

If you find issues, then feel free to create and issue, or help with making the code better.
