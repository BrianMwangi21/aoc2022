package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

type blueprint struct {
	id         int
	oreRobot   oreRobot
	clayRobot  clayRobot
	obsRobot   obsRobot
	geodeRobot geodeRobot
}

type oreRobot struct {
	oreCost int
}

type clayRobot struct {
	oreCost int
}

type obsRobot struct {
	oreCost, clayCost int
}

type geodeRobot struct {
	oreCost, obsCost int
}

var globalBest = 0

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

func parseInput(r io.Reader) []blueprint {
	scanner := bufio.NewScanner(r)
	blueprints := []blueprint{}

	for scanner.Scan() {
		var id int
		oreRobot := oreRobot{}
		clayRobot := clayRobot{}
		obsRobot := obsRobot{}
		geodeRobot := geodeRobot{}

		fmt.Sscanf(scanner.Text(),
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&id, &oreRobot.oreCost, &clayRobot.oreCost, &obsRobot.oreCost, &obsRobot.clayCost, &geodeRobot.oreCost, &geodeRobot.obsCost)

		bp := blueprint{
			id,
			oreRobot,
			clayRobot,
			obsRobot,
			geodeRobot,
		}

		blueprints = append(blueprints, bp)
	}

	return blueprints
}

func search(bp blueprint, ore, clay, obs, time, oreRobots, clayRobots, obsRobots, geodeRobots, geodes int) int {
	if time == 0 || globalBest >= geodes+rangeSum(geodeRobots, geodeRobots+time-1) {
		return 0
	}
	if oreRobots >= bp.geodeRobot.oreCost && obsRobots >= bp.geodeRobot.obsCost {
		return rangeSum(geodeRobots, geodeRobots+time-1)
	}

	oreLimitHit := oreRobots >= int(math.Max(float64(bp.geodeRobot.oreCost), math.Max(float64(bp.clayRobot.oreCost), float64(bp.obsRobot.oreCost))))
	clayLimitHit := clayRobots >= bp.obsRobot.clayCost
	obsLimitHit := obsRobots >= bp.geodeRobot.obsCost
	best := 0

	if !oreLimitHit {
		best = int(math.Max(
			float64(best),
			float64(geodeRobots+search(
				bp, ore+oreRobots, clay+clayRobots, obs+obsRobots,
				time-1, oreRobots, clayRobots, obsRobots, geodeRobots, geodes+geodeRobots))))
	}
	if ore >= bp.oreRobot.oreCost && !oreLimitHit {
		best = int(math.Max(
			float64(best),
			float64(geodeRobots+search(
				bp, ore-bp.oreRobot.oreCost+oreRobots, clay+clayRobots, obs+obsRobots,
				time-1, oreRobots+1, clayRobots, obsRobots, geodeRobots, geodes+geodeRobots))))
	}
	if ore >= bp.clayRobot.oreCost && !clayLimitHit {
		best = int(math.Max(
			float64(best),
			float64(geodeRobots+search(
				bp, ore-bp.clayRobot.oreCost+oreRobots, clay+clayRobots, obs+obsRobots,
				time-1, oreRobots, clayRobots+1, obsRobots, geodeRobots, geodes+geodeRobots))))
	}
	if ore >= bp.obsRobot.oreCost && clay >= bp.obsRobot.clayCost && !obsLimitHit {
		best = int(math.Max(
			float64(best),
			float64(geodeRobots+search(
				bp, ore-bp.obsRobot.oreCost+oreRobots, clay-bp.obsRobot.clayCost+clayRobots, obs+obsRobots,
				time-1, oreRobots, clayRobots, obsRobots+1, geodeRobots, geodes+geodeRobots))))
	}
	if ore >= bp.geodeRobot.oreCost && obs >= bp.geodeRobot.obsCost {
		best = int(math.Max(
			float64(best),
			float64(geodeRobots+search(
				bp, ore-bp.geodeRobot.oreCost+oreRobots, clay+clayRobots, obs-bp.geodeRobot.obsCost+obsRobots,
				time-1, oreRobots, clayRobots, obsRobots, geodeRobots+1, geodes+geodeRobots))))
	}

	globalBest = int(math.Max(float64(best), float64(globalBest)))
	return best
}

func rangeSum(first, last int) int {
	return last*(last+1)/2 - ((first - 1) * first / 2)
}

func partOne(blueprints []blueprint) int {
	result := 0
	for _, bp := range blueprints {

		result += bp.id * search(bp, 0, 0, 0, 24, 1, 0, 0, 0, 0)
		globalBest = 0
	}

	return result
}

func partTwo(blueprints []blueprint) int {
	if len(blueprints) < 3 {
		return -1
	}
	result := 1
	for i := 0; i < 3; i++ {
		result *= search(blueprints[i], 0, 0, 0, 32, 1, 0, 0, 0, 0)
		globalBest = 0
	}

	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	blueprints := parseInput(file)
	fmt.Println("Part 1:", partOne(blueprints))
	fmt.Println("Part 2:", partTwo(blueprints))
}
