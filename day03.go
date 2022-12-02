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
	fmt.Printf("03.1 ==> sum of invalid items: %d\n", result)
	return
}

func Day03_2(filename string) (result int) {
	rs := NewRucksacks(filename)
	result = rs.badgesVal()
	fmt.Printf("03.1 ==> sum of badges of all groups: %d\n", result)
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

func (rs *Rucksacks) badgesVal() (result int) {
	for i := 0; i < len(rs.rucksacks)-2; i = i + 3 {
		inter := funk.IntersectString(rs.rucksacks[i].items, rs.rucksacks[i+1].items)
		inter = funk.IntersectString(inter, rs.rucksacks[i+2].items)
		inter = funk.UniqString(inter)
		result += stringToItemValue(inter[0])
	}
	return result
}

type Sack struct {
	items      []string
	l          []string
	r          []string
	invalid    []string
	invalidVal int
}

func NewSack(line string) (s *Sack) {
	s = new(Sack)
	s.items = strings.Split(line, "")
	halfLen := len(s.items) / 2
	s.l = s.items[:halfLen:halfLen]
	s.r = s.items[halfLen:len(s.items):len(s.items)]
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
