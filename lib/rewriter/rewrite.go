package rewrite

import "io"

type Rewrite struct {
	Target string
	With   string
}

type Reader struct {
	pos    int
	cur    int
	target []byte
	with   []byte
	r      io.Reader
}

func (rw *Rewrite) NewReader(r io.Reader) *Reader {
	return &Reader{
		pos:    0,
		cur:    0,
		target: []byte(rw.Target),
		with:   []byte(rw.With),
		r:      r,
	}
}

func (r *Reader) Read(p []byte) (int, error) {
	flag := r.target[0]
	max := len(p)
}
