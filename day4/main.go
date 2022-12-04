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

func isPairInclusive(pairs []string) bool {
	limitsLeft := strings.Split(pairs[0], "-")
	limitsRight := strings.Split(pairs[1], "-")

	limitsLeft_lower, _ := strconv.Atoi(limitsLeft[0])
	limitsLeft_upper, _ := strconv.Atoi(limitsLeft[1])
	limitsRight_lower, _ := strconv.Atoi(limitsRight[0])
	limitsRight_upper, _ := strconv.Atoi(limitsRight[1])

	if limitsLeft_lower <= limitsRight_lower && limitsLeft_upper >= limitsRight_upper {
		return true
	}

	if limitsRight_lower <= limitsLeft_lower && limitsRight_upper >= limitsLeft_upper {
		return true
	}

	return false
}

func isPairShared(pairs []string) bool {
	limitsLeft := strings.Split(pairs[0], "-")
	limitsRight := strings.Split(pairs[1], "-")

	limitsLeft_lower, _ := strconv.Atoi(limitsLeft[0])
	limitsLeft_upper, _ := strconv.Atoi(limitsLeft[1])
	limitsRight_lower, _ := strconv.Atoi(limitsRight[0])
	limitsRight_upper, _ := strconv.Atoi(limitsRight[1])

	if limitsLeft_lower >= limitsRight_lower && limitsLeft_lower <= limitsRight_upper {
		return true
	}

	if limitsLeft_upper >= limitsRight_lower && limitsLeft_lower <= limitsRight_upper {
		return true
	}

	if limitsRight_lower >= limitsLeft_lower && limitsRight_lower <= limitsLeft_upper {
		return true
	}

	if limitsRight_upper >= limitsLeft_lower && limitsRight_upper <= limitsLeft_upper {
		return true
	}

	return false
}

func partOne(data []string) int {
	total := 0

	for d := range data {
		pairs := strings.Split(data[d], ",")
		isInclusive := isPairInclusive(pairs)

		if isInclusive {
			total += 1
		}
	}

	return total
}

func partTwo(data []string) int {
	total := 0

	for d := range data {
		pairs := strings.Split(data[d], ",")
		isShared := isPairShared(pairs)

		if isShared {
			total += 1
		}
	}

	return total
}

func main() {
	data := readFile("input.txt")

	fmt.Println(partOne(data))
	fmt.Println(partTwo(data))
}
