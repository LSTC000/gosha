![Go](https://img.shields.io/badge/Go-1.22.0-blue)

# Gosha

A package for working with bash scripts on Go

## Getting started

### Getting Gosha
With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```
import "github.com/LSTC000/gosha"
```

to your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.

Otherwise, run the following Go command to install the package:

```
go get -u github.com/LSTC000/gosha
```

### Creating scripts

Create the following `bash` scripts in the `./scripts` directory:

`hello.sh`

```shell
#!/bin/bash
echo "Hello"
```

`world.sh`

```shell
#!/bin/bash
echo "World"
```

### Running Gosha

```go
package main

import (
	"gosha"
	"log"
	"path"
	"time"
)

const (
	scriptsPath = "./scripts"
	helloScript = "hello.sh"
	worldScript = "world.sh"
)

func main() {
	helloCmd := gosha.Cmd{
		Path:    path.Join(scriptsPath, helloScript),
		Timeout: 0 * time.Second,
	}
	worldCmd := gosha.Cmd{
		Path:    path.Join(scriptsPath, worldScript),
		Timeout: 5 * time.Second,
	}
	
	commands := []gosha.Cmd{helloCmd, worldCmd}
	
	goshaExec := gosha.GetExec(gosha.GetScanner(), commands)

	
	if errs := goshaExec.SyncRun(); errs != nil {
		log.Fatal(errs)
	}
}
```

And use the Go command to run it

```
go run main.go
```