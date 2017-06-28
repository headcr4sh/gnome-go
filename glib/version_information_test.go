package glib

import (
	"testing"
)

func TestCheckVersion(t *testing.T) {
	var tests = []struct {
		major uint
		minor uint
		micro uint
		ok    bool
		msg   string
	}{
		{MajorVersion, MinorVersion, MicroVersion, true, ""},
		{0, 0, 0, false, "GLib version too new (major mismatch)"},
	}
	for _, test := range tests {
		ok, msg := CheckVersion(test.major, test.minor, test.micro)
		if ok != test.ok || msg != test.msg {
			t.Errorf(`CheckVersion(%d, %d, %d) => %v "%s", expected: %v "%s"`, test.major, test.minor, test.micro, ok, msg, test.ok, test.msg)
		}
	}
}
