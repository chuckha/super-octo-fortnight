package main

import (
	"fmt"

	_ "github.com/chuckha/pkgtest/somedir"
	_ "github.com/chuckha/pkgtest/somedir/someotherdir"
)

func main() {
	fmt.Println("hello world and packages")
}
