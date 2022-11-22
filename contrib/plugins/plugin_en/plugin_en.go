package main

import (
  greeter "github.com/odra/greeter/pkg"
)

type GreetPlugin struct {

}

func (gp *GreetPlugin) Greet() string {
    return "Hello"
}

func Build() greeter.Greeter {
    return &GreetPlugin{}
}
