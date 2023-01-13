// Package day04 - Solution for Advent of Code 2018/04
// Problem Link: http://adventofcode.com/2018/day/04
package day04

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const (
	asleep = -1
	awake  = -2
)

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	guards := parse(input)
	return solvePart1(guards), solvePart2(guards)
}

func solvePart1(guards map[int]Guard) int {
	maxSleepTime := math.MinInt
	var superSleeper Guard

	for _, guard := range guards {
		if guard.totalSleepTime() > maxSleepTime {
			maxSleepTime = guard.totalSleepTime()
			superSleeper = guard
		}
	}

	minuteAsleepMost, _ := superSleeper.minuteAsleepMost()
	return minuteAsleepMost * superSleeper.id
}

func solvePart2(guards map[int]Guard) int {
	var (
		minuteAsleepMost int
		habitualSleeper  Guard
		maxFrequency     = math.MinInt
	)

	for _, guard := range guards {
		minute, frequency := guard.minuteAsleepMost()
		if frequency > maxFrequency {
			maxFrequency = frequency
			minuteAsleepMost = minute
			habitualSleeper = guard
		}
	}

	return minuteAsleepMost * habitualSleeper.id
}

func parse(input string) map[int]Guard {
	events := parseEvents(input)

	guards := map[int]Guard{}
	var currentGuard, sleep int

	for _, event := range events {
		switch event.eventType {
		case asleep:
			sleep = event.minute
		case awake:
			guard := guards[currentGuard]
			guard.sleep(sleep, event.minute)
			sleep = 0
		default:
			guardId := event.eventType
			if _, ok := guards[guardId]; ok {
				currentGuard = guardId
			} else {
				guards[guardId] = Guard{id: guardId, minutesAsleep: make([]int, 60)}
				currentGuard = guardId
			}
		}
	}

	return guards
}

func parseEvents(input string) (events []Event) {
	for _, eventLine := range strings.Split(input, "\n") {
		events = append(events, parseEvent(eventLine))
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].date < events[j].date
	})

	return events
}

func parseEvent(eventLine string) Event {
	eventLine = stripBrackets(eventLine)
	tokens := strings.SplitN(eventLine, " ", 3)

	return Event{
		fmt.Sprintf("%s:%s", tokens[0], tokens[1]),
		extractMinute(tokens[1]),
		parseEventType(tokens[2]),
	}
}

func parseEventType(rawEventType string) int {
	switch rawEventType {
	case "wakes up":
		return awake
	case "falls asleep":
		return asleep
	default:
		var guard int
		fmt.Sscanf(rawEventType, "Guard #%d begins shift", &guard)
		return guard
	}
}

func stripBrackets(line string) string {
	withoutBrackets := strings.FieldsFunc(line, func(r rune) bool {
		return r == '[' || r == ']'
	})

	return strings.Join(withoutBrackets, "")
}

func extractMinute(timeStr string) int {
	timeValues := strings.Split(timeStr, ":")
	minute, _ := strconv.Atoi(timeValues[1])

	return minute
}

type Event struct {
	date      string
	minute    int
	eventType int
}

type Guard struct {
	id            int
	minutesAsleep []int
}

func (guard *Guard) sleep(from, to int) {
	for i := from; i < to; i++ {
		guard.minutesAsleep[i]++
	}
}

func (guard *Guard) totalSleepTime() (total int) {
	for _, freq := range guard.minutesAsleep {
		total += freq
	}

	return total
}

func (guard *Guard) minuteAsleepMost() (sleepyMinute int, minutesAsleep int) {
	for minute, tally := range guard.minutesAsleep {
		if tally > minutesAsleep {
			sleepyMinute = minute
			minutesAsleep = tally
		}
	}

	return sleepyMinute, minutesAsleep
}
