package gospel

import (
	"fmt"
)

var (
	states map[string]any
)

func Test() {
	fmt.Println("Hello World")
}

func SetState(_name string, _value any) {
	states[_name] = _value
}

func GetState(_name string) any {
	return states[_name]
}
