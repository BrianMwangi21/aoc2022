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

func shapeScore(shape string) int {
	switch shape {
	case "A", "X": // Rock
		return 1
	case "B", "Y": // Paper
		return 2
	case "C", "Z": // Scissors
		return 3
	}
	return 0
}

func roundScore(opp_score int, hero_score int) int {
	if opp_score == hero_score {
		return 3 // Same shape
	}
	if hero_score == 1 && opp_score == 3 {
		return 6 // Rock beats Scissors
	}
	if hero_score == 2 && opp_score == 1 {
		return 6 // Paper beats Rock
	}
	if hero_score == 3 && opp_score == 2 {
		return 6 // Scissors beats Paper
	}
	return 0
}

func decisionScore(shape string) int {
	switch shape {
	case "X": // Need to Lose
		return 0
	case "Y": // Need to draw
		return 3
	case "Z": // Need to win
		return 6
	}
	return 0
}

func changeMind(d_score int, opponent string, hero string) string {
	opp_score := shapeScore(opponent)
	if d_score == 3 {
		return opponent
	}
	if d_score == 0 {
		if opp_score == 1 {
			return "Z"
		}
		if opp_score == 2 {
			return "X"
		}
		if opp_score == 3 {
			return "Y"
		}
	}
	if d_score == 6 {
		if opp_score == 1 {
			return "Y"
		}
		if opp_score == 2 {
			return "Z"
		}
		if opp_score == 3 {
			return "X"
		}
	}
	return hero
}

func partOne(data []string) int {
	score := 0

	for d := range data {
		roundData := strings.Split(data[d], " ")

		s_score := shapeScore(roundData[1])
		r_score := roundScore(shapeScore(roundData[0]), shapeScore(roundData[1]))

		score += s_score + r_score
	}

	return score
}

func partTwo(data []string) int {
	score := 0

	for d := range data {
		roundData := strings.Split(data[d], " ")

		d_score := decisionScore(roundData[1])
		new_shape := changeMind(d_score, roundData[0], roundData[1])
		s_score := shapeScore(new_shape)
		r_score := roundScore(shapeScore(roundData[0]), shapeScore(new_shape))

		score += s_score + r_score
	}

	return score
}

func main() {
	data := readFile("input.txt")

	fmt.Println(partOne(data))
	fmt.Println(partTwo(data))
}
