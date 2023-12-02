package main

import "testing"

const (
	answer1 = 54667
	answer2 = 54203
)

// goos: darwin
// goarch: arm64
// pkg: github.com/kyrod/advent-of-code/2023/day01
// Benchmark/part1-10         	1000000000	         0.0001600 ns/op
// Benchmark/part2-10         	1000000000	         0.002654 ns/op
// Benchmark/part1-goroutines-10         	1000000000	         0.0006572 ns/op
// Benchmark/part2-goroutines-10         	1000000000	         0.0008865 ns/op
// PASS
// ok  	github.com/kyrod/advent-of-code/2023/day01	0.442s
func Benchmark(b *testing.B) {
	inputLines := readInput()

	for i := 0; i < b.N; i++ {
		var ans int
		b.Run("part1", func(b *testing.B) {
			ans = solve(inputLines, false)
		})
		if ans != answer1 {
			b.Errorf("got %d, want %d", ans, answer1)
		}
	}

	for i := 0; i < b.N; i++ {
		var ans int
		b.Run("part2", func(b *testing.B) {
			ans = solve(inputLines, true)
		})
		if ans != answer2 {
			b.Errorf("got %d, want %d", ans, answer2)
		}
	}

	// goroutines benchmark
	for i := 0; i < b.N; i++ {
		var ans int
		b.Run("part1-goroutines", func(b *testing.B) {
			ans = solveGoroutine(inputLines, false)
		})
		if ans != answer1 {
			b.Errorf("got %d, want %d", ans, answer1)
		}
	}

	for i := 0; i < b.N; i++ {
		var ans int
		b.Run("part2-goroutines", func(b *testing.B) {
			ans = solveGoroutine(inputLines, true)
		})
		if ans != answer2 {
			b.Errorf("got %d, want %d", ans, answer2)
		}
	}
}
