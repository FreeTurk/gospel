package gospel

import (
	"fmt"
	"os"
)

var (
	file *os.File
)

func Test() {
	fmt.Println(file.Name())
}

func Initial() {
	f, err := os.CreateTemp("", "sample.*.yaml")
	if err != nil {
		panic(err)
	}
	file = f
}
