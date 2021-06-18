package str

import "testing"

func TestIsAnagram(t *testing.T) {
	param1, param2 := "anagram", "nagaram"
	expected := true
	acl := isAnagram(param1, param2)
	if acl != expected {
		t.Errorf("Expected \n%#v, Actual: \n%#v", expected, acl)
	}
}
