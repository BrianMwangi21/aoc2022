package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func splitCratesAndProcedures(data []string) ([]string, []string) {
	var (
		crates     []string
		procedures []string
	)

	for d := range data {
		if len(data[d]) == 0 {
			crates = data[:d]
			procedures = data[d+1:]
		}
	}

	return crates, procedures
}

func getIndividualCrates(data []string) []string {
	crateHolder := make([]string, 100)

	for _, d := range data[:len(data)-1] {
		for i, l := range d {
			if unicode.IsLetter(l) {
				crateHolder[i] += string(l)
			}
		}
	}

	var actualCrateHolder []string
	for _, c := range crateHolder {
		if len(c) != 0 {
			actualCrateHolder = append(actualCrateHolder, reverseCrate(c))
		}
	}

	return actualCrateHolder
}

func reverseCrate(crate string) string {
	var reversed string

	for i := len(crate) - 1; i >= 0; i-- {
		reversed += string(crate[i])
	}

	return reversed
}

func getIndividualProcedure(data string) (int, int, int) {
	moves := strings.Split(data, " ")
	num, _ := strconv.Atoi(moves[1])
	src, _ := strconv.Atoi(moves[3])
	dest, _ := strconv.Atoi(moves[5])

	return num, src, dest
}

func printCrate(data []string) {
	for _, d := range data {
		fmt.Println(d)
	}
	fmt.Println("====")
}

func partOne(data []string) string {
	crates, procedures := splitCratesAndProcedures(data)
	var topCrates string

	crateHolder := getIndividualCrates(crates)

	for _, p := range procedures {
		num, src, dest := getIndividualProcedure(p)
		var tmp string

		src_crate := crateHolder[src-1]
		dest_crate := crateHolder[dest-1]

		if len(src_crate)-num >= 0 {
			tmp = src_crate[len(src_crate)-num:]
		}

		src_crate = strings.ReplaceAll(src_crate, tmp, "")
		dest_crate += reverseCrate(tmp)

		crateHolder[src-1] = src_crate
		crateHolder[dest-1] = dest_crate

		// printCrate(crateHolder)
	}

	for _, c := range crateHolder {
		if len(c) > 0 {
			topCrates += string(c[len(c)-1])
		}
	}

	return topCrates
}

func partTwo(data []string) string {
	crates, procedures := splitCratesAndProcedures(data)
	var topCrates string

	crateHolder := getIndividualCrates(crates)

	for _, p := range procedures {
		num, src, dest := getIndividualProcedure(p)
		var tmp string

		src_crate := crateHolder[src-1]
		dest_crate := crateHolder[dest-1]

		if len(src_crate)-num >= 0 {
			tmp = src_crate[len(src_crate)-num:]
		}

		src_crate = strings.ReplaceAll(src_crate, tmp, "")
		dest_crate += tmp

		crateHolder[src-1] = src_crate
		crateHolder[dest-1] = dest_crate

		// printCrate(crateHolder)
	}

	for _, c := range crateHolder {
		if len(c) > 0 {
			topCrates += string(c[len(c)-1])
		}
	}

	return topCrates
}

func main() {
	data := readFile("input.txt")

	fmt.Println(partOne(data)) // SHMSDGZVC
	fmt.Println(partTwo(data)) // VRZGHDFBQ
}
