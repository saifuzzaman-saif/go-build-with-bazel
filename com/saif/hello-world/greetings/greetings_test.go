package greetings

import (
	"testing"
)

func Test_Greet(t *testing.T) {
	greetMsg := Greet("M. Geremy")
	expectedMsg := "Hello M. Geremy!"

	if greetMsg != expectedMsg {
		t.Errorf("Greet() returned %q, but expected %q", greetMsg, expectedMsg)
	}
}
