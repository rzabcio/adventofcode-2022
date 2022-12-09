package main

import (
	"fmt"
	// "regexp"
	"strconv"
	"strings"
	// "github.com/thoas/go-funk"
)

func Day07_1(filename string) (result int) {
	dc := NewDirectoryCrawler(filename)
	// fmt.Printf("directory crawler:\n%v\n", dc.Print())
	// fmt.Printf("flat dir list:\n%v\n", dc.root.DirsFlat())
	dirs, result := dc.dirsSizeLeq(100000)
	fmt.Printf("07.1 ==> dirs with size not greater than 100000 (totaling to %d): %v\n", result, dirs)
	return result
}

func Day07_2(filename string) (result int) {
	dc := NewDirectoryCrawler(filename)
	sizeNeeded := 30000000 - (70000000 - dc.root.SizeR())
	dirs, result := dc.dirsSizeGeq(sizeNeeded)
	fmt.Printf("07.2 ==> dirs with size greater than %d (smallest's size %d): %v\n", sizeNeeded, result, dirs)
	return result
}

type DirectoryCrawler struct {
	root    *Directory
	current *Directory
}

func NewDirectoryCrawler(filename string) (dc DirectoryCrawler) {
	dc = DirectoryCrawler{}
	dc.root = NewDirectory("")
	dc.current = dc.root
	for line := range inputCh(filename) {
		if strings.HasPrefix(line, "$ ") {
			dc.processCommand(line[2:])
		} else {
			dc.processContent(line)
		}
	}
	return dc
}

func (dc *DirectoryCrawler) processCommand(line string) {
	// fmt.Printf("|> processing command: %s\n", line)
	if line == "cd .." {
		dc.current = dc.current.parent
	} else if strings.HasPrefix(line, "cd ") {
		newDir, ok := dc.current.dirs[line[3:]]
		if ok {
			dc.current = newDir
		}
	}
}

func (dc *DirectoryCrawler) processContent(line string) {
	// fmt.Printf("-> processing content: %s\n", line)
	if strings.HasPrefix(line, "dir ") {
		dir := NewDirectory(line[4:])
		dir.parent = dc.current
		dc.current.dirs[dir.name] = dir
	} else {
		split := strings.Split(line, " ")
		size, _ := strconv.Atoi(split[0])
		dc.current.files[split[1]] = size
	}
}

func (dc *DirectoryCrawler) dirsFlat() (dirs []*Directory) {
	return dc.root.DirsFlat()
}

func (dc *DirectoryCrawler) dirsSizeLeq(maxSize int) (dirs []*Directory, totalSize int) {
	for _, dir := range dc.dirsFlat() {
		size := dir.SizeR()
		if size <= maxSize {
			dirs = append(dirs, dir)
			totalSize += size
		}
	}
	return dirs, totalSize
}

func (dc *DirectoryCrawler) dirsSizeGeq(maxSize int) (dirs []*Directory, smallestSize int) {
	for _, dir := range dc.dirsFlat() {
		size := dir.SizeR()
		if size >= maxSize {
			dirs = append(dirs, dir)
			if smallestSize == 0 {
				smallestSize = size
			} else if smallestSize > size {
				smallestSize = size
			}
		} else {
			// fmt.Printf("    %s size: %d\n", dir.name, size)
		}
	}
	return dirs, smallestSize
}

func (dc DirectoryCrawler) Print() (s string) {
	return "- / (dir)\n" + dc.root.Print()
}

type Directory struct {
	name   string
	parent *Directory
	dirs   map[string]*Directory
	files  map[string]int
}

func NewDirectory(name string) (dir *Directory) {
	dir = &Directory{}
	dir.name = name
	dir.dirs = make(map[string]*Directory)
	dir.files = make(map[string]int)
	return dir
}

func (dir Directory) Pwd() (pwd string) {
	if dir.parent == nil {
		pwd = "/" // root
	} else {
		pwd = fmt.Sprintf("%s%s/", dir.parent.Pwd(), dir.name)
	}
	return pwd
}

func (dir Directory) Level() (level int) {
	if dir.parent == nil {
		level = 0 // root
	} else {
		level = dir.parent.Level() + 1
	}
	return level
}

func (dir Directory) String() (s string) {
	return dir.name
}

func (dir Directory) Print() (s string) {
	indent := ""
	for i := 0; i <= dir.Level(); i++ {
		indent += "  "
	}
	for _, subdir := range dir.dirs {
		s += fmt.Sprintf("%s- %s (dir)\n", indent, subdir.name)
		s += subdir.Print()
	}
	for filename, filesize := range dir.files {
		s += fmt.Sprintf("%s- %s (file, size=%d)\n", indent, filename, filesize)
	}
	return s
}

func (dir Directory) Size() (size int) {
	for _, filesize := range dir.files {
		size += filesize
	}
	return size
}

func (dir Directory) SizeR() (size int) {
	size = dir.Size()
	for _, subdir := range dir.dirs {
		size += subdir.SizeR()
	}
	return size
}

func (dir Directory) DirsFlat() (subdirs []*Directory) {
	subdirs = make([]*Directory, 0)
	for _, subdir := range dir.dirs {
		subdirs = append(subdirs, subdir)
		subdirs = append(subdirs, subdir.DirsFlat()...)
	}
	return subdirs
}
