package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/go-deeper/chunks"
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

func getValueOfLetter(letter string) int {
	value, _ := utf8.DecodeRuneInString(letter)
	result := 0

	if strings.ToUpper(letter) == letter {
		result = int(value) - 38
		return result
	}

	if strings.ToLower(letter) == letter {
		result = int(value) - 96
		return result
	}

	return int(value)
}

func getCommonLetterInString(runsack string) string {
	first := runsack[:len(runsack)/2]
	second := runsack[len(runsack)/2:]

	for l := range first {
		if strings.Contains(second, string(first[l])) {
			return string(first[l])
		}
	}

	return ""
}

func getCommonLetterInArray(runsacks []string) string {
	first := runsacks[0]
	second := runsacks[1]
	third := runsacks[2]

	for l := range first {
		if strings.Contains(second, string(first[l])) && strings.Contains(third, string(first[l])) {
			return string(first[l])
		}
	}

	return ""
}

func partOne(data []string) int {
	total := 0

	for d := range data {
		commonLetter := getCommonLetterInString(data[d])
		total += getValueOfLetter(commonLetter)
	}

	return total
}

func partTwo(data []string) int {
	total := 0
	chunk := chunks.Split(data, 3)

	for c := range chunk {
		commonLetter := getCommonLetterInArray(chunk[c])
		total += getValueOfLetter(commonLetter)
	}

	return total
}

func main() {
	data := readFile("input.txt")

	fmt.Println(partOne(data))
	fmt.Println(partTwo(data))
}
