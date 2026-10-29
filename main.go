package main

import (
	"bytes"
	"fmt"

	"github.com/hsm-gustavo/compression/files"
	"github.com/hsm-gustavo/compression/rle"
)

func main() {
	compressed := rle.Compress(files.Lorem)

	decompressed := rle.Decompress(compressed)

	fmt.Println("Original: ", files.Lorem)
	fmt.Println("Compressed: ", compressed)
	fmt.Println("Decompressed: ", decompressed)
	fmt.Println(bytes.Equal(files.Lorem, decompressed)) // should be true
}
