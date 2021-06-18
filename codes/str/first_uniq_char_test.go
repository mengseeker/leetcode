package str

import "testing"

func TestFirstUniqChar(t *testing.T) {
	param := "loveleetcode"
	expected := 2
	acl := firstUniqChar(param)
	if acl != expected {
		t.Errorf("Expected \n%#v, Actual: \n%#v", expected, acl)
	}
}
