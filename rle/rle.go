package rle

import (
	"bytes"
	"os"
)

func Compress(data []byte) []byte {
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

func Decompress(data []byte) []byte {
	if len(data) == 0 {
		return []byte{}
	}

	if len(data)%2 != 0 {
		panic("the compressed file should have pairs of bytes")
	}

	var runs []byte

	limit := len(data)

	for i := 0; i < limit; i += 2 {
		count := data[i]
		value := data[i+1]

		runs = append(runs, bytes.Repeat([]byte{value}, int(count))...)
	}

	return runs
}
