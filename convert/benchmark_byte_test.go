package convert

import (
	"bytes"
	"strings"
	"testing"
)

var str = strings.Repeat("test", 1024)
var bs = bytes.Repeat([]byte("test"), 1024)

func BenchmarkStringToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = StringToBytes(str)
	}
}

func BenchmarkSTDStringToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = []byte(str)
	}
}

func BenchmarkBytesToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = BytesToString(bs)
	}
}

func BenchmarkSTDBytesToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string(bs)
	}
}
