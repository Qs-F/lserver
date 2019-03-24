package lserver

import (
	"net/http"
	"os"
	"time"
)

type File struct {
	o   http.File
	ovr Overrider
}

func NewFile(origin http.File, overrider Overrider) *File {
	return &File{
		o:   origin,
		ovr: overrider,
	}
}

func (f *File) Close() error {
	return f.o.Close()
}

func (f *File) Read(p []byte) (int, error) {
	r, err := f.ovr.Map(f.o)
	if err != nil {
		return 0, err
	}
	return r.Read(p)
}

func (f *File) Seek(offset int64, whence int) (int64, error) {
}

func (f *File) Readdir(count int) ([]os.FileInfo, error) {
	return f.o.Readdir(count)
}

func (f *File) Stat() (os.FileInfo, error) {
	ofi, err := f.o.Stat()
	if err != nil {
		return nil, err
	}
	fi := NewFileInfo(ofi)
}

type FileInfo struct {
	o   os.FileInfo
	ovr Overrider
}

func NewFileInfo(origin os.FileInfo, overrider Overrider) *FileInfo {
	return &FileInfo{
		o:   origin,
		ovr: overrider,
	}
}

func (fi *FileInfo) Size() int64 {
	r, err := fi.ovr.Map(fi.o)
	if err != nil {

	b, err := ioutil.ReadAll(r)
}

func (fi *FileInfo) Name() string { return fi.o.Name() }

func (fi *FileInfo) Mode() os.FileMode { return fi.o.Mode() }

func (fi *FileInfo) ModTime() time.Time { return fi.o.ModTime() }

func (fi *FileInfo) IsDir() bool { return fi.o.IsDir() }

func (fi *FileInfo) Sys() interface{} { return fi.o.Sys() }

func (fi *FileInfo) MapFileInfo(orig os.FileInfo) (os.FileInfo, error) {
}
