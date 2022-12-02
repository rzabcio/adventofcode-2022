package main

import (
	"fmt"
	// "regexp"
	// "strconv"
	"strings"

	"github.com/thoas/go-funk"
	// "github.com/thoas/go-funk"
)

func Day03_1(filename string) (result int) {
	rs := NewRucksacks(filename)
	for _, s := range rs.rucksacks {
		result += s.invalidVal
	}
	fmt.Printf("03.1 ==> sum of invalid items is %d\n", result)
	return
}

func Day03_2(filename string) (result int) {
	fmt.Printf("03.1 ==> %d\n", result)
	return
}

type Rucksacks struct {
	rucksacks []Sack
}

func NewRucksacks(filename string) (rs *Rucksacks) {
	rs = new(Rucksacks)
	rs.rucksacks = make([]Sack, 0)

	for line := range inputCh(filename) {
		rs.rucksacks = append(rs.rucksacks, *NewSack(line))
	}

	return rs
}

type Sack struct {
	l          []string
	r          []string
	invalid    []string
	invalidVal int
}

func NewSack(line string) (s *Sack) {
	s = new(Sack)
	split := strings.Split(line, "")
	halfLen := len(split) / 2
	s.l = split[:halfLen:halfLen]
	s.r = split[halfLen:len(split):len(split)]
	s.invalid = funk.UniqString(funk.IntersectString(s.l, s.r))
	s.invalidVal = stringToItemValue(s.invalid[0])
	return s
}

func stringToItemValue(item string) (val int) {
	val = int([]byte(item)[0])
	if val > 96 {
		val -= 96
	} else {
		val -= 38
	}
	return val
}
