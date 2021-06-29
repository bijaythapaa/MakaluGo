package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Vijay")

	got := buffer.String()
	want := "Hello, Vijay"

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}
