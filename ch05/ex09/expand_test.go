package main

import "testing"

func TestExpand(t *testing.T) {
	if expand("abcde$fooghi", To123) != "abcde123ghi" {
		t.Error(`error`)
	}
	if expand("$foofoo", To123) != "123foo" {
		t.Error(`error`)
	}
	if expand("$foo$foo", To123) != "123123" {
		t.Error(`error`)
	}
}
