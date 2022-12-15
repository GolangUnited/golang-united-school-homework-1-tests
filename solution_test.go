package solution

import (
	"strings"
	"testing"
)

func TestGetMessage(t *testing.T) {
	expected := string([]rune{72, 101, 108, 108, 111, 32, 240, 159, 151, 186, 239, 184, 143, 32})
	msg := GetMessage()

	if !strings.EqualFold(msg, expected) {
		t.Errorf("Unexpected result:\n\tExpected: %q\n\tGot: %q", expected, msg)
	}
}
