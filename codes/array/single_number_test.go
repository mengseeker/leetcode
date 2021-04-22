package array

import "testing"

func TestSingleNumber(t *testing.T) {
	nums := []int{1, 2, 3, 1, 3}
	expected := 2
	actual := singleNumber(nums)
	if actual != expected {
		t.Errorf("Expected %#v, Actual: %#v", expected, actual)
	}
}
