package array

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	param := [][]int{{1, 2, 2, 1}, {2, 2}}
	expected := []int{2, 2}
	actual := intersect(param[0], param[1])
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %#v, Actual: %#v", expected, actual)
	}
}
