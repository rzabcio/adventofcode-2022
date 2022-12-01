package main

import (
	"fmt"
	// "regexp"
	"strconv"
	// "strings"

	"github.com/thoas/go-funk"
)

func Day1_1(filename string) (result int) {
	f := NewFellowship(filename)
	_, result = minMax(f.totals)
	fmt.Printf("==> backpack with most rations has %d calories\n", result)
	return
}

func Day1_2(filename string) (result int) {
	// f := NewFellowship(filename)

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
