package main

import (
  greeter "github.com/odra/greeter/pkg"
)

type GreetPlugin struct {

}

func (gp *GreetPlugin) Greet() string {
    return "Oi"
}

func Build() greeter.Greeter {
    return &GreetPlugin{}
}
