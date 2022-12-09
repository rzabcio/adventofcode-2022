package main

import (
	"fmt"
	// "regexp"
	"strconv"
	"strings"
	// "github.com/thoas/go-funk"
)

func Day08_1(filename string) (result int) {
	tf := NewTreeForest(filename)
	tf.checkVisibility()
	result = tf.countVisible()
	fmt.Printf("08.1 ==> there are %d visible trees\n", result)
	return result
}

func Day08_2(filename string) (result int) {
	fmt.Printf("08.2 ==> %d\n", result)
	return result
}

type TreeForest struct {
	trees [][]Tree
}

type Tree struct {
	h    int
	nInv bool
	sInv bool
	wInv bool
	eInv bool
}

func NewTreeForest(filename string) (tf TreeForest) {
	tf = TreeForest{}
	for line := range inputCh(filename) {
		treeLine := make([]Tree, 0, len(line))
		for _, char := range strings.Split(line, "") {
			tree := Tree{nInv: false, sInv: false, wInv: false, eInv: false}
			tree.h, _ = strconv.Atoi(string(char))
			treeLine = append(treeLine, tree)
		}
		tf.trees = append(tf.trees, treeLine)
	}
	return tf
}

func (tf *TreeForest) checkVisibility() {
	for i := 0; i < len(tf.trees); i++ {
		heighest := 0
		for j := 0; j < len(tf.trees[i]); j++ {
			if j == 0 || heighest < tf.trees[i][j].h {
				heighest = tf.trees[i][j].h
			} else {
				tf.trees[i][j].wInv = true
			}
		}
		heighest = 0
		for j := len(tf.trees[i]) - 1; j >= 0; j-- {
			if j == len(tf.trees[i])-1 || heighest < tf.trees[i][j].h {
				heighest = tf.trees[i][j].h
			} else {
				tf.trees[i][j].eInv = true
			}
		}
	}

	for j := 0; j < len(tf.trees[0]); j++ {
		heighest := 0
		for i := 0; i < len(tf.trees); i++ {
			if i == 0 || heighest < tf.trees[i][j].h {
				heighest = tf.trees[i][j].h
			} else {
				tf.trees[i][j].nInv = true
			}
		}
		heighest = 0
		for i := len(tf.trees) - 1; i >= 0; i-- {
			if i == len(tf.trees)-1 || heighest < tf.trees[i][j].h {
				heighest = tf.trees[i][j].h
			} else {
				tf.trees[i][j].sInv = true
			}
		}
	}
}

func (tf TreeForest) countVisible() (i int) {
	for _, treeLine := range tf.trees {
		for _, tree := range treeLine {
			if !tree.IsInvisible() {
				i++
			}
		}
	}
	return i
}

func (tf TreeForest) Print() (s string) {
	for _, treeLine := range tf.trees {
		for _, tree := range treeLine {
			s += fmt.Sprintf("%s ", tree.Print())
		}
		s += "\n"
	}
	return s
}

func (t Tree) IsInvisible() (result bool) {
	return t.wInv && t.nInv && t.eInv && t.sInv
}

func (t Tree) Print() (s string) {
	inv := ""
	if t.wInv {
		inv += "_"
	} else {
		inv += "w"
	}
	if t.nInv {
		inv += "_"
	} else {
		inv += "n"
	}
	if t.eInv {
		inv += "_"
	} else {
		inv += "e"
	}
	if t.sInv {
		inv += "_"
	} else {
		inv += "s"
	}
	if t.IsInvisible() {
		// s = fmt.Sprintf("%d_(%s)", t.h, inv)
		s = fmt.Sprintf("%d_", t.h)
	} else {
		// s = fmt.Sprintf("%d*(%s)", t.h, inv)
		s = fmt.Sprintf("%d*", t.h)
	}
	return s
}
