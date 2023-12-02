package main

import "testing"

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
