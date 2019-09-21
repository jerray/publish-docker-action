package main

import (
	"reflect"
	"testing"
)

func Test_resolveSemanticVersionTag(t *testing.T) {
	cases := []struct {
		input  string
		expect []string
	}{
		// invalid semantic version
		{"1", []string{"1"}},
		{"master", []string{"master"}},
		{"20190921-actions", []string{"20190921-actions"}},

		// valid semantic version
		{"0.0.0", []string{"0", "0.0", "0.0.0"}},
		{"1.0.0", []string{"1", "1.0", "1.0.0"}},
		{"2.7.10", []string{"2", "2.7", "2.7.10"}},
		{"v1.0.0", []string{"1", "1.0", "1.0.0"}},
		{"v2.3.6", []string{"2", "2.3", "2.3.6"}},
		{"3.0.0-rc1", []string{"3-rc1", "3.0-rc1", "3.0.0-rc1"}},
		{"v3.0.0-alpha.1", []string{"3-alpha.1", "3.0-alpha.1", "3.0.0-alpha.1"}},
		{"4.0.0-alpha.1.2", []string{"4-alpha.1.2", "4.0-alpha.1.2", "4.0.0-alpha.1.2"}},
	}
	for _, c := range cases {
		actual := resolveSemanticVersionTag(c.input)
		if !reflect.DeepEqual(c.expect, actual) {
			t.Errorf("expect %v, actual %s", c.expect, actual)
		}
	}
}
