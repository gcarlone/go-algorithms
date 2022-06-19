package unique

import (
	"testing"
)

func TestUniqueInts(t *testing.T) {
	values := []int{1, 2, 3, 3, 4}
	uniqueInts := Ints(values)
	if len(uniqueInts) != 4 {
		t.Errorf("Expected %v unique ints, but founded %v", 4, len(uniqueInts))
	}
	if uniqueInts[0] != 1 {
		t.Errorf("Expected %v, but founded %v", 1, uniqueInts[0])
	}
	if uniqueInts[1] != 2 {
		t.Errorf("Expected %v, but founded %v", 2, uniqueInts[1])
	}
	if uniqueInts[2] != 3 {
		t.Errorf("Expected %v, but founded %v", 4, uniqueInts[3])
	}
}
