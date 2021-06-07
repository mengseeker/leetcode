package array

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	param := [][]int{{0, 1, 0, 3, 12}}
	expected := []int{1, 3, 12, 0, 0}
	moveZeroes(param[0])
	actual := param[0]
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %#v, Actual: %#v", expected, actual)
	}
}
