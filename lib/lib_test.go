package lib_test

import (
	"testing"

	"github.com/lecuong04/compressembed/lib"
)

var text = `Hello World!`

func TestTool(t *testing.T) {
	key := []byte(lib.KeyGen())
	com := lib.Compress([]byte(text), key)
	store := lib.Decompress(com, key)
	t.Error(text, "\t", string(store), "\t", string(key))
}
