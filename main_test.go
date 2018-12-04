package main

import "testing"

func TestHashURL(t *testing.T) {
	h := hashURL("abcd")
	t.Log(h)
}
