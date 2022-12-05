package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
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

func getCranesAndProcedures(data []string) (map[int][]string, map[int][]int) {
	var (
		tmp_crates     []string
		tmp_procedures []string
	)

	for i, d := range data {
		if len(d) == 0 {
			tmp_crates = data[:i-1]     // Don't really need the keys
			tmp_procedures = data[i+1:] // Don't really need the blank space
		}
	}

	// Crates
	crates := map[int][]string{}

	for _, d := range tmp_crates {
		for i, l := range d {
			if unicode.IsLetter(l) {
				crates[i] = append(crates[i], string(l))
			}
		}
	}

	for i, c := range crates {
		crates[i] = reverseCrate(c)
	}

	// Procedures
	procedures := map[int][]int{}

	for i, p := range tmp_procedures {
		split_p := strings.Split(p, " ")
		num, _ := strconv.Atoi(split_p[1])
		src, _ := strconv.Atoi(split_p[3])
		dest, _ := strconv.Atoi(split_p[5])
		procedures[i] = append(procedures[i], num, src, dest)
	}

	return crates, procedures
}

func reverseCrate(data []string) []string {
	var reversed []string
	for i := len(data) - 1; i >= 0; i-- {
		reversed = append(reversed, data[i])
	}
	return reversed
}

func orderCrates(data map[int][]string) ([]int, map[int][]string) {
	var keys []int
	for key := range data {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	ordered := map[int][]string{}
	for i, key := range keys {
		ordered[i+1] = data[key]
	}
	return keys, ordered
}

func orderProcedures(data map[int][]int) ([]int, map[int][]int) {
	var keys []int
	for key := range data {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	ordered := map[int][]int{}
	for i, key := range keys {
		ordered[i] = data[key]
	}
	return keys, ordered
}

func partOne(data []string) string {
	var result string

	allCrates, allProcedures := getCranesAndProcedures(data)

	_, orderedCrates := orderCrates(allCrates)
	_, orderedProcedures := orderProcedures(allProcedures)

	fmt.Println(orderedCrates)
	fmt.Println(orderedProcedures)

	return result
}

func main() {
	data := readFile("input.test.txt")

	// Part One : SHMSDGZVC
	// Part Two : VRZGHDFBQ
	fmt.Println(partOne(data))
}
