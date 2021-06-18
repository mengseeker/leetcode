package str

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	str := []byte{'h', 'e', 'l', 'l', 'o'}
	expected := []byte{'o', 'l', 'l', 'e', 'h'}
	reverseString(str)
	if !reflect.DeepEqual(str, expected) {
		t.Errorf("Expected \n%#v, Actual: \n%#v", expected, str)
	}
}
