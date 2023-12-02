package main

import "testing"

// goos: darwin
// goarch: arm64
// pkg: github.com/kyrod/advent-of-code/2023/day01
// Benchmark
// Benchmark/part1
// Benchmark/part1-10         	1000000000	         0.0000911 ns/op
// Benchmark/part2
// Benchmark/part2-10         	1000000000	         0.002177 ns/op
// PASS
func Benchmark(b *testing.B) {
	inputLines := readInput()

	for i := 0; i < b.N; i++ {
		b.Run("part1", func(b *testing.B) {
			part1(inputLines)
		})
	}

	for i := 0; i < b.N; i++ {
		b.Run("part2", func(b *testing.B) {
			part2(inputLines)
		})
	}
}
