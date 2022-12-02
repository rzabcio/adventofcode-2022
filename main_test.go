package main

import (
	"testing"
)

func TestDay01(t *testing.T) {
	got, want := Day01_1("input-files/day01-test1.txt"), 24000
	if got != want {
		t.Errorf("Day01_1(test1) = %d; want %d", got, want)
	}
	got, want = Day01_2("input-files/day01-test1.txt"), 45000
	if got != want {
		t.Errorf("Day01_2(test1) = %d; want %d", got, want)
	}
}

func TestDay02(t *testing.T) {
	got, want := Day02_1("input-files/day02-test1.txt"), 15
	if got != want {
		t.Errorf("Day02_1(test1) = %d; want %d", got, want)
	}
	got, want = Day02_2("input-files/day02-test1.txt"), 12
	if got != want {
		t.Errorf("Day02_2(test1) = %d; want %d", got, want)
	}
}

func TestDay03(t *testing.T) {
	got, want := Day03_1("input-files/day03-test1.txt"), 157
	if got != want {
		t.Errorf("Day03_1(test1) = %d; want %d", got, want)
	}
	got, want = Day03_2("input-files/day03-test1.txt"), 70
	if got != want {
		t.Errorf("Day03_2(test1) = %d; want %d", got, want)
	}
}

func TestDay04(t *testing.T) {
	got, want := Day04_1("input-files/day04-test1.txt"), 2
	if got != want {
		t.Errorf("Day04_1(test1) = %d; want %d", got, want)
	}
	got, want = Day04_2("input-files/day04-test1.txt"), 0
	if got != want {
		t.Errorf("Day04_2(test1) = %d; want %d", got, want)
	}
}
