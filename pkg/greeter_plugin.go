package greeter_plugin

import "fmt"

type Greeter interface {
    Greet() string
}

func Run(greeter Greeter) {
    fmt.Println(greeter.Greet())
}
