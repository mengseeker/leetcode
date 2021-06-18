package str

import "testing"

func TestIsPalindrome(t *testing.T) {
	param1 := "A man, a plan, a canal: Panama"
	expected := true
	acl := isPalindrome(param1)
	if acl != expected {
		t.Errorf("Expected \n%#v, Actual: \n%#v", expected, acl)
	}
}
