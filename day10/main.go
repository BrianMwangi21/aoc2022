package main

import (
	"bufio"
	"fmt"
	"log"
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
	tmp := []rune(screen)
	tmp[pointToCheck] = '#'
	screen = string(tmp)
	return screen
}

func checkToDraw(screen []string, currentCycle int, X int) []string {
	if currentCycle <= 40 {
		pointToCheck := currentCycle - 1
		Xlower := X - 1
		XUpper := X + 1

		if pointToCheck >= Xlower && pointToCheck <= XUpper {
			screen[0] = changePointInScreen(screen[0], pointToCheck)
		}
	} else if currentCycle <= 80 {
		pointToCheck := currentCycle - 41
		Xlower := X - 1
		XUpper := X + 1

		if pointToCheck >= Xlower && pointToCheck <= XUpper {
			screen[1] = changePointInScreen(screen[1], pointToCheck)
		}
	} else if currentCycle <= 120 {
		pointToCheck := currentCycle - 81
		Xlower := X - 1
		XUpper := X + 1

		if pointToCheck >= Xlower && pointToCheck <= XUpper {
			screen[2] = changePointInScreen(screen[2], pointToCheck)
		}
	} else if currentCycle <= 160 {
		pointToCheck := currentCycle - 121
		Xlower := X - 1
		XUpper := X + 1

		if pointToCheck >= Xlower && pointToCheck <= XUpper {
			screen[3] = changePointInScreen(screen[3], pointToCheck)
		}
	} else if currentCycle <= 200 {
		pointToCheck := currentCycle - 161
		Xlower := X - 1
		XUpper := X + 1

		if pointToCheck >= Xlower && pointToCheck <= XUpper {
			screen[4] = changePointInScreen(screen[4], pointToCheck)
		}
	} else if currentCycle <= 240 {
		pointToCheck := currentCycle - 201
		Xlower := X - 1
		XUpper := X + 1

		if pointToCheck >= Xlower && pointToCheck <= XUpper {
			screen[5] = changePointInScreen(screen[5], pointToCheck)
		}
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

	fmt.Println("After", currentCycle, "cycles, here are the secret 8 letters")
	printScreen(screen)
}

func main() {
	data := readFile("input.txt")

	fmt.Println(partOne(data))
	partTwo(data)
}
