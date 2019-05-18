package main

import (
	"io"
	"testing"
)

var (
	w   io.Writer
	err error
)

func init()  {
	if w, err = NewLog(); err != nil {
		panic(err)
	}
}

func BenchmarkFLog_Write(b *testing.B) {
	for i:=0;i<b.N;i++{
		_,err = w.Write([]byte([]byte("hello world")))
		if err != nil{
			b.Error(err)
		}
		_,err = w.Write([]byte("\n"))
		if err != nil{
			b.Error(err)
		}
	}
}
