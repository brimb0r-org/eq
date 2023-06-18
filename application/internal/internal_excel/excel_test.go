package internal_excel

import "testing"

func TestGetStatus(t *testing.T) {
	tests := map[string]struct {
		a        float64
		expected string
	}{
		"safe": {
			expected: "safe",
			a:        11000,
		},
		"managed": {
			expected: "managed",
			a:        6500,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := getStatus(test.a)
			if actual != test.expected {
				t.Errorf("doesnt match: expected [%s] got [%s]", test.expected, actual)
			}
		})
	}
}
