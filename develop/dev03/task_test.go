package main

import (
	"reflect"
	"testing"
)

func TestStrings(t *testing.T) {
	strct := []struct {
		input  []string
		expect []string
	}{
		{[]string{"cde", "bca", "abc", "zxc"}, []string{"abc", "bca", "cde", "zxc"}},
		{[]string{"2", "1", "4", "4", "5"}, []string{"1", "2", "4", "4", "5"}},
	}

	for _, element := range strct {
		result := SortUtil(element.input, 0, false, false, false)
		if eq := reflect.DeepEqual(result, element.expect); !eq {
			t.Errorf("got %strct, wanted %strct", element.input, element.expect)
		}
	}
}

func TestUniq(t *testing.T) {
	strct := []struct {
		input  []string
		expect []string
	}{
		{[]string{"cde", "bca", "abc", "zxc", "abc"}, []string{"abc", "bca", "cde", "zxc"}},
		{[]string{"2", "1", "4", "4", "5"}, []string{"1", "2", "4", "5"}},
	}

	for _, element := range strct {
		result := SortUtil(element.input, 0, false, false, true)
		if eq := reflect.DeepEqual(result, element.expect); !eq {
			t.Errorf("got %strct, wanted %strct", element.input, element.expect)
		}
	}
}

func TestReverse(t *testing.T) {
	strct := []struct {
		input  []string
		expect []string
	}{
		{[]string{"cde", "bca", "abc", "zxc"}, []string{"zxc", "cde", "bca", "abc"}},
		{[]string{"2", "1", "4", "4", "5"}, []string{"5", "4", "4", "2", "1"}},
	}

	for _, element := range strct {
		result := SortUtil(element.input, 0, false, true, false)
		if eq := reflect.DeepEqual(result, element.expect); !eq {
			t.Errorf("got %strct, wanted %strct", element.input, element.expect)
		}
	}
}
