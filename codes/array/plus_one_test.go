package array

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	param := [][]int{{1, 2, 3}}
	expected := []int{1, 2, 4}
	actual := plusOne(param[0])
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %#v, Actual: %#v", expected, actual)
	}
}
