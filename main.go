package main

import (
	"os"
)

type Run struct {
	Count byte
	Value byte
}

func main() {
	data, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	if len(data) == 0 {
		return
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

	err = os.WriteFile("output.rle", runs, 0644)
	if err != nil {
		panic(err)
	}
}
