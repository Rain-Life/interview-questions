package question_one

import (
	"testing"
)

func TestPerm(t *testing.T) {
	input := []rune("ABC")
	result := map[string]bool {
		"ABC": true,
		"ACB": true,
		"BAC": true,
		"BCA": true,
		"CBA": true,
		"CAB": true,
	}
	Perm(input, func(a []rune) {
		_, ok := result[string(a)]
		if !ok {
			t.Errorf("input is %s, expected is %s", string(input), string(a))
		}
		delete(result, string(a))
	})
	if len(result) > 0 {
		t.Errorf("input is %s, expected is %v", string(input), result)
	}
}
