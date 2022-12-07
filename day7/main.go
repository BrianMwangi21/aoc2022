package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Files map[string]*File

type File struct {
	Name     string
	Size     int
	Children Files
}

func readFile(filename string) []string {
	var data []string

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func getFileStructure(data []string) *File {
	current := &File{Name: "/", Children: Files{}}
	root := current

	for _, d := range data[1:] {
		switch {
		case d[:4] == "$ cd":
			name := d[5:]
			current = current.Children[name]
		case d[:3] == "dir":
			name := d[4:]
			current.Children[name] = &File{
				Name:     name,
				Children: Files{"..": current},
			}
		case '0' <= d[0] && d[0] <= '9':
			split := strings.Split(d, " ")
			size, _ := strconv.Atoi(split[0])
			name := split[1]
			current.Children[name] = &File{
				Name: name,
				Size: size,
			}
		}
	}

	return root
}

func computeSizes(dir *File) int {
	if len(dir.Children) == 0 {
		return dir.Size
	}

	size := 0
	for k, file := range dir.Children {
		if k == ".." {
			continue
		}
		size += computeSizes(file)
	}
	dir.Size = size

	return size
}

func differenceCalc(dir *File, total *int, spaceNeeded int, minSize *int) {
	if len(dir.Children) > 0 {
		if dir.Size <= 100000 {
			*total += dir.Size
		}

		if dir.Size >= spaceNeeded && dir.Size < *minSize {
			*minSize = dir.Size
		}
	}

	for k, file := range dir.Children {
		if k == ".." {
			continue
		}
		differenceCalc(file, total, spaceNeeded, minSize)
	}
}

func partOneAndTwo(data []string) (int, int) {
	root := getFileStructure(data)
	computeSizes(root)

	total := 0

	freeSpace := 70000000 - root.Size
	spaceNeeded := 30000000 - freeSpace
	minSize := 1<<32 - 1

	differenceCalc(root, &total, spaceNeeded, &minSize)

	return total, minSize
}

func main() {
	data := readFile("input.txt")

	partOne, partTwo := partOneAndTwo(data)

	fmt.Println(partOne)
	fmt.Println(partTwo)
}
