package win

import "testing"

func TestControlTest(t *testing.T) {
	_, err := controlText(0x3515c2)
	if err != nil {
		t.Fatalf("controlText: %s", err)
	}
}
