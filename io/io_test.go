package io

import (
	"io"
	"os"
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
		expected int64
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