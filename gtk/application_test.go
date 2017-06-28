package gtk

import "testing"

func TestIsValidApplicationID(t *testing.T) {
	var tests = []struct {
		id    string
		valid bool
	}{
		{"ExampleApp", false},
	}
	for _, test := range tests {
		if valid := IsValidApplicationID(test.id); !valid {
			t.Errorf("IsValidApplicationID(%s) => %v, expected: %v", test.id, valid, test.valid)
		}
	}
}
