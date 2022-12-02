package main

import (
	"fmt"
	"regexp"
	// "strings"
	// "regexp"
	"strconv"
	// "strings"
	"github.com/thoas/go-funk"
)

func Day04_1(filename string) (result int) {
	bc := NewBeachCleaning(filename)
	result = bc.countFullyOverlappingTeams()
	fmt.Printf("04.1 ==> number of fully overlapping cleaning tems: %d\n", result)
	return
}

func Day04_2(filename string) (result int) {
	bc := NewBeachCleaning(filename)
	result = bc.countAllOverlappingTeams()
	fmt.Printf("04.1 ==> number of all overlaping cleaning teams: %d\n", result)
	return
}

type BeachCleaning struct {
	teams []CleaningTeam
}

func NewBeachCleaning(filename string) (bc *BeachCleaning) {
	bc = new(BeachCleaning)
	bc.teams = make([]CleaningTeam, 0)

	for line := range inputCh(filename) {
		bc.teams = append(bc.teams, *NewCleaningTeam(line))
	}
	return bc
}

func (bc *BeachCleaning) countFullyOverlappingTeams() (result int) {
	for _, ct := range bc.teams {
		intersection := funk.Intersect(ct.s1, ct.s2).([]int)
		diff1, diff1_ := funk.DifferenceInt(ct.s1, intersection)
		diff2, diff2_ := funk.DifferenceInt(ct.s2, intersection)
		if (len(diff1) == 0 && len(diff1_) == 0) || (len(diff2) == 0 && len(diff2_) == 0) {
			result++
		}
	}
	return result
}

func (bc *BeachCleaning) countAllOverlappingTeams() (result int) {
	for _, ct := range bc.teams {
		intersection := funk.Intersect(ct.s1, ct.s2).([]int)
		if len(intersection) > 0 {
			result++
		}
	}
	return result
}

type CleaningTeam struct {
	s1 []int
	s2 []int
}

func NewCleaningTeam(line string) (ct *CleaningTeam) {
	ct = new(CleaningTeam)
	ct.s1 = make([]int, 0)
	ct.s2 = make([]int, 0)
	regex, _ := regexp.Compile("[0-9]*")
	nos := regex.FindAllString(line, -1)
	s1from, _ := strconv.Atoi(nos[0])
	s1to, _ := strconv.Atoi(nos[1])
	s2from, _ := strconv.Atoi(nos[2])
	s2to, _ := strconv.Atoi(nos[3])

	for i := s1from; i <= s1to; i++ {
		ct.s1 = append(ct.s1, i)
	}
	for i := s2from; i <= s2to; i++ {
		ct.s2 = append(ct.s2, i)
	}
	return ct
}
