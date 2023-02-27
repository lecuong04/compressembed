package lib

import (
	"bytes"
	"compress/flate"
	"compress/zlib"
	"io"
)

func Compress(data, key []byte) []byte {
	var buf bytes.Buffer
	w, err := zlib.NewWriterLevelDict(&buf, flate.BestCompression, key)
	if err != nil {
		return nil
	}
	defer w.Close()
	_, _ = w.Write(data)
	return buf.Bytes()
}

func Decompress(data, key []byte) []byte {
	var buf bytes.Buffer
	r, err := zlib.NewReaderDict(bytes.NewReader(data), key)
	if err != nil {
		return nil
	}
	_, _ = io.Copy(&buf, r)
	r.Close()
	return buf.Bytes()
}
