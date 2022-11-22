package main

import (
    "os"
    "fmt"
    "plugin"
    greeter "github.com/odra/greeter/pkg"
)

func main() {
    pname := fmt.Sprintf("build/plugin_%s.so", os.Args[1])

    p, err := plugin.Open(pname)
    if err != nil {
        panic(err)
    }

    fn, err := p.Lookup("Build")
    if err != nil {
        panic(err)
    }

    impl := fn.(func() greeter.Greeter)()

    greeter.Run(impl)

    //fmt.Printf("%v", fn.(func() PluginRunner))

    //gp := fn.(func())()

    //run(gp.(PluginRunner))
}
