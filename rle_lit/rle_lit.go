package rle_lit

import (
	"bytes"
	"os"
)

type Run struct {
	Marker byte   // 00 repeated 01 literal
	Count  byte   // how many literal/how many reps
	Values []byte // if repeated its a 1 value list, if literal its a list of values
}

func makeRun(count byte, values []byte) Run {
	if count > 3 {
		// append repeated
		// marker, count repeat, byte to repeat
		return Run{
			Marker: byte(0),
			Count:  count,
			Values: []byte{values[0]},
		}
	}

	// append literal
	// marker, count many, literal list
	return Run{
		Marker: byte(1),
		Count:  count,
		Values: values,
	}
}

func Compress(data []byte) []byte {
	if len(data) == 0 {
		return []byte{}
	}

	// counting process
	var runs []Run

	currentCount := byte(1)
	currentByte := data[0]
	// data    = aaaab
	// indexes = 01234
	// slice   = aaab
	// indexes = 0123
	//
	// count starts equals 1
	// so when we reach b:
	// b = "b"
	// idx = 3 (idx is the index inside the slice)
	// count = 4 (starts at one and increases one each
	// iter if b == currentByte)
	// start := 3 + 1 - 4 = 0
	// values := data[0:3+1] = data[0:4]
	// since slices second arg is exclusive, it returns
	// me an array from index 0 to 3 of the original array
	// so values = {aaaa}
	for idx, b := range data[1:] {
		if b == currentByte && currentCount < 255 {
			currentCount++
			continue
		}

		start := idx + 1 - int(currentCount) // literal slice idx start
		values := data[start : idx+1]

		runs = append(runs, makeRun(currentCount, values))
		currentByte = b
		currentCount = 1
	}

	// last one
	start := len(data) - int(currentCount)
	values := data[start:]
	runs = append(runs, makeRun(currentCount, values))

	/* err := os.WriteFile("output.rle", runs, 0644)
	if err != nil {
		panic(err)
	}

	return runs */

	markers := len(runs)
	count := len(runs)
	var val int
	for _, r := range runs {
		val += len(r.Values)
	}
	// fixed array capacity
	output := make([]byte, 0, markers+count+val)
	for _, r := range runs {
		output = append(output, r.Marker, r.Count)
		output = append(output, r.Values...)
	}

	err := os.WriteFile("output.rle", output, 0644)
	if err != nil {
		panic(err)
	}

	return output
}

func Decompress(data []byte) []byte {
	if len(data) == 0 {
		return []byte{}
	}

	var runs []Run

	limit := len(data)

	for i := 0; i < limit; i++ {
		marker := data[i]
		count := data[i+1]
		var values []byte

		if count <= 3 {
			values = append(values, data[i+2:i+2+int(count)]...)
			// i+2 = assumption that we are always ending this after the last value
			// i + 2 + count = same as above, ends after the count, since its exclusive that value wont get in
			// value is marker
			i += 2 + int(count) - 1 // for loop already adds 1
		} else {
			values = append(values, data[i+2])
			i += 3 - 1 // {marker count value marker ...} so indexes 0 1 2 3, we want to end the loop on a marker
			// same as before, loop already adds 1 so we remove 1
		}

		runs = append(runs, Run{
			Marker: marker,
			Count:  count,
			Values: values,
		})
	}

	var capacity int
	for _, r := range runs {
		capacity += int(r.Count)
	}
	output := make([]byte, 0, capacity)

	for _, r := range runs {
		if r.Count > 3 {
			output = append(output, bytes.Repeat([]byte{r.Values[0]}, int(r.Count))...)
		} else {
			output = append(output, r.Values...)
		}
	}

	return output
}
