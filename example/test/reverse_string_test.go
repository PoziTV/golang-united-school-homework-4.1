package test_reverse_string

import (
	"reverse_string"
	"testing"
)

func Test_reverse_string(b *testing.T) {

	inputStr := "Gladys"
	want := "sydalG"
	// msg, err := Hello("Gladys")
	if want != reverse_string.ReverseString(inputStr) {
		b.Fatalf(inputStr, want)
	}
}
