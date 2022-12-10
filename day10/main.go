package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
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

func initializeScreen() []string {
	var screen []string
	for i := 0; i < 6; i++ {
		s := ""
		for j := 0; j < 40; j++ {
			s += "."
		}
		screen = append(screen, s)
	}
	return screen
}

func printScreen(screen []string) {
	for d := range screen {
		fmt.Println(screen[d])
	}
}

func changePointInScreen(screen string, pointToCheck int) string {
	if pointToCheck > 0 && pointToCheck <= len(screen) {
		tmp := []rune(screen)
		tmp[pointToCheck] = '#'
		screen = string(tmp)
	}
	return screen
}

func checkToDraw(screen []string, currentCycle int, X int) []string {
	rowToBe := int(math.Floor(float64(currentCycle) / 40))

	pointToCheck := currentCycle - (rowToBe * 40) - 1
	Xlower := X - 1
	XUpper := X + 1

	if pointToCheck >= Xlower && pointToCheck <= XUpper {
		screen[rowToBe] = changePointInScreen(screen[rowToBe], pointToCheck)
	}

	return screen
}

func partOne(data []string) int {
	currentCycle, X, total := 1, 1, 0
	cyclesToWatch := [6]int{20, 60, 100, 140, 180, 220}

	cycles := make(map[int]int)
	cycles[currentCycle] = X

	for _, d := range data {
		splits := strings.Split(d, " ")
		command := splits[0]

		if command == "noop" {
			currentCycle++
			cycles[currentCycle] = X
		} else if command == "addx" {
			value, _ := strconv.Atoi(splits[1])

			currentCycle++
			cycles[currentCycle] = X

			currentCycle++
			X += value
			cycles[currentCycle] = X
		}
	}

	for _, c := range cyclesToWatch {
		total += c * cycles[c]
	}

	return total
}

func partTwo(data []string) {
	currentCycle, X := 1, 1
	cycles := make(map[int]int)
	cycles[currentCycle] = X
	screen := initializeScreen()
	screen = checkToDraw(screen, currentCycle, X)

	for _, d := range data {
		splits := strings.Split(d, " ")
		command := splits[0]

		if command == "noop" {
			currentCycle++
			cycles[currentCycle] = X
			screen = checkToDraw(screen, currentCycle, X)
		} else if command == "addx" {
			value, _ := strconv.Atoi(splits[1])

			currentCycle++
			cycles[currentCycle] = X
			screen = checkToDraw(screen, currentCycle, X)

			currentCycle++
			X += value
			cycles[currentCycle] = X
			screen = checkToDraw(screen, currentCycle, X)
		}
	}

	printScreen(screen)
}

func main() {
	data := readFile("input.txt")

	fmt.Println(partOne(data))
	partTwo(data)
}
