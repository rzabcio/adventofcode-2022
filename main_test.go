package main

import (
	"testing"
)

func TestDay1(t *testing.T) {
	got, want := Day1_1("input-files/day01-test1.txt"), 24000
	if got != want {
		t.Errorf("Day1_1(test1) = %d; want %d", got, want)
	}
	got, want = Day1_2("input-files/day01-test1.txt"), 45000
	if got != want {
		t.Errorf("Day1_2(test1) = %d; want %d", got, want)
	}
}

func TestDay2(t *testing.T) {
	got, want := Day2_1("input-files/day02-test1.txt"), 24000
	if got != want {
		t.Errorf("Day2_1(test1) = %d; want %d", got, want)
	}
	got, want = Day2_2("input-files/day02-test1.txt"), 45000
	if got != want {
		t.Errorf("Day2_2(test1) = %d; want %d", got, want)
	}
}
