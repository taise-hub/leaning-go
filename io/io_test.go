package io

import (
	"io"
	"os"
	"fmt"
	"time"
	"testing"
	"strings"
)

func Test_Copy(t *testing.T) {
	type args struct {
		dst io.Writer
		src io.Reader
	}
	tests := map[string]struct {
		args args
	}{
		"標準出力": {
			args: args{
				src: strings.NewReader("some io.Reader stream to be read\n"),
				dst: os.Stdout,
			},
		},
	}
	for tName, test := range tests {
		t.Run(tName, func(t *testing.T) {
			if _, err := io.Copy(test.args.dst, test.args.src); err != nil {
				panic(err)
			}
		})
	}
}

func Test_CopyBuffer(t *testing.T) {//Copyとほぼ変わらない。bufferを自分で用意する点が違う？
	type args struct {
		dst io.Writer
		src io.Reader
		buf []byte
	}
	tests := map[string]struct {
		args args
	}{
		"一般的な使い方": {
			args{
				src: strings.NewReader("first reader\n"),
				dst: os.Stdout,
				buf: make([]byte, 1),
			},
		},
	}
	for tName, test := range tests {
		t.Run(tName, func(t *testing.T) {
			if _, err := io.CopyBuffer(test.args.dst, test.args.src, test.args.buf); err != nil{
				panic(err)
			}
		})
	}
}

func Test_CopyN(t *testing.T) {
	type args struct {
		dst io.Writer
		src io.Reader
		n 	int64
	}
	tests := map[string]struct{
		args args
	}{
		"使い方": {
			args: args{
				src: strings.NewReader("some io.Reader stream to be read"),
				dst: os.Stdout,
				n: 14,
			},
		},
	}
	for tName, test := range tests {
		t.Run(tName, func(t *testing.T) {
			if _, err := io.CopyN(test.args.dst, test.args.src, test.args.n); err != nil {
				panic(err)
			}
		})
	}

}

func Test_Pipe(t *testing.T) {
	r, w := io.Pipe()
	go func() {
		time.Sleep(3000 * time.Millisecond)
		fmt.Fprintf(w, "some io.Reader stream to be read\n")
		w.Close()
	}()
	if _, err := io.Copy(os.Stdout, r); err != nil {
		panic(err)
	}
} 