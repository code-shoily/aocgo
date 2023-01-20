// Package day07 - Solution for Advent of Code 2016/07
// Problem Link: http://adventofcode.com/2016/day/07
package day07

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	ips := parse(input)
	return solvePart1(ips), solvePart2(ips)
}

func solvePart1(ips []IPv7) (total int) {
	for _, ip := range ips {
		if ip.supportsTLS() {
			total++
		}
	}

	return total
}

func solvePart2(ips []IPv7) (total int) {
	for _, ip := range ips {
		if ip.supportsSSL() {
			total++
		}
	}

	return total
}

func parse(input string) (ips []IPv7) {
	for _, line := range strings.Split(input, "\n") {
		supernets, hypernets := parseLine(line)
		ips = append(ips, IPv7{supernets, hypernets})
	}

	return ips
}

func parseLine(line string) (supernets []string, hypernets []string) {
	var segment []rune

	for _, c := range line {
		switch c {
		case '[':
			supernets = append(supernets, string(segment))
			segment = segment[:0]
		case ']':
			hypernets = append(hypernets, string(segment))
			segment = segment[:0]
		default:
			segment = append(segment, c)
		}
	}

	supernets = append(supernets, string(segment))

	return supernets, hypernets
}

type IPv7 struct {
	supernets []string
	hypernets []string
}

func (ip IPv7) supportsTLS() bool {
	var valid bool
	for _, supernet := range ip.supernets {
		if hasABBA(supernet) {
			valid = true
			break
		}
	}

	for _, hypernet := range ip.hypernets {
		if hasABBA(hypernet) {
			return false
		}
	}

	return valid
}

func (ip IPv7) supportsSSL() bool {
	babSet := map[string]bool{}

	for _, supernet := range ip.supernets {
		for k, v := range getABAtoBAB(supernet) {
			babSet[k] = v
		}
	}

	for _, hypernet := range ip.hypernets {
		from, to := 0, 2
		for to < len(hypernet) {
			if hypernet[from] == hypernet[to] && hypernet[from+1] != hypernet[to] {
				if _, exists := babSet[hypernet[from:to+1]]; exists {
					return true
				}
			}
			from++
			to++
		}
	}
	return false
}

func getABAtoBAB(sequence string) map[string]bool {
	set := map[string]bool{}
	from, to := 0, 2

	for to < len(sequence) {
		if sequence[from] == sequence[to] && sequence[from+1] != sequence[to] {
			bab := string([]byte{sequence[from+1], sequence[to], sequence[from+1]})
			set[bab] = true
		}
		from++
		to++
	}

	return set
}

func hasABBA(sequence string) bool {
	from, to := 0, 3

	for to < len(sequence) {
		if sequence[from] == sequence[to] && sequence[from+1] == sequence[to-1] && sequence[from] != sequence[from+1] {
			return true
		}
		to++
		from++
	}

	return false
}
