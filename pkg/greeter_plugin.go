package greeter_plugin

import (
    "fmt"
    "plugin"
)

type Greeter interface {
    Greet() string
}

func Load(path string) (Greeter, error) {
    p, err := plugin.Open(path)
    if err != nil {
        return nil, err
    }

    fn, err := p.Lookup("Build")
    if err != nil {
        return nil, err
    }

    return fn.(func() Greeter)(), nil
}

func Run(greeter Greeter) {
    fmt.Println(greeter.Greet())
}
