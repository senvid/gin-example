package util

import (
	"testing"
)

func TestCreatesecureCookie(t *testing.T)  {

	v:="test"
	n:="user"

	s := createSignedValue(v,n)
	if s!="" {
		t.Log(s)
	}
}