package pkgtest

import "fmt"
import _ "github.com/chuckha/pkgtest/somedir"

func init() {
	fmt.Println("welcome to package pkgtest from file two")
}
