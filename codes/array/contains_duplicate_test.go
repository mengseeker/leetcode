package array

import "testing"

func TestContainsDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	expected := false
	actual := containsDuplicate(nums)
	if actual != expected {
		t.Errorf("Expected %#v, Actual: %#v", expected, actual)
	}
}
