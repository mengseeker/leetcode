package codes

import (
	"reflect"
	"testing"
)

func TestRotateArray(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	expected := []int{5, 6, 7, 1, 2, 3, 4}
	rotate(nums, 3)
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("Expected %#v, Actual: %#v", expected, nums)
	}
}
