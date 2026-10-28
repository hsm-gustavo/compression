package main

import (
	"bytes"
	"fmt"
	"os"
)

func compress(data []byte) []byte {
	if len(data) == 0 {
		return []byte{}
	}

	// counting process
	var runs []byte

	currentCount := byte(1)
	currentByte := data[0]

	for _, b := range data[1:] {
		if b == currentByte && currentCount < 255 {
			currentCount++
			continue
		}

		runs = append(runs, currentCount, currentByte)

		currentByte = b
		currentCount = 1
	}

	// last one
	runs = append(runs, currentCount, currentByte)

	err := os.WriteFile("output.rle", runs, 0644)
	if err != nil {
		panic(err)
	}

	return runs
}

func decompress(data []byte) []byte {
	return []byte{}
}

func main() {
	data, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	compressed := compress(data)

	decompressed := decompress(compressed)

	fmt.Println(bytes.Equal(data, decompressed)) // should be true
}
