package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckBrackets(t *testing.T) {
	tests := map[string]struct {
		value string
		want  bool
	}{
		"empty string": {
			value: "",
			want:  true,
		},
		"without brackets ": {
			value: "Some string",
			want:  true,
		},
		"correct simple brackets": {
			value: "([])",
			want:  true,
		},
		"incorrect simple brackets": {
			value: "{[(]}",
			want:  false,
		},
		"incorrect extra opening brackets": {
			value: "(([])",
			want:  false,
		},
		"incorrect brackets starting with a closing bracket": {
			value: ")(][}{)",
			want:  false,
		},
		"correct nested brackets with symbols": {
			value: "String (14[22{06}])",
			want:  true,
		},
		"incorrect nested brackets with symbols": {
			value: "String (14[22{06]})",
			want:  false,
		},
		"correct brackets with UTF-8 symbols": {
			value: "Строка (14)[22]{06}",
			want:  true,
		},
		"incorrect brackets with UTF-8 symbols": {
			value: "Строка (14}[22]{06)",
			want:  false,
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			req := require.New(t)
			res := checkBrackets(testCase.value)
			req.Equal(testCase.want, res)
		})
	}

}
