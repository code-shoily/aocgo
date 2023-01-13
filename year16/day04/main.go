// Package day04 - Solution for Advent of Code 2016/04
// Problem Link: http://adventofcode.com/2016/day/04
package day04

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const northPoleObjectStorage = "northpole object storage"

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	data := parse(input)
	return solvePart1(data), solvePart2(data)
}

func solvePart1(rooms []Room) (total int) {
	for _, room := range rooms {
		if room.isReal() {
			total += room.sectorId
		}
	}

	return total
}

func solvePart2(rooms []Room) int {
	for _, room := range rooms {
		if room.isReal() && room.decipherRoomName() == northPoleObjectStorage {
			return room.sectorId
		}
	}

	panic("Can't store anything to north pole!!!")
}

func parse(input string) (data []Room) {
	for _, line := range strings.Split(input, "\n") {
		data = append(data, newRoom(line))
	}

	return data
}

type Room struct {
	encryptedName string
	sectorId      int
	checksum      string
}

func (r Room) isReal() bool {
	frequencyMap := getFrequencies(r.encryptedName)
	descendingFrequencies := getDescendingFrequencies(frequencyMap)
	previousChar := rune(r.checksum[0])
	previousFrequency := frequencyMap[previousChar]

	for idx, char := range r.checksum {
		frequency := frequencyMap[char]

		if descendingFrequencies[idx] != frequency {
			return false
		}

		if idx != 0 {
			if previousFrequency < frequency || previousFrequency == frequency && char < previousChar {
				return false
			}

			previousChar = char
			previousFrequency = frequency
		}
	}

	return true
}

func (r Room) decipherRoomName() string {
	nameSlice := make([]rune, len(r.encryptedName))

	for idx, char := range r.encryptedName {
		nameSlice[idx] = rotate(char, r.sectorId)
	}

	return string(nameSlice)
}

func getFrequencies(name string) map[rune]int {
	frequencies := map[rune]int{}
	for _, char := range name {
		frequencies[char]++
	}

	delete(frequencies, '-')

	return frequencies
}

func getDescendingFrequencies(frequencies map[rune]int) []int {
	values := make([]int, len(frequencies))

	for _, v := range frequencies {
		values = append(values, v)
	}

	sort.SliceStable(values, func(i, j int) bool {
		return values[i] > values[j]
	})

	return values
}

func newRoom(line string) Room {
	tokens := strings.Split(line, "-")
	encryptedName := tokens[:len(tokens)-1]
	meta := strings.FieldsFunc(tokens[len(tokens)-1], func(r rune) bool {
		return r == '[' || r == ']'
	})

	sectorId, error := strconv.Atoi(meta[0])

	if error != nil {
		panic("Invalid parsing logic")
	}

	return Room{strings.Join(encryptedName, "-"), sectorId, meta[1]}
}

func rotate(char rune, by int) rune {
	if char != '-' {
		if offset := char + rune(by%26); offset > 'z' {
			return offset%'z' + 'a' - 1
		} else {
			return offset
		}
	}
	return ' '
}
