package fibbonaci

import "testing"

func compareResult(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestNthFiboZeroCase(t *testing.T) {
	input1 := 0
	exp := []int{}
	f := NthFibo()
	get1 := f(input1)
	if !compareResult(exp, get1) {
		t.Error("TestFailed")
		t.Errorf("Want: %v\n Get: %v\n", exp, get1)
	}
}
func TestNthFiboBaseCase(t *testing.T) {
	input1 := 3
	exp := []int{0, 1, 1}
	f := NthFibo()
	get1 := f(input1)
	if !compareResult(exp, get1) {
		t.Error("TestFailed")
		t.Errorf("Want: %v\n Get: %v\n", exp, get1)
	}
}
func TestNthFiboFirstTen(t *testing.T) {
	input1 := 10
	exp := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	f := NthFibo()
	get1 := f(input1)
	if !compareResult(exp, get1) {
		t.Errorf("Want: %v\n Get: %v\n", exp, get1)
		t.Error("TestFailed")
	}
}

func TestNthFiboMultiCall(t *testing.T) {
	input1 := 10
	input2 := 8
	input3 := 15
	exp1 := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	exp2 := []int{0, 1, 1, 2, 3, 5, 8, 13}
	exp3 := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377}

	f := NthFibo()
	get1 := f(input1)
	get2 := f(input2)
	get3 := f(input3)
	if !compareResult(exp1, get1) {
		t.Error("TestFailed")
		t.Errorf("want: %v\n Get: %v", exp1, get1)
	}
	if !compareResult(exp2, get2) {
		t.Error("TestFailed")
		t.Errorf("want: %v\n Get: %v", exp2, get2)
	}
	if !compareResult(exp3, get3) {
		t.Error("TestFailed")
		t.Errorf("want: %v\n Get: %v", exp3, get3)
	}
}
