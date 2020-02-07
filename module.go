package module2

import (
	"fmt"

	module1 "github.com/cocorambo/go-module1"
)

const ModuleName = "Go module 2"

func Print() {
	fmt.Println(module1.ModuleName)
}
