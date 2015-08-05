package fibonacci

import "testing"

func expect(t *testing.T, cond string, expectation, result bool) {
	if result != expectation {
		t.Errorf("Expected %s to be %t, got %t", cond, expectation, result)
	}
}

func TestIsFib(t *testing.T) {
	expect(t, "IsFib(0)", true, IsFib(0))
	expect(t, "IsFib(1)", true, IsFib(1))
	expect(t, "IsFib(55)", true, IsFib(55))
	expect(t, "IsFib(100)", false, IsFib(100))
}

/*func TestIsFib(t *testing.T) {
	if !IsFib(0) {
		t.Errorf("Testing IsFib failed for 0, expected true got false")
	}
	if !IsFib(1) {
		t.Errorf("Testing IsFib failed for 1, expected true got false")
	}
	if !IsFib(55) {
		t.Errorf("Testing IsFib failed for 55, expected true got false")
	}
	if IsFib(100) {
		t.Errorf("Testing IsFib failed for 100, expected false got true")
	}
}*/
