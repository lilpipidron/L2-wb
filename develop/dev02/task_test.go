package main

import (
	"errors"
	"testing"
)

func TestUnpackString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      error
	}{
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abcd", "abcd", nil},
		{"45", "", errors.New("invalid string")},
		{"", "", nil},
		{`qwe\4\5`, "qwe45", nil},
		{`qwe\45`, "qwe44444", nil},
		{`qwe\\5`, `qwe\\\\\`, nil},
		{`qwe\\`, `qwe\`, nil},
		{`qwe\`, "", errors.New("invalid escape sequence")},
	}

	for _, test := range tests {
		result, err := Unpack(test.input)
		if (err != nil) && err.Error() != test.err.Error() {
			t.Errorf("UnpackString(%s) error = %v, expected error = %v", test.input, err, test.err)
		}
		if result != test.expected {
			t.Errorf("UnpackString(%s) = %s, expected %s", test.input, result, test.expected)
		}
	}
}
