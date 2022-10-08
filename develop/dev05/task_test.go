package main

import (
	"testing"
)

func TestGrep(t *testing.T) {
	flgs := customFlags{
		after:      0,
		before:     0,
		context:    0,
		count:      false,
		ignoreCase: false,
		invert:     false,
		fixed:      false,
		index:      false,
	}

	grep := &customGrep{
		flags:  flgs,
		regExp: "aba",
		content: []fileContent{
			{"ggff", 1},
			{"gabaf", 2},
			{"yhfrg", 3},
			{"vfdabar", 4},
			{"ffff", 5},
			{"sssss", 6},
		},
	}

	expect := "gabaf\nvfdabar"
	result, err := grep.Grep()

	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	if result != expect {
		t.Errorf("got %q, wanted %q", result, expect)
	}
}

func TestJoin(t *testing.T) {
	content := []fileContent{
		{"yhfrg", 3},
		{"vfdabar", 4},
		{"sssss", 6},
	}

	expect := "3: yhfrg\n4: vfdabar\n6: sssss"
	result := Join(content, true)

	if expect != result {
		t.Errorf("wanted %q, got %q", expect, result)
	}
}
