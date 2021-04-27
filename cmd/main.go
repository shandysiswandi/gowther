package main

import (
	"fmt"

	"github.com/shandysiswandi/gowther"
)

func main() {
	v := gowther.Version()

	fmt.Println("Version of Gowther is", v)
}
