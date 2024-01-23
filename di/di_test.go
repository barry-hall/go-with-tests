package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	Greet(&buffer, "River")

	got := buffer.String()
	want := "Hello, River"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
