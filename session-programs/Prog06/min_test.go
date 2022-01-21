package utility

import "testing"

func TestGetMin(t *testing.T) {
	ans := GetMin(10, 3)
	if ans != 3 {
		t.Error("GetMin(10,3) = %v, want 3", ans)
	}
}
