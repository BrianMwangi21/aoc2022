package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	id            int
	startingItems []int
	operation     []string
	test          int
	ifTrue        int
	ifFalse       int
	inspectCount  int
}

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

func parseData(data []string) []Monkey {
	var (
		chunks  [][]string
		chunk   []string
		monkeys []Monkey
	)

	for _, d := range data {
		if len(d) > 0 {
			chunk = append(chunk, d)
		} else {
			chunks = append(chunks, chunk)
			chunk = nil
		}
	}

	chunks = append(chunks, chunk)

	for _, chunk := range chunks {
		var (
			id            int
			startingItems []int
			operation     []string
			test          int
			ifTrue        int
			ifFalse       int
		)

		for i, c := range chunk {
			switch i {
			case 0:
				idString := strings.Split(c, " ")
				idString = strings.Split(idString[1], "")
				_id, _ := strconv.Atoi(idString[0])
				id = _id
			case 1:
				startingItemsString := strings.ReplaceAll(c, ",", "")
				startingItemsSplit := strings.Split(startingItemsString, " ")
				for _, s := range startingItemsSplit[4:] {
					item, _ := strconv.Atoi(s)
					startingItems = append(startingItems, item)
				}
			case 2:
				operationString := strings.Split(c, " ")
				operation = operationString[6:]
			case 3:
				testString := strings.Split(c, " ")
				_test, _ := strconv.Atoi(testString[len(testString)-1])
				test = _test
			case 4:
				ifTrueString := strings.Split(c, " ")
				_ifTrue, _ := strconv.Atoi(ifTrueString[len(ifTrueString)-1])
				ifTrue = _ifTrue
			case 5:
				ifFalseString := strings.Split(c, " ")
				_ifFalse, _ := strconv.Atoi(ifFalseString[len(ifFalseString)-1])
				ifFalse = _ifFalse
			}
		}

		monkey := Monkey{id, startingItems, operation, test, ifTrue, ifFalse, 0}
		monkeys = append(monkeys, monkey)
	}

	return monkeys
}

func performOperation(oldValue int, operation []string) int {
	newValue, operationValue := 1, 0

	if operation[1] == "old" {
		operationValue = oldValue
	} else {
		operationValue, _ = strconv.Atoi(operation[1])
	}

	switch operation[0] {
	case "+":
		newValue = oldValue + operationValue
	case "*":
		newValue = oldValue * operationValue
	}

	return newValue
}

func getLevelOfMonkeyBusiness(monkeys []Monkey) int {
	monkeyBusiness := 0
	var inspections []int

	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.inspectCount)
	}
	inspections = append(inspections, 0)

	sort.Ints(inspections)
	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))
	monkeyBusiness = inspections[0] * inspections[1]

	return monkeyBusiness
}

func getSuperMod(monkeys []Monkey) int {
	// Still don't get this part but, thanks to Reddit
	result := 1
	for _, monkey := range monkeys {
		result *= monkey.test
	}
	return result
}

func partOneOrPartTwo(monkeys []Monkey, rounds int, isWorried bool) int {
	monkeyBusiness := 0
	superMod := getSuperMod(monkeys)

	for i := 1; i <= rounds; i++ {
		for index, monkey := range monkeys {
			for _, item := range monkey.startingItems {
				new_value := performOperation(item, monkey.operation)
				if isWorried {
					new_value /= 3
				} else {
					new_value %= superMod
				}
				if (new_value % monkey.test) == 0 {
					monkeys[monkey.ifTrue].startingItems = append(monkeys[monkey.ifTrue].startingItems, new_value)
				} else {
					monkeys[monkey.ifFalse].startingItems = append(monkeys[monkey.ifFalse].startingItems, new_value)
				}
				monkeys[index].inspectCount += 1
			}
			monkeys[index].startingItems = nil
		}
	}

	monkeyBusiness = getLevelOfMonkeyBusiness(monkeys)
	return monkeyBusiness
}

func main() {
	data := readFile("input.txt")

	monkeys := parseData(data)
	fmt.Println(partOneOrPartTwo(monkeys, 20, true))
	monkeys = parseData(data)
	fmt.Println(partOneOrPartTwo(monkeys, 10000, false))
}
