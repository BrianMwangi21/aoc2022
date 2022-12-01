package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

func partOne(data []string) int {
	var totals []int
	sum := 0

	for d := range data {
		if data[d] != "" {
			calories, _ := strconv.Atoi(data[d])
			sum += calories
		} else {
			totals = append(totals, sum)
			sum = 0
		}
	}

	totals = append(totals, sum)
	sort.Ints(totals)

	return totals[len(totals)-1]
}

func partTwo(data []string) int {
	var totals []int
	sum := 0

	for d := range data {
		if data[d] != "" {
			calories, _ := strconv.Atoi(data[d])
			sum += calories
		} else {
			totals = append(totals, sum)
			sum = 0
		}
	}

	totals = append(totals, sum)
	sort.Ints(totals)

	return totals[len(totals)-1] + totals[len(totals)-2] + totals[len(totals)-3]
}

func main() {
	data := readFile("input.txt")

	fmt.Println(partOne(data))
	fmt.Println(partTwo(data))
}
