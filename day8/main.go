package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
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

func getTreesGrid(data []string) [][]int {
	var grid [][]int

	for i := 0; i < len(data); i++ {
		var vals []int
		for j := 0; j < len(data[i]); j++ {
			val, _ := strconv.Atoi(string(data[i][j]))
			vals = append(vals, val)
		}
		grid = append(grid, vals)
	}

	return grid
}

func getEdgeTotal(grid [][]int) int {
	return (len(grid[0]) * 2) + ((len(grid) - 2) * 2)
}

func isTreeVisible(grid [][]int, tree int, row int, col int, ROW_HEIGHT int, COL_HEIGHT int) (bool, int) {
	toTop, toBottom, toLeft, toRight := true, true, true, true
	sceneTop, sceneBottom, sceneLeft, sceneRight, sceneScore := 0, 0, 0, 0, 0

	// To top
	for i := row - 1; i >= 0; i-- {
		if grid[i][col] >= tree {
			toTop = false
			sceneTop += 1
			break
		}
		sceneTop += 1
	}

	// To bottom
	for i := row + 1; i < ROW_HEIGHT; i++ {
		if grid[i][col] >= tree {
			toBottom = false
			sceneBottom += 1
			break
		}
		sceneBottom += 1
	}

	// To left
	for i := col - 1; i >= 0; i-- {
		if grid[row][i] >= tree {
			toLeft = false
			sceneLeft += 1
			break
		}
		sceneLeft += 1
	}

	// To right
	for i := col + 1; i < COL_HEIGHT; i++ {
		if grid[row][i] >= tree {
			toRight = false
			sceneRight += 1
			break
		}
		sceneRight += 1
	}

	if toTop || toBottom || toLeft || toRight {
		sceneScore = sceneTop * sceneLeft * sceneRight * sceneBottom
		return true, sceneScore
	}

	return false, sceneScore
}

func partOneAndTwo(data []string) (int, int) {
	total, max_scenic_score := 0, 0

	grid := getTreesGrid(data)

	ROW_HEIGHT := len(grid)
	COL_HEIGHT := len(grid[0])

	for i := 1; i < ROW_HEIGHT-1; i++ {
		for j := 1; j < COL_HEIGHT-1; j++ {
			tree := grid[i][j]
			visible, scenic_score := isTreeVisible(grid, tree, i, j, ROW_HEIGHT, COL_HEIGHT)

			if visible {
				total += 1
				max_scenic_score = int(math.Max(float64(max_scenic_score), float64(scenic_score)))
			}
		}
	}

	total += getEdgeTotal(grid)

	return total, max_scenic_score
}

func main() {
	data := readFile("input.txt")

	fmt.Println(partOneAndTwo(data))
}
