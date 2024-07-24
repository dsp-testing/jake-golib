package jake_golib_test

import (
	jake_golib "github.com/dsp-testing/jake-golib"
	"testing"
)

func TestAdd(t *testing.T) {
	if jake_golib.Add(1, 2) != 3 {
		t.Fail()
	}
}
