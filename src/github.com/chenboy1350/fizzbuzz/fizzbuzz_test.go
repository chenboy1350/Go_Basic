package fizzbuzz

import (
	"testing"
)

func TestInput1ShouldBeDisplay1(t *testing.T) {
	v := FizzBuzz(1)
	if v != "1" {
		t.Error("fizzbuzz of 1 should be '1' but have", v)
	}
}

func TestInput2ShouldBeDisplay2(t *testing.T) {
	v := FizzBuzz(2)
	if v != "2" {
		t.Error("fizzbuzz of 2 should be '2' but have", v)
	}
}
