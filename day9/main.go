package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(x int) int {
	if x < 0 {
		return -1
	} else if x == 0 {
		return 0
	}
	return 1
}

func followTheLeader(head Coordinate, tail Coordinate) Coordinate {
	deltaX := head.X - tail.X
	deltaY := head.Y - tail.Y

	if abs(deltaX) > 1 || abs(deltaY) > 1 {
		return Coordinate{tail.X + sign(deltaX), tail.Y + sign(deltaY)}
	}

	return tail
}

func partOne(data []string) int {
	grid := make(map[Coordinate]bool)
	head := Coordinate{0, 0}
	tail := Coordinate{0, 0}
	grid[head] = true

	for _, d := range data {
		instructions := strings.Split(d, " ")
		direction := instructions[0]
		distance, _ := strconv.Atoi(instructions[1])

		for i := 0; i < distance; i++ {
			switch direction {
			case "U":
				head.Y++
			case "D":
				head.Y--
			case "L":
				head.X--
			case "R":
				head.X++
			}
			tail = followTheLeader(head, tail)
			grid[tail] = true
		}
	}

	return len(grid)
}

func partTwo(data []string) int {
	grid := make(map[Coordinate]bool)
	knotsXY := make([]Coordinate, 10)
	grid[Coordinate{0, 0}] = true

	for _, d := range data {
		instructions := strings.Split(d, " ")
		direction := instructions[0]
		distance, _ := strconv.Atoi(instructions[1])

		for i := 0; i < distance; i++ {
			switch direction {
			case "U":
				knotsXY[0].Y++
			case "D":
				knotsXY[0].Y--
			case "L":
				knotsXY[0].X--
			case "R":
				knotsXY[0].X++
			}

			for i := 1; i < 10; i++ {
				knotsXY[i] = followTheLeader(knotsXY[i-1], knotsXY[i])
			}

			grid[knotsXY[9]] = true
		}
	}

	return len(grid)
}

func main() {
	data := readFile("input.txt")

	fmt.Println(partOne(data))
	fmt.Println(partTwo(data))
}
