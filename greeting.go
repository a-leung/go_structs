package main

type Greeter struct {
	GreetingMessage string
}

func (g *Greeter) SayHello() string {
	return g.GreetingMessage
}
