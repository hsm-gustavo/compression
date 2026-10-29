package rle

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/hsm-gustavo/compression/files"
)

func TestRLELit(t *testing.T) {
	inputs := [][]byte{
		[]byte(""),
		[]byte("A"),
		[]byte("ABCDEFGH"),
		[]byte("AAAAAAA"),
		[]byte("AAABCDFFFFFGH"),
		bytes.Repeat([]byte("A"), 300),
		files.Lorem,
	}

	for _, input := range inputs {
		compressed := Compress(input)
		decompressed := Decompress(compressed)

		originalSize := len(input)
		compressedSize := len(compressed)
		ratio := float64(len(compressed)) / float64(len(input))
		reductionPercent := (float64(1) - ratio) * 100

		fmt.Println("Original size:", originalSize, "bytes")
		fmt.Println("Compression size:", compressedSize, "bytes")
		fmt.Printf("Reduction: %.2f%%\n", reductionPercent)
		if !bytes.Equal(input, decompressed) {
			t.Errorf(
				"round-trip failed for %q: got %q",
				input,
				decompressed,
			)
		}
	}
}
