package implementation

import "testing"

func TestKangaroo(t *testing.T) {
	var x1 int32 = 0; var x2 int32 = 5
	var v1 int32 = 2; var v2 int32 = 3
	if match := Kangaroo(x1, v1, x2, v2); match != "NO" {
		t.Errorf("Incorrect kangaroo match, got: %v, want: NO.", 1)
	}
}
