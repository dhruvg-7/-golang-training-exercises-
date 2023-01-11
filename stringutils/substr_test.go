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
func TestOverLapStringNullCondition(t *testing.T) {
	input1 := ""
	input2 := ""
	exp := ""
	get1 := OverLapString(input1, input2)
	if get1 != exp {
		t.Error("TestFailed empty case")
	}
}
func TestOverLapStringSecondCondition(t *testing.T) {
	input1 := "hello"
	input2 := ""
	exp := ""
	get1 := OverLapString(input1, input2)
	if get1 != exp {
		t.Error("TestFailed second case")
	}
}
func TestOverLapStringThirdCondition(t *testing.T) {
	input1 := "hello"
	input2 := "walp"
	exp := ""
	get1 := OverLapString(input1, input2)
	if get1 != exp {
		t.Error("TestFailed third case")
	}
}
func TestOverLapStringfourthCondition(t *testing.T) {
	input1 := "h"
	input2 := "walp"
	exp := ""
	get1 := OverLapString(input1, input2)
	if get1 != exp {
		t.Error("TestFailed fourth case")
	}
}
