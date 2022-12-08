package main

import (
	"fmt"
	// "regexp"
	"strconv"
	"strings"
	// "github.com/thoas/go-funk"
)

func Day05_1(filename string) (result int) {
	ls := NewLoadingSpace(filename)
	ls.Rearrange()
	fmt.Printf("05.1 ==> top crates: %v\n", ls.Tops())
	return
}

func Day05_2(filename string) (result int) {
	ls := NewLoadingSpace(filename)
	ls.newCrane = true
	ls.Rearrange()
	fmt.Printf("05.2 ==> top crates: %v\n", ls.Tops())
	return
}

type LoadingSpace struct {
	stacks   []CrateStack
	moves    [][]int
	newCrane bool
}

func NewLoadingSpace(filename string) (ls *LoadingSpace) {
	ls = &LoadingSpace{}
	ls.stacks = []CrateStack{}
	for line := range inputCh(filename) {
		if strings.Contains(line, "move ") {
			splitted := strings.Split(line, " ")
			m1, _ := strconv.Atoi(splitted[1])
			m2, _ := strconv.Atoi(splitted[3])
			m3, _ := strconv.Atoi(splitted[5])
			ls.moves = append(ls.moves, []int{m1, m2 - 1, m3 - 1})
		} else if len(line) > 0 {
			ls.stacks = append(ls.stacks, CrateStack{crates: strings.Split(line, "")})
		}
	}
	return ls
}

func (ls *LoadingSpace) Rearrange() {
	for _, move := range ls.moves {
		ls.Move(move[0], move[1], move[2])
	}
}

func (ls *LoadingSpace) Move(count, from, to int) {
	crates, ok := ls.stacks[from].PopCount(count)
	if !ok {
		fmt.Printf("---error---\n")
		return
	}
	if ls.newCrane {
		crates = reverseStrArr(crates)
	}
	ls.stacks[to].PushAll(crates)
}

func (ls *LoadingSpace) Tops() (tops string) {
	for _, stack := range ls.stacks {
		top, ok := stack.Top()
		if ok {
			tops += top
		} else {
			tops += "_"
		}
	}
	return tops
}

type CrateStack struct {
	crates []string
}

func (cs *CrateStack) Push(crate string) {
	cs.crates = append(cs.crates, crate)
}

func (cs *CrateStack) PushAll(crates []string) {
	for _, crate := range crates {
		cs.Push(crate)
	}
}

func (cs *CrateStack) Top() (crate string, ok bool) {
	if len(cs.crates) == 0 {
		return "", false
	}
	crate = cs.crates[len(cs.crates)-1]
	return crate, true
}

func (cs *CrateStack) Pop() (crate string, ok bool) {
	if len(cs.crates) == 0 {
		return "", false
	}
	crate = cs.crates[len(cs.crates)-1]
	cs.crates = cs.crates[:len(cs.crates)-1]
	return crate, true
}

func (cs *CrateStack) PopCount(count int) (crates []string, ok bool) {
	if len(cs.crates) < count {
		return nil, false
	}
	for i := 0; i < count; i++ {
		crate, ok := cs.Pop()
		if !ok {
			return nil, false
		}
		crates = append(crates, crate)
	}
	return crates, true
}
