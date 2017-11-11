package main

import "testing"

func TestGreeting(t *testing.T) {
	english := &Greeter {
		GreetingMessage: "Good day",
	}

	if (english.SayHello() != "Good day") {
		t.Error("not getting value from struct")
	}

	french := &Greeter {
		GreetingMessage: "Bon Jour",
	}

	if (french.SayHello() == "Good day") {
		t.Error("not getting value from struct")
	}
}
