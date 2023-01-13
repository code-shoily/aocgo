// Package day04 - Solution for Advent of Code 2020/04
// Problem Link: http://adventofcode.com/2020/day/04
package day04

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	data := parse(input)
	return solvePart1(data), solvePart2(data)
}

func solvePart1(data [][][2]string) (valid int) {
	for _, info := range data {
		if isValidPassportV1(info) {
			valid++
		}
	}

	return valid
}

func solvePart2(data [][][2]string) (valid int) {
	for _, info := range data {
		passport := NewPassport(info)
		if passport.isValid() {
			valid++
		}
	}

	return valid
}

func parse(input string) (data [][][2]string) {
	for _, info := range strings.Split(input, "\n\n") {
		data = append(data, parsePassportInfo(info))
	}

	return data
}

func parsePassportInfo(info string) (fields [][2]string) {
	kvs := strings.FieldsFunc(info, func(token rune) bool {
		return token == '\n' || token == ' '
	})

	for _, kv := range kvs {
		kvInfo := strings.Split(kv, ":")
		fields = append(fields, [2]string{kvInfo[0], kvInfo[1]})
	}

	return fields
}

func isValidPassportV1(info [][2]string) bool {
	fieldInfo := map[string]bool{
		"byr": true,
		"iyr": true,
		"eyr": true,
		"hgt": true,
		"hcl": true,
		"ecl": true,
		"pid": true,
		"cid": true,
	}

	for _, field := range info {
		delete(fieldInfo, field[0])
	}

	return len(fieldInfo) == 0 || len(fieldInfo) == 1 && fieldInfo["cid"]
}

type Passport struct {
	byr int    // (Birth Year)
	iyr int    // (Issue Year)
	eyr int    // (Expiration Year)
	hgt string // (Height)
	hcl string // (Hair Color)
	ecl string // (Eye Color)
	pid string // (Passport ID)
	cid string // (Country ID)
}

func NewPassport(kvs [][2]string) *Passport {
	var byr, iyr, eyr int
	var hgt, hcl, ecl, pid, cid string
	for _, field := range kvs {
		switch field[0] {
		case "byr":
			if value, error := strconv.Atoi(field[1]); error == nil {
				byr = value
			}
		case "iyr":
			if value, error := strconv.Atoi(field[1]); error == nil {
				iyr = value
			}
		case "eyr":
			if value, error := strconv.Atoi(field[1]); error == nil {
				eyr = value
			}
		case "hgt":
			hgt = field[1]
		case "hcl":
			hcl = field[1]
		case "ecl":
			ecl = field[1]
		case "pid":
			pid = field[1]
		case "cid":
			cid = field[1]
		}
	}

	return &Passport{byr, iyr, eyr, hgt, hcl, ecl, pid, cid}
}

func (p Passport) isValid() bool {
	return p.validateBirthYear() && p.validateIssueYear() && p.validateExpirationYear() &&
		p.validateHeight() && p.validateHairColor() && p.validateEyeColor() && p.validatePid()
}

func (p Passport) validateBirthYear() bool {
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	return p.byr >= 1920 && p.byr <= 2002
}

func (p Passport) validateIssueYear() bool {
	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	return p.iyr >= 2010 && p.iyr <= 2020

}

func (p Passport) validateExpirationYear() bool {
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	return p.eyr >= 2020 && p.eyr <= 2030
}

func (p Passport) validateHeight() bool {
	//hgt (Height) - a number followed by either cm or in:
	//If cm, the number must be at least 150 and at most 193.
	//If in, the number must be at least 59 and at most 76.
	var height int
	var unit string
	fmt.Sscanf(p.hgt, "%d%s", &height, &unit)

	switch unit {
	case "cm":
		return height >= 150 && height <= 193
	case "in":
		return height >= 59 && height <= 76
	default:
		return false
	}
}

func (p Passport) validateEyeColor() bool {
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	for _, color := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if color == p.ecl {
			return true
		}
	}

	return false
}

func (p Passport) validateHairColor() bool {
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	var colour string
	fmt.Sscanf(p.hcl, "#%s", &colour)

	if _, error := strconv.ParseInt(colour, 16, 64); error == nil {
		return len(colour) == 6
	}
	return false
}

func (p Passport) validatePid() bool {
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	if _, error := strconv.Atoi(p.pid); error == nil {
		return len(p.pid) == 9
	}
	return false
}
