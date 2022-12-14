package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

var (
	minX, maxX, maxY int
	field            map[string]rune
	sand             [2]int
	startX           = 500
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

func crd(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func parseData(data []string) {
	field = make(map[string]rune)
	minX = math.MaxInt

	for _, line := range data {
		parts := strings.Split(line, " -> ")

		for i := 1; i < len(parts); i++ {
			a, b := parts[i-1], parts[i]

			var sx, sy, ex, ey int
			fmt.Sscanf(a, "%d,%d", &sx, &sy)
			fmt.Sscanf(b, "%d,%d", &ex, &ey)

			minX = min(sx, min(ex, minX))
			maxX = max(sx, max(ex, maxX))
			maxY = max(sy, max(ey, maxY))

			if sx == ex {
				for y := min(sy, ey); y <= max(sy, ey); y++ {
					field[crd(sx, y)] = '#'
				}
			} else if sy == ey {
				for x := min(sx, ex); x <= max(sx, ex); x++ {
					field[crd(x, sy)] = '#'
				}
			} else {
				panic("Non-straight line asked for!")
			}
		}
	}
}

func partOne() int {
	sand = [2]int{startX, 1}
	stopped := 0

	for sand[1] <= maxY {
		var next = sand
		var ok bool
		next[1] += 1

		if _, ok = field[crd(next[0], next[1])]; !ok {
			sand = next
		} else if _, ok = field[crd(next[0]-1, next[1])]; !ok {
			sand = [2]int{next[0] - 1, next[1]}
		} else if _, ok = field[crd(next[0]+1, next[1])]; !ok {
			sand = [2]int{next[0] + 1, next[1]}
		} else {
			field[crd(sand[0], sand[1])] = 'o'
			stopped++
			sand = [2]int{startX, 1}
		}
	}

	return stopped
}

func partTwo(stopped int) int {
	sand = [2]int{startX, 0}
	maxY += 2

	for {
		var next = sand
		var ok bool
		next[1] += 1
		didStop := false
		for i := -1; i <= 1; i++ {
			field[crd(next[0]+i, maxY)] = '#'
		}
		if _, ok = field[crd(next[0], next[1])]; !ok {
			sand = next
		} else if _, ok = field[crd(next[0]-1, next[1])]; !ok {
			sand = [2]int{next[0] - 1, next[1]}
		} else if _, ok = field[crd(next[0]+1, next[1])]; !ok {
			sand = [2]int{next[0] + 1, next[1]}
		} else {
			field[crd(sand[0], sand[1])] = 'o'
			didStop = true
		}

		if sand[0] < minX {
			minX = sand[0]
		} else if sand[0] > maxX {
			maxX = sand[0]
		}
		if didStop {
			stopped++
			if sand[0] == startX && sand[1] == 0 {
				break
			}
			sand = [2]int{startX, 0}
		}
	}

	return stopped
}

func main() {
	data := readFile("input.txt")

	parseData(data)
	fmt.Println(partOne())
	parseData(data)
	fmt.Println(partTwo(partOne()))
}
