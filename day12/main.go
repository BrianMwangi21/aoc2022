package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// https://www.youtube.com/watch?v=xcIUM003HS0

type Coords struct {
	row int
	col int
}

func (c Coords) GetNeighbours() (Coords, Coords, Coords, Coords) {
	up := Coords{c.row - 1, c.col}
	right := Coords{c.row, c.col + 1}
	down := Coords{c.row + 1, c.col}
	left := Coords{c.row, c.col - 1}
	return up, right, down, left
}

func isCoordContained(landscape [][]int) bool {
	return false
}

func isMoveValid(from Coords, to Coords) bool {
	return false
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

func parseData(data []string) ([][]int, Coords, Coords) {
	var (
		landscape  [][]int
		startPoint Coords
		endPoint   Coords
	)

	for row, d := range data {
		var m []int
		for col, l := range d {
			m = append(m, int(l))
			if string(l) == "S" {
				startPoint.row = row
				startPoint.col = col
			} else if string(l) == "E" {
				endPoint.row = row
				endPoint.col = col
			}
		}
		landscape = append(landscape, m)
	}

	return landscape, startPoint, endPoint
}

func partOne(landscape [][]int, startPoint Coords, endPoint Coords) int {
	leastSteps := 0

	for _, d := range landscape {
		fmt.Println(d)
	}

	return leastSteps
}

func main() {
	data := readFile("input.test.txt")

	landscape, startPoint, endPoint := parseData(data)
	fmt.Println(partOne(landscape, startPoint, endPoint))
}
