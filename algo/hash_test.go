package algo

import "testing"

func TestGetMD5Hash(t *testing.T) {
	expect := "900150983cd24fb0d6963f7d28e17f72"
	if got := GetMD5Hash("abc"); got != expect {
		t.Errorf("Fail - expected %s but got %s", got, expect)
	}
}
