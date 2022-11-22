package main

import (
    "os"
    "fmt"
    greeter "github.com/odra/greeter/pkg"
)

func main() {
    pname := fmt.Sprintf("build/plugin_%s.so", os.Args[1])
    impl, err := greeter.Load(pname)
    if err != nil {
        panic(err)
    }
    
    greeter.Run(impl)
}
