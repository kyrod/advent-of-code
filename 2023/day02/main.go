package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() []string {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(f), "\n")
}

func main() {
	lines := readInput()
	ans1 := part1(lines)
	fmt.Printf("Part 1: %v\n", ans1)
	ans2 := part2(lines)
	fmt.Printf("Part 2: %v\n", ans2)
}

var bagCubes = map[string]int{
	red:   12,
	green: 13,
	blue:  14,
}

const (
	red   = "red"
	green = "green"
	blue  = "blue"
)

func parseGame(line string) (int, map[string]int) {
	maxCubes := map[string]int{
		red:   0,
		green: 0,
		blue:  0,
	}
	spl := strings.Split(line, ":")
	gameNum, _ := strconv.Atoi(strings.TrimPrefix(spl[0], "Game "))

	for _, roll := range strings.Split(spl[1], ";") {
		roll = strings.TrimSpace(roll)

		for _, color := range strings.Split(roll, ",") {
			color = strings.TrimSpace(color)
			cSplit := strings.Split(color, " ")
			numString, c := cSplit[0], cSplit[1]

			num, _ := strconv.Atoi(numString)
			maxCubes[c] = max(maxCubes[c], num)
		}
	}
	return gameNum, maxCubes
}

func part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		gameNum, maxCubes := parseGame(line)
		possible := true
		for k, v := range bagCubes {
			if maxCubes[k] > v {
				possible = false
			}
		}
		if possible {
			sum += gameNum
		}
	}
	return sum
}

func part2(lines []string) int {
	sum := 0
	for _, line := range lines {
		_, maxCubes := parseGame(line)
		x := 1
		for _, v := range maxCubes {
			x *= v
		}
		sum += x
	}
	return sum
}
