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

	startTime := 0
	subTimeLength := 1

	for i, entry := range dialogueEntries {
		timestamp := fmt.Sprintf("00:00:%02d,000 --> 00:00:%02d,000", startTime, startTime+subTimeLength)
		fmt.Fprintln(file, strconv.Itoa(i+1))
		fmt.Fprintln(file, timestamp)
		fmt.Fprintln(file, entry)
		fmt.Fprintln(file)
		startTime += subTimeLength
	}
	return filename, nil
}
