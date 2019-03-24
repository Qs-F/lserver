package lserver

import (
	"bufio"
	"io"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	o := &Override{}
	tc := []string{"Hello World"}
	for _, tstr := range tc {
		t.Log("test case: ", tstr)
		r, err := o.Map(strings.NewReader(tstr))
		if err != nil {
			t.Fatal(err)
			return
		}
		buf := bufio.NewReader(r)
		b, err := buf.ReadBytes(byte(0x00))
		if err != nil && err != io.EOF {
			t.Fatal(err)
			return
		}
		if string(b) == tstr {
			t.Log(string(b))
		} else {
			t.Fatal("not the same test string: ", b)
		}
	}
}
