# Golang Gowther
lorem ipsum dolor amet

## How to install
```sh
go get -u github.com/shandysiswandi/gowther
```

## Usage
```go
package main

import (
	"fmt"

	"github.com/shandysiswandi/gowther"
)

func main() {
	v := gowther.Version(1)

	fmt.Println("Version of Gowther is", v)
}
```