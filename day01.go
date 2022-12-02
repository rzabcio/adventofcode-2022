package main

import (
	"fmt"
	// "regexp"
	"strconv"
	// "strings"

	"github.com/thoas/go-funk"
)

func Day01_1(filename string) (result int) {
	f := NewFellowship(filename)
	_, result = minMax(f.totals)
	fmt.Printf("01.1 ==> backpack with most rations has %d calories\n", result)
	return
}

func Day01_2(filename string) (result int) {
	f := NewFellowship(filename)
	totals := make([]int, len(f.totals))
	var max int
	copy(totals, f.totals)

	for range []int{0, 1, 2} {
		_, max = minMax(totals)
		totals = removeInt(totals, max)
		result += max
	}

	fmt.Printf("01.2 ==> three best backpacks have %d calories\n", result)
	return
}

type Fellowship struct {
	backpacks []Backpack
	totals    []int
}

func NewFellowship(filename string) Fellowship {
	f := new(Fellowship)
	f.backpacks = make([]Backpack, 0)
	f.totals = make([]int, 0)
	backpack := *new(Backpack)

	for line := range inputCh(filename) {
		if len(line) == 0 {
			f.backpacks = append(f.backpacks, backpack)
			f.totals = append(f.totals, backpack.sum())
			backpack = *new(Backpack)
		} else {
			ration, _ := strconv.Atoi(line)
			backpack.rations = append(backpack.rations, ration)
		}
	}
	f.backpacks = append(f.backpacks, backpack)
	f.totals = append(f.totals, backpack.sum())
	return *f
}

type Backpack struct {
	rations []int
}

func (b *Backpack) sum() (r int) {
	r = funk.Reduce(b.rations, '+', 0).(int)
	return
}
