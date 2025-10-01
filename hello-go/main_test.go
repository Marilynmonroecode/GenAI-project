package main

import "testing"

func TestSayHello(t *testing.T) {
    got := SayHello()
    want := "Hello, VS Code!"

    if got != want {
        t.Errorf("SayHello() = %q, want %q", got, want)
    }
}
