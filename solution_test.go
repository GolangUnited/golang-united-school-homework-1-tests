package solution

import (
	"reflect"
	"testing"
)

func Test_GetMessage(t *testing.T) {
	expected := []byte{72, 101, 108, 108, 111, 32, 240, 159, 151, 186, 239, 184, 143, 32}
	msg := []byte(GetMessage())

	if !reflect.DeepEqual(expected, msg) {
		t.Errorf("Unexpected result:\n\tExpected: %q\n\tGot: %q", expected, msg)
	}
}
