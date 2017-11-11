package main

import "testing"

func TestHello(t *testing.T) {
	if true == false {
		t.Error("true is not false?!")
	}
}
