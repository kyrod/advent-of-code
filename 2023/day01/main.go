package main

import (
	"fmt"
	"os"
	_ "runtime/pprof"
	"strconv"
	"strings"
	"unicode"
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
	fmt.Printf("Part 1: %d\n", ans1)
	ans2 := part2(lines)
	fmt.Printf("Part 2: %d\n", ans2)
}

func part1(lines []string) int {
	return solve(lines, false)
}

func part2(lines []string) int {
	return solve(lines, true)
}

func solve(lines []string, findTextNumbers bool) int {
	sum := 0
	for _, line := range lines {
		var first, last int
		if line == "" {
			continue
		}
		for i := range line {
			if num, isNumber := lineToNumber(line, i, findTextNumbers); isNumber {
				if first == 0 {
					first = num
				}
				last = num
			}
		}
		value, _ := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
		sum += value
	}
	return sum
}

var textNumMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func lineToNumber(line string, i int, findTextNums bool) (int, bool) {
	if unicode.IsNumber(rune(line[i])) {
		num, _ := strconv.Atoi(string(line[i]))
		return num, true
	}
	// part 2
	if findTextNums {
		for word, num := range textNumMap {
			if strings.HasPrefix(line[i:], word) {
				return num, true
			}
		}
	}
	return 0, false
}
