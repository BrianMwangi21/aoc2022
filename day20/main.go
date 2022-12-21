package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"log"
	"os"
	"strconv"
)

const decryptionKey = 811589153

type positionKey struct {
	index int
	value int
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

func parseData(data []string) []int {
	var dataAsInt []int

	for _, d := range data {
		num, _ := strconv.Atoi(d)
		dataAsInt = append(dataAsInt, num)
	}

	return dataAsInt
}

func constructList(numbers []int) (map[positionKey]*ring.Ring, positionKey) {
	positions := make(map[positionKey]*ring.Ring, len(numbers))
	r := ring.New(len(numbers))
	zeroKey := positionKey{value: 0}
	for idx, number := range numbers {
		if number == 0 {
			zeroKey.index = idx
		}
		positions[positionKey{idx, number}] = r
		r.Value = number
		r = r.Next()
	}
	return positions, zeroKey
}

func mix(numbers []int, n int) (coordinateSum int) {
	positions, zeroKey := constructList(numbers)
	length := len(numbers) - 1
	halflen := length >> 1

	for ; n > 0; n-- {
		for idx, number := range numbers {
			r := positions[positionKey{idx, number}].Prev()
			removed := r.Unlink(1)
			if (number > halflen) || (number < -halflen) {
				number %= length
				switch {
				case number > halflen:
					number -= length
				case number < -halflen:
					number += length
				}
			}
			r.Move(number).Link(removed)
		}
	}

	r := positions[zeroKey]
	for i := 1; i <= 3; i++ {
		r = r.Move(1000)
		coordinateSum += r.Value.(int)
	}

	return coordinateSum
}

func partOne(data []int) int {
	sum := mix(data, 1)
	return sum
}

func partTwo(data []int) int {
	for index := range data {
		data[index] *= decryptionKey
	}
	sum := mix(data, 10)
	return sum
}

func main() {
	data := readFile("input.txt")
	dataAsInt := parseData(data)

	fmt.Println(partOne(dataAsInt))
	fmt.Println(partTwo(dataAsInt))
}
