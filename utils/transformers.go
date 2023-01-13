package utils

import (
	"strconv"
	"strings"
)

func SplitLines(lines string) []string {
	if lines == "" {
		return []string{}
	}
	return strings.Split(strings.TrimSpace(lines), "\n")
}

func SplitIntLines(lines string) (intLines []int) {
	for _, line := range SplitLines(lines) {
		if intLine, error := strconv.Atoi(line); error == nil {
			intLines = append(intLines, intLine)
		}
	}

	return intLines
}

func SplitBy(line string, by string) []string {
	if line == "" {
		return []string{}
	}
	return strings.Split(line, by)
}

func SplitByInts(line string, by string) (intWords []int) {
	for _, word := range SplitBy(line, by) {
		if intWord, error := strconv.Atoi(word); error == nil {
			intWords = append(intWords, intWord)
		}
	}

	return intWords
}
