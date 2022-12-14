package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

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

func parseData(data []string) ([][]int, []Coords, Coords) {
	var (
		landscape   [][]int
		startPoints []Coords
		endPoint    Coords
	)

	for row, d := range data {
		var m []int
		for col, l := range d {
			if string(l) == "S" || string(l) == "a" {
				m = append(m, int('a'))
				startPoints = append(startPoints, Coords{row, col})
			} else if string(l) == "E" {
				m = append(m, int('z'))
				endPoint.row = row
				endPoint.col = col
			} else {
				m = append(m, int(l))
			}
		}
		landscape = append(landscape, m)
	}

	return landscape, startPoints, endPoint
}

type Coords struct {
	row int
	col int
}

func getNeighbours(c Coords, landscape [][]int) []Coords {
	var validCoords []Coords
	up := Coords{c.row - 1, c.col}
	right := Coords{c.row, c.col + 1}
	down := Coords{c.row + 1, c.col}
	left := Coords{c.row, c.col - 1}

	if isCoordContained(up, landscape) {
		validCoords = append(validCoords, up)
	}
	if isCoordContained(right, landscape) {
		validCoords = append(validCoords, right)
	}
	if isCoordContained(down, landscape) {
		validCoords = append(validCoords, down)
	}
	if isCoordContained(left, landscape) {
		validCoords = append(validCoords, left)
	}

	return validCoords
}

func isCoordContained(c Coords, landscape [][]int) bool {
	return c.row >= 0 && c.col >= 0 && c.row < len(landscape) && c.col < len(landscape[0])
}

func isMoveValid(from Coords, to Coords, landscape [][]int) bool {
	if !isCoordContained(from, landscape) || !isCoordContained(to, landscape) {
		return false
	}
	fromVal := landscape[from.row][from.col]
	toVal := landscape[to.row][to.col]
	return toVal <= fromVal+1
}

func initializeDistances(landscape [][]int) [][]int {
	var distances [][]int

	for row := range landscape {
		var d []int
		for range landscape[row] {
			d = append(d, -1)
		}
		distances = append(distances, d)
	}

	return distances
}

func printDistances(distances [][]int) {
	for _, d := range distances {
		fmt.Println(d)
	}
}

func isExploring(queue []Coords) bool {
	return len(queue) > 0
}

func explore(queue []Coords, distances [][]int, landscape [][]int) ([][]int, []Coords) {
	coordToExplore := queue[0]
	newQueue := queue[1:]

	neighbours := getNeighbours(coordToExplore, landscape)

	for _, neighbour := range neighbours {
		isVisited := distances[neighbour.row][neighbour.col]
		isValid := isMoveValid(coordToExplore, neighbour, landscape)

		if isVisited == -1 && isValid {
			distances[neighbour.row][neighbour.col] = distances[coordToExplore.row][coordToExplore.col] + 1
			newQueue = append(newQueue, neighbour)
		}
	}

	return distances, newQueue
}

func partOne(startPoint Coords, endPoint Coords, landscape [][]int) int {
	var (
		queue       []Coords
		endDistance int
	)
	distances := initializeDistances(landscape)
	distances[startPoint.row][startPoint.col] = 0
	queue = append(queue, startPoint)

	for isExploring(queue) {
		distances, queue = explore(queue, distances, landscape)
		endDistance = distances[endPoint.row][endPoint.col]

		if endDistance != -1 {
			break
		}
	}

	return endDistance
}

func partTwo(startPoints []Coords, endPoint Coords, landscape [][]int) int {
	var distances []int

	for _, startPoint := range startPoints {
		endDistance := partOne(startPoint, endPoint, landscape)
		if endDistance != -1 {
			distances = append(distances, endDistance)
		}
	}

	sort.Ints(distances)
	return distances[0]
}

func main() {
	data := readFile("input.txt")

	landscape, startPoints, endPoint := parseData(data)
	fmt.Println(partOne(startPoints[0], endPoint, landscape))
	fmt.Println(partTwo(startPoints, endPoint, landscape))
}
