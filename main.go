package main

import (
	"fmt"
	"os"

	"github.com/hsm-gustavo/compression/rle_lit"
)

func main() {
	data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}

	compressed := rle_lit.Compress(data)

	decompressed := rle_lit.Decompress(compressed)

	fmt.Println("Original: ", data)
	fmt.Println("Compressed: ", compressed)
	fmt.Println("Decompressed: ", decompressed)
	//fmt.Println(bytes.Equal(data, decompressed)) // should be true
}
