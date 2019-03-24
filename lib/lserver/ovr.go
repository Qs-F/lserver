package lserver

import (
	"bytes"
	"io"
	"io/ioutil"
)

type Overrider interface {
	Map(io.Reader) (io.Reader, error)
}

type Override struct {
	Target string
	With   string
}

func New(target, with string) *Override {
	return &Override{
		Target: target,
		With:   with,
	}
}

func (o *Override) Map(in io.Reader) (io.Reader, error) {
	origin, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(bytes.Replace(origin, []byte(o.Target), []byte(o.With), -1)), nil
}

// satisfy text/transform
// type tf struct {
// 	ovr Overrider
// }
//
// func NewTF(overrider Overrider) tf {
// 	return tf{
// 		ovr: overrider,
// 	}
// }
//
// func (o *Override) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
// 	dst = bytes.Replace(src, []byte(o.Target), []byte(o.With), -1)
// 	return len(dst), len(src), nil
// }
//
// func (_ *Override) Reset() {}
