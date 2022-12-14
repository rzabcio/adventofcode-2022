package main

import (
	"fmt"
	// "regexp"
	// "strconv"
	"strings"
	// "github.com/thoas/go-funk"
)

func Day02_1(filename string) (result int) {
	s := NewStrategy(filename)
	for _, r := range s.rounds {
		r.calculateResult()
		// fmt.Printf("- %v => result: %d, score: %d\n", r, r.result, r.myScore())
		result += r.myScore()
	}
	fmt.Printf("02.1 ==> my total score: %d\n", result)
	return
}

func Day02_2(filename string) (result int) {
	s := NewStrategy(filename)
	for _, r := range s.rounds {
		r.calculateMyMove()
		// fmt.Printf("- %v => result: %d, score: %d\n", r, r.result, r.myScore())
		result += r.myScore()
	}
	fmt.Printf("02.2 ==> my total score: %d\n", result)
	return
}

type Strategy struct {
	rounds []Round
}

func NewStrategy(filename string) *Strategy {
	s := new(Strategy)
	s.rounds = make([]Round, 0)

	for line := range inputCh(filename) {
		s.rounds = append(s.rounds, *NewRound(line))
	}
	return s
}

type Round struct {
	line   string
	their  int
	my     int
	result int
}

func NewRound(line string) *Round {
	r := new(Round)
	r.line = line
	return r
}

func (r *Round) decodeTheir() {
	split := strings.Split(r.line, " ")
	switch split[0] {
	case "A":
		r.their = 1
	case "B":
		r.their = 2
	case "C":
		r.their = 3
	}
}

func (r *Round) decodeMy() {
	split := strings.Split(r.line, " ")
	switch split[1] {
	case "X":
		r.my = 1
	case "Y":
		r.my = 2
	case "Z":
		r.my = 3
	}
}

func (r *Round) decodeResult() {
	split := strings.Split(r.line, " ")
	switch split[1] {
	case "X":
		r.result = 0
	case "Y":
		r.result = 3
	case "Z":
		r.result = 6
	}
}

func (r *Round) calculateResult() {
	r.decodeTheir()
	r.decodeMy()
	if r.my-r.their == 1 {
		r.result = 6
	} else if r.my-r.their == 0 {
		r.result = 3
	} else if r.my-r.their == -2 {
		r.result = 6
	} else {
		r.result = 0
	}
}

func (r *Round) calculateMyMove() {
	r.decodeTheir()
	r.decodeResult()
	if r.result == 6 && r.their == 3 {
		r.my = 1
	} else if r.result == 6 {
		r.my = r.their + 1
	} else if r.result == 3 {
		r.my = r.their
	} else if r.their == 1 {
		r.my = 3
	} else {
		r.my = r.their - 1
	}
}

func (r *Round) myScore() int {
	return r.result + r.my
}
