package bufio

import(
	"bufio"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ScanBytes(t *testing.T) {
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