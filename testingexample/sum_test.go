package testingexample

import (
	"testing"
)

func TestSum(t *testing.T) {
	actual := Sum(10, 20)
	expected := 30
	if actual != expected {
		t.Errorf("got  %v\nwant %v", actual, expected)
	}

	actual = Sum(0, 0)
	expected = 0
	if actual != expected {
		t.Errorf("got  %v\nwant %v", actual, expected)
	}
	actual = Sum(0, -1)
	expected = -1
	if actual != expected {
		t.Errorf("got  %v\nwant %v", actual, expected)
	}
}
