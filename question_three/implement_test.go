package question_three

import "testing"

func TestRand13ToRand5(t *testing.T) {
	randVal := Rand13ToRand5()
	if randVal > 5 || randVal < 0 {
		t.Errorf("Rand13ToRand5 receive %d, out of rang [0,5]", randVal)
	}
}

func TestRand5ToRand13(t *testing.T) {
	randVal := Rand13ToRand5()
	if randVal > 13 || randVal < 0 {
		t.Errorf("TestRand5ToRand13 receive %d, out of rang [0,13]", randVal)
	}
}