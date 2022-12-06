package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func isRepeated(data string) bool {
	for _, d := range data {
		if strings.Count(data, string(d)) > 1 {
			return true
		}
	}
	return false
}

func partOneOrTwo(data string, WINDOW_SIZE int) int {
	first_marker := 0

	for d := range data[:len(data)-WINDOW_SIZE-1] {
		window := ""

		for i := 0; i < WINDOW_SIZE; i++ {
			window += string(data[d+i])
		}

		if !isRepeated(window) {
			first_marker = d + WINDOW_SIZE
			break
		}
	}

	return first_marker
}

func main() {
	data := readFile("input.txt")

	fmt.Println(partOneOrTwo(data[0], 4))
	fmt.Println(partOneOrTwo(data[0], 14))
}
