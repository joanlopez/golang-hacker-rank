package implementation

import "testing"

type viralAdvertisingTest struct {
	n        int32
	expected int32
}

var viralAdvertisingTests = []viralAdvertisingTest {
	{1, 2},
	{2, 5},
	{3, 9},
	{4, 15},
	{5, 24},
}

func TestViralAdvertising(t *testing.T) {
	for _, tt := range viralAdvertisingTests {
		actual := ViralAdvertising(tt.n)
		if actual != tt.expected {
			t.Errorf("ViralAdvertising(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}