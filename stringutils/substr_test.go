package stringutils

import "testing"

func TestOverLapStringBorderCondition(t *testing.T) {
	input1 := "helloworld"
	input2 := "hello"
	exp := "hello"
	get1 := OverLapString(input1, input2)
	if get1 != exp {
		t.Error("TestFailed")
	}
}
