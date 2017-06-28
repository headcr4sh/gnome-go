package glib

import (
	"testing"
)

func TestIsDirDirSeparator(t *testing.T) {
	var tests = []struct {
		sep rune
		res bool
	}{
		{'/', true},
		{'a', false},
		{'b', false},
		{'C', false},
	}
	for _, test := range tests {
		if res := IsDirSeparator(test.sep); res != test.res {
			t.Errorf(`IsDirSeparator(%v) => %v, expected: %v`, test.sep, res, test.res)
		}
	}
}
