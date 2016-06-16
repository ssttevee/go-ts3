package ts3util_test

import (
	"github.com/ssttevee/go-ts3/util"
	"testing"
)

func TestEscape(t *testing.T) {
	test := "\\/ |\a\b\f\n\r\t\v"
	expect := `\\\/\s\p\a\b\f\n\r\t\v`
	if s := ts3util.Escape(test); s != expect {
		t.Errorf("got unexpected result: %s\n", s)
	} else {
		t.Logf("%s -> %s\n", test, expect)
	}
}

func TestUnescape(t *testing.T) {
	test := `\\\/\s\p\a\b\f\n\r\t\v`
	expect := "\\/ |\a\b\f\n\r\t\v"
	if s := ts3util.Unescape(test); s != expect {
		t.Errorf("got unexpected result: %s\n", s)
	} else {
		t.Logf("%s -> %s\n", test, expect)
	}
}
