package codegenerator

import (
	"testing"
)

func TestCodeGenerator(t *testing.T) {
	testCase := []struct {
		length int
	}{
		{4},
		{6},
		{10},
	}

	for _, tc := range testCase {
		t.Run("", func(t *testing.T) {
			randomString := CodeGenerator(tc.length)

			if len(randomString) != tc.length {
				t.Errorf("Expected random string of length %d, but got length %d", tc.length, len(randomString))
			}

			for _, char := range randomString {
				if !isValidCharacter(char) {
					t.Errorf("Invalid character in random string %c", char)
				}
			}
		})
	}
}

func isValidCharacter(c rune) bool {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for _, validChar := range charset {
		if c == validChar {
			return true
		}
	}
	return false
}
