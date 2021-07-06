package bufio

import(
	"bufio"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ScanBytes(t *testing.T) {//byte列の最初の文字を取得する関数
	type args struct {
		data []byte
		atEOF bool
	}
	tests := map[string]struct {
		args args
		expected []byte
	} {
		"byte列の最初の値を取得できる。": {
			args: args{
				data: []byte{0x40, 0x41, 0x42, 0x43, 0x44},
				atEOF: true,
			},
			expected: []byte{0x40},
		},

		"文字列をbyteに変換した時に最初の文字のbyteが取得できる。": {
			args: args{
				data : []byte("hello"),
				atEOF: true,
			},
			expected: []byte("h"),
		},
	}
	for tName, test := range tests {
		t.Run(tName, func(t *testing.T) {
			advance, sut, err := bufio.ScanBytes(test.args.data, test.args.atEOF)
			if err != nil {
				panic(err)
			}
			assert.Equal(t, advance, 1)
			assert.Equal(t, test.expected, sut)
		})
	}
}


func Test_ScanLines(t *testing.T) {
	type args struct {
		data []byte
		atEOF bool
	}
	tests := map[string]struct {
		args args
		expected []byte
	}{
		"一行をdataとして与えると全て取得する。": {
			args: args{
				data: []byte("Hello, World."),
				atEOF: true,
			},
			expected: []byte("Hello, World."),
		},
		"二行をdataとして与えると最初の1行を取得する。": {
			args: args{
				data: []byte("Hello,\nWorld."),
				atEOF: true,
			},
			expected: []byte("Hello,"),
		},
		"raw文字列で改行をすると最初の1行を取得できる。": {
			args: args{
				data: []byte(
					`Hello,
					World`),
				atEOF: true,
			},
			expected: []byte("Hello,"),
		},
	}
	for tName, test := range tests {
		t.Run(tName, func(t *testing.T) {
			_, sut, err := bufio.ScanLines(test.args.data, test.args.atEOF)
			if err != nil {
				panic(err)
			}		
			assert.Equal(t, test.expected, sut)
		})
	}
}

func Test_ScanRune(t *testing.T) {
	type args struct {
		data []byte
		atEOF bool
	}
	tests := map[string]struct {
		args args
		expected []byte
	}{
		"英字": {
			args: args{
				data: []byte("Hello, World."),
				atEOF: true,
			},
			expected: []byte("H"),
		},

		"数字": {
			args: args{
				data : []byte("1234"),
				atEOF: true,
			},
			expected: []byte("1"),
		},
		"日本語": {
			args: args{
				data : []byte("こんにちは"),
				atEOF: true,
			},
			expected: []byte("こ"),
		},
		"絵文字": {
			args: args{
				data : []byte("😭これ"),
				atEOF: true,
			},
			expected: []byte("😭"),
		},
	}
	for tName, test := range tests {
		t.Run(tName, func(t *testing.T) {
			_, sut, err := bufio.ScanRunes(test.args.data, test.args.atEOF)
			if err != nil {
				panic(err)
			}
			assert.Equal(t, test.expected, sut)
		})
	}
}