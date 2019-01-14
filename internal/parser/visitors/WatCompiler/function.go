package wast

import (
	"bytes"
	"fmt"
	"io"
)

type locals struct {
	name string
	tp   string
}

type function struct {
	code      *bytes.Buffer
	locals    []locals
	export    bool
	tblIndex  int
	name      string
	signature string
	mangle    string
}

func (f function) writeTo(w io.Writer) error {
	w.Write([]byte("("))
	w.Write([]byte(f.signature))
	w.Write([]byte("\n"))
	for _, local := range f.locals {
		w.Write([]byte(fmt.Sprintf("(local %s %s)\n", local.name, local.tp)))
	}
	f.code.WriteTo(w)
	w.Write([]byte(")\n"))
	if f.export {
		data := []byte(fmt.Sprintf("(export \"%s\" (func %s))\n", f.name, f.mangle))
		w.Write(data)
	}
	w.Write([]byte(fmt.Sprintf("(elem (i32.const %d) %s)\n", f.tblIndex, f.mangle)))
	return nil
}
