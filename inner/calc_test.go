package inner

import "testing"

func TestAdd(t *testing.T) {
	rsl := Add(1, 1)
	if rsl != 2 {
		t.Errorf("Fail when add 1 + 1, expect: %d, actual %d", 2, rsl)
	}
}
