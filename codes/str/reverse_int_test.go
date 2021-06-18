package str

import (
	"testing"
)

func TestReverseInt(t *testing.T) {
	param := 120
	expected := 21
	acl := reverseInt(param)
	if acl != expected {
		t.Errorf("Expected \n%#v, Actual: \n%#v", expected, acl)
	}
}
