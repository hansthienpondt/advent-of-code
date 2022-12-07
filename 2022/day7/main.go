package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type File struct {
	size int    // the size of the cur file
	name string // the name of the cur file
}

func (f *File) Size() int {
	return f.size
}
func (f *File) String() string {
	return f.name
}
func (f *File) IsDir() bool {
	return false
}
func (f *File) IsFile() bool {
	return true
}

type FileSystem struct {
	name        string        // the name of the cur directory
	files       []*File       // list of files in directory
	directories []*FileSystem // list of childs in directory
	prev        *FileSystem   // pointer to parent
}

func (fs *FileSystem) Size() int {
	size := 0
	for _, f := range fs.files {
		size += f.Size()
	}
	for _, d := range fs.directories {
		size += d.Size()
	}
	return size
}
func (fs *FileSystem) String() string {
	return fs.name
}
func (fs *FileSystem) IsDir() bool {
	return true
}
func (fs *FileSystem) IsFile() bool {
	return false
}

func NewRoot() *FileSystem {
	return &FileSystem{name: "/", files: []*File{}, directories: []*FileSystem{}}
}
func NewDir(name string, parent *FileSystem) *FileSystem {
	return &FileSystem{name: name, files: []*File{}, directories: []*FileSystem{}, prev: parent}
}
func NewFile(name string, size int) *File {
	return &File{name: name, size: size}
}

func readInput(file string) []string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	data := string(bytes)
	lines := strings.Split(data, "\n")

	return lines
}
func Parse(input []string) *FileSystem {
	var root *FileSystem = NewRoot()
	var curdir *FileSystem = root
	for _, l := range input {
		// Parse commands
		if strings.HasPrefix(l, "$") {
			switch {
			case strings.HasPrefix(l, "$ cd /"):
				//fmt.Println(curdir)
				continue
			case strings.HasPrefix(l, "$ cd .."):
				//fmt.Println("traversing up", l)
				curdir = curdir.prev
			case strings.HasPrefix(l, "$ cd"):
				dir := strings.Split(l, "$ cd ")[1]
				//fmt.Println("CHANGING DIRS", dir)
				for _, lsdir := range curdir.directories {
					if lsdir.name == dir {
						curdir = lsdir
					}
				}
			case strings.HasPrefix(l, "$ ls"):
				//fmt.Println("listing files", l)
				continue
			default:
				fmt.Println("non-recognized command")
			}
		} else {
			// parse files/dirs
			switch {
			case strings.HasPrefix(l, "dir "):
				dir := strings.Split(l, "dir ")[1]
				curdir.directories = append(curdir.directories, NewDir(dir, curdir))
				//fmt.Println("DIR", dir)
			default:
				size, _ := strconv.Atoi(strings.Split(l, " ")[0])
				fname := strings.Split(l, " ")[1]
				curdir.files = append(curdir.files, NewFile(fname, size))
				//fmt.Println("FILE", fname)
			}
		}
	}
	return root
}

func Count7a(dir *FileSystem, limit int, total *int) {
	if len(dir.directories) > 0 || len(dir.files) > 0 {
		if dir.Size() <= limit {
			*total += dir.Size()
		}
	}
	for _, subdir := range dir.directories {
		Count7a(subdir, limit, total)
	}
}
func Count7b(dir *FileSystem, spaceNeeded int, minSize *int) {
	if len(dir.directories) > 0 || len(dir.files) > 0 {
		if dir.Size() >= spaceNeeded && dir.Size() < *minSize {
			*minSize = dir.Size()
		}
	}
	for _, subdir := range dir.directories {
		Count7b(subdir, spaceNeeded, minSize)
	}
}

func day7a(input []string) int {
	root := Parse(input)
	count := 0
	for _, i := range root.directories {
		Count7a(i, 100000, &count)
	}
	return count
}

func day7b(input []string) int {
	root := Parse(input)
	freeSpace := 70000000 - root.Size()
	spaceNeeded := 30000000 - freeSpace
	minSize := 1<<32 - 1
	for _, i := range root.directories {
		Count7b(i, spaceNeeded, &minSize)
	}

	return minSize
}

func main() {
	filename := os.Args[1]

	fmt.Printf("Part 1: %d\n", day7a(readInput(filename)))
	fmt.Printf("Part 2: %d\n", day7b(readInput(filename)))

}
