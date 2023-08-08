package main

import (
	"fmt"
	"os"
	"strconv"
)

func createSubtitlesFile(filename string, dialogueEntries []string) (string, error) {
	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Define the pause durations
	pauseDurations := map[rune]int{
		'.': 450,
		',': 275,
		';': 400,
		':': 400,
		'?': 450,
		'!': 450,
		'—': 600,  // Assuming this is an em dash
		'(': 400,
		')': 400,
		'“': 400,  // Beginning quotation mark
		'”': 400,  // Ending quotation mark
	}

	/* pauseDurations := map[rune]int{
		'.': 700,
		',': 600,
		';': 800,
		':': 800,
		'?': 1000,
		'!': 1000,
		'—': 600,  // Assuming this is an em dash
		'(': 400,
		')': 400,
		'“': 400,  // Beginning quotation mark
		'”': 400,  // Ending quotation mark
	}*/

	startTime := 0
	defaultSubTimeLength := 275

	for i, entry := range dialogueEntries {
		subTimeLength := defaultSubTimeLength

		// Adjust the subTimeLength if there's a punctuation mark at the end
		if len(entry) > 0 {
			lastRune := rune(entry[len(entry)-1])
			if duration, exists := pauseDurations[lastRune]; exists {
				subTimeLength += duration
			}
		}

		endTime := startTime + subTimeLength
		timestamp := fmt.Sprintf("00:00:%02d,%03d --> 00:00:%02d,%03d", startTime/1000, startTime%1000, endTime/1000, endTime%1000)
		fmt.Fprintln(file, strconv.Itoa(i+1))
		fmt.Fprintln(file, timestamp)
		fmt.Fprintln(file, entry)
		fmt.Fprintln(file)

		startTime = endTime
	}
	return filename, nil

}
