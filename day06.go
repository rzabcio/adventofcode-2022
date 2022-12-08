package main

import (
	"fmt"
	// "regexp"
	// "strconv"
	// "strings"
	// "github.com/thoas/go-funk"
)

func Day06_1(filename string) (result int) {
	for _, rs := range loadFile(filename) {
		result = rs.findMarker(4)
		fmt.Printf("06.1 ==> marker for message '%s' is '%s' on %d position\n", rs.message, rs.marker, result)
	}
	return
}

func Day06_2(filename string) (result int) {
	fmt.Printf("06.1 ==> %d\n", result)
	return
}

func loadFile(filename string) (rss []RadioSequence) {
	for line := range inputCh(filename) {
		rss = append(rss, *NewRadioSequence(line))
	}
	return rss
}

type RadioSequence struct {
	message string
	marker  string
}

func NewRadioSequence(line string) (rs *RadioSequence) {
	rs = &RadioSequence{
		message: line,
	}
	return rs
}

func (rs *RadioSequence) findMarker(length int) (pos int) {
	rs.marker = ""
	for _, char := range rs.message {
		indexOfChar := lastIndexOfRune(rs.marker, char)
		if indexOfChar > -1 {
			newMarker := ""
			for i := indexOfChar + 1; i < len(rs.marker); i++ {
				newMarker += string(rs.marker[i])
			}
			// fmt.Printf("  next letter: '%s' -> cut '%s' from %d to %d: %s\n", string(char),rs.marker,indexOfChar,len(rs.marker),newMarker,)
			rs.marker = newMarker
		}
		rs.marker += string(char)
		pos++
		// fmt.Printf("- %d: %s -> %s\n", pos, string(char), rs.marker)
		if len(rs.marker) == length {
			return pos
		}
	}
	return pos
}
