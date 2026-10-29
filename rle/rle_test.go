package rle

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/hsm-gustavo/compression/files"
)

func TestRLE(t *testing.T) {
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

		fmt.Println("Original size:", originalSize, "bytes")
		fmt.Println("Compression size:", compressedSize, "bytes")

		if originalSize > 0 {
			ratio := float64(len(compressed)) / float64(len(input))
			reductionPercent := (float64(1) - ratio) * 100
			fmt.Printf("Reduction: %.2f%%\n", reductionPercent)
		} else {
			fmt.Println("Reduction: N/A")
		}
		if !bytes.Equal(input, decompressed) {
			t.Errorf(
				"round-trip failed for %q: got %q",
				input,
				decompressed,
			)
		}
	}
}

func TestBinaryRLE(t *testing.T) {
	input := []byte{
		0x00, 0x00, 0x00, 0x00,
		0xFF, 0xFF, 0xFF, 0xFF,
		0x01, 0x02, 0x03,
		0x00, 0x00,
	}

	compressed := Compress(input)
	decompressed := Decompress(compressed)

	originalSize := len(input)
	compressedSize := len(compressed)

	fmt.Println("Original size:", originalSize, "bytes")
	fmt.Println("Compression size:", compressedSize, "bytes")

	if originalSize > 0 {
		ratio := float64(len(compressed)) / float64(len(input))
		reductionPercent := (float64(1) - ratio) * 100
		fmt.Printf("Reduction: %.2f%%\n", reductionPercent)
	} else {
		fmt.Println("Reduction: N/A")
	}

	if !bytes.Equal(input, decompressed) {
		t.Fatal("binary round-trip failed")
	}
}
