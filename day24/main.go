package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	SolvePart1()
	fmt.Println("====")
	SolvePart2()
}

type blizzard struct {
	x   int
	y   int
	dir byte
}

type valley struct {
	bl              []*blizzard
	numStepsToLoc   [][]int
	currentFrame    [][]byte
	nextFrame       [][]byte
	currentFrameNum int
	W               int
	H               int
	startX          int
	startY          int
}

func dx(dir byte) int {
	if dir == '>' {
		return 1
	}
	if dir == '<' {
		return -1
	}
	return 0
}

func dy(dir byte) int {
	if dir == '^' {
		return -1
	}
	if dir == 'v' {
		return 1
	}
	return 0
}

func addmod(a int, b int, mod int) int {
	t := (a + b) % mod
	if t < 0 {
		return t + mod
	}
	return t
}

func ParseValley(s []string) *valley {
	var v valley
	v.H = len(s) - 2
	v.W = len(s[0]) - 2
	v.bl = make([]*blizzard, 0)
	v.numStepsToLoc = make([][]int, v.H)
	v.currentFrame = make([][]byte, v.H)
	v.nextFrame = make([][]byte, v.H)
	v.startX = 0
	v.startY = 0

	for y := 0; y < v.H; y++ {
		v.numStepsToLoc[y] = make([]int, v.W)
		v.currentFrame[y] = make([]byte, v.W)
		v.nextFrame[y] = make([]byte, v.W)
	}

	for y := 1; y < len(s)-1; y++ {
		for x := 1; x < len(s[y])-1; x++ {
			if s[y][x] != '.' {
				var b blizzard
				b.x = x - 1
				b.y = y - 1
				b.dir = s[y][x]
				v.bl = append(v.bl, &b)
				v.currentFrame[y-1][x-1] = s[y][x]
			} else {
				v.currentFrame[y-1][x-1] = '.'
			}
		}
	}

	v.PopulateNextFrame()
	for y := 0; y < v.H; y++ {
		for x := 0; x < v.W; x++ {
			v.numStepsToLoc[y][x] = -1
		}
	}
	v.numStepsToLoc[0][0] = 0
	return &v
}

func (v *valley) PopulateNextFrame() {
	for y := 0; y < v.H; y++ {
		for x := 0; x < v.W; x++ {
			v.nextFrame[y][x] = '.'
		}
	}

	for _, b := range v.bl {
		x := addmod(b.x, dx(b.dir), v.W)
		y := addmod(b.y, dy(b.dir), v.H)
		if v.nextFrame[y][x] == '.' {
			v.nextFrame[y][x] = b.dir
		} else if v.nextFrame[y][x] == '>' || v.nextFrame[y][x] == '<' || v.nextFrame[y][x] == '^' || v.nextFrame[y][x] == 'v' {
			v.nextFrame[y][x] = '2'
		} else if v.nextFrame[y][x] == '2' {
			v.nextFrame[y][x] = '3'
		} else if v.nextFrame[y][x] == '3' {
			v.nextFrame[y][x] = '4'
		} else {
			v.nextFrame[y][x] = '?'
		}
	}
}

func (v *valley) Step() {
	v.StepExpeditionDists()

	// Move the nextFrame to currentFrame
	for y := 0; y < v.H; y++ {
		copy(v.currentFrame[y], v.nextFrame[y])
	}

	// Step all the blizzards
	for _, b := range v.bl {
		b.x = addmod(b.x, dx(b.dir), v.W)
		b.y = addmod(b.y, dy(b.dir), v.H)
	}

	v.PopulateNextFrame()
	v.currentFrameNum++
}

func (v *valley) isInBounds(x, y int) bool {
	return x >= 0 && x < v.W && y >= 0 && y < v.H
}

func (v *valley) StepExpeditionDists() {
	nextDists := make([][]int, v.H)
	for y := 0; y < v.H; y++ {
		nextDists[y] = make([]int, v.W)
		for x := 0; x < v.W; x++ {
			nextDists[y][x] = -1
		}
	}

	// it is always possible to enter the arena
	if v.nextFrame[v.startY][v.startX] == '.' {
		nextDists[v.startY][v.startX] = v.currentFrameNum + 1
	}

	for y := 0; y < v.H; y++ {
		for x := 0; x < v.W; x++ {
			if v.numStepsToLoc[y][x] != -1 {
				// try going up
				ny := y - 1
				nx := x
				if v.isInBounds(nx, ny) && v.nextFrame[ny][nx] == '.' {
					if nextDists[ny][nx] == -1 || nextDists[ny][nx] > (v.numStepsToLoc[y][x]+1) {
						nextDists[ny][nx] = v.numStepsToLoc[y][x] + 1
					}
				}

				// try going right
				ny = y
				nx = x + 1
				if v.isInBounds(nx, ny) && v.nextFrame[ny][nx] == '.' {
					if nextDists[ny][nx] == -1 || nextDists[ny][nx] > (v.numStepsToLoc[y][x]+1) {
						nextDists[ny][nx] = v.numStepsToLoc[y][x] + 1
					}
				}

				// try going down
				ny = y + 1
				nx = x
				if v.isInBounds(nx, ny) && v.nextFrame[ny][nx] == '.' {
					if nextDists[ny][nx] == -1 || nextDists[ny][nx] > (v.numStepsToLoc[y][x]+1) {
						nextDists[ny][nx] = v.numStepsToLoc[y][x] + 1
					}
				}

				// try going left
				ny = y
				nx = x - 1
				if v.isInBounds(nx, ny) && v.nextFrame[ny][nx] == '.' {
					if nextDists[ny][nx] == -1 || nextDists[ny][nx] > (v.numStepsToLoc[y][x]+1) {
						nextDists[ny][nx] = v.numStepsToLoc[y][x] + 1
					}
				}

				// try staying where you are
				ny = y
				nx = x
				if v.isInBounds(nx, ny) && v.nextFrame[ny][nx] == '.' {
					if nextDists[ny][nx] == -1 || nextDists[ny][nx] > (v.numStepsToLoc[y][x]+1) {
						nextDists[ny][nx] = v.numStepsToLoc[y][x] + 1
					}
				}
			}
		}
	}

	// copy it to dists
	for y := 0; y < v.H; y++ {
		copy(v.numStepsToLoc[y], nextDists[y])
	}
}

func (v *valley) checkWin() bool {
	return v.numStepsToLoc[v.H-1][v.W-1] != -1
}

func fill(a [][]int, val int) {
	for y := 0; y < len(a); y++ {
		for x := 0; x < len(a[0]); x++ {
			a[y][x] = val
		}
	}
}

func (v *valley) String() string {
	var b strings.Builder
	for _, s := range v.currentFrame {
		b.WriteString(string(s))
		b.WriteByte('\n')
	}
	return b.String()
}

func (v *valley) DistString() string {
	var b strings.Builder
	for y := 0; y < v.H; y++ {
		for x := 0; x < v.W; x++ {
			if v.numStepsToLoc[y][x] != -1 {
				b.WriteString(fmt.Sprint(v.numStepsToLoc[y][x]))
			} else {
				b.WriteByte(' ')
			}
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func SolvePart1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([]string, 0)
	for scanner.Scan() {
		s := scanner.Text()
		input = append(input, s)
	}
	fmt.Printf("read %v lines\n", len(input))

	v := ParseValley(input)

	for !v.checkWin() {
		//fmt.Printf("== Time %v == \n", (v.currentFrameNum + 1))
		v.Step()
		//fmt.Println(v)
		//fmt.Println("\ndists:")
		//fmt.Println(v.DistString())
	}
	fmt.Println("-----")
	fmt.Println("final answer", (v.numStepsToLoc[v.H-1][v.W-1] + 1))
}

func SolvePart2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([]string, 0)
	for scanner.Scan() {
		s := scanner.Text()
		input = append(input, s)
	}
	fmt.Printf("read %v lines\n", len(input))

	v := ParseValley(input)

	for !v.checkWin() {
		v.Step()
	}
	there := (v.numStepsToLoc[v.H-1][v.W-1] + 1)
	fmt.Println("steps there: ", there)
	// one step out
	v.Step()

	// try to go back
	fill(v.numStepsToLoc, -1)
	v.currentFrameNum = 0
	v.startX = v.W - 1
	v.startY = v.H - 1
	for v.numStepsToLoc[0][0] == -1 {
		v.Step()
		//fmt.Printf("\n== Step %v ==\n%v\n", v.currentFrameNum, v.DistString())
	}
	back := v.numStepsToLoc[0][0] + 1
	// one more step to go out (move the blizzards)
	v.Step()
	fmt.Println("steps back: ", back)

	// and back again
	fill(v.numStepsToLoc, -1)
	v.currentFrameNum = 0
	v.startX = 0
	v.startY = 0
	for v.numStepsToLoc[v.H-1][v.W-1] == -1 {
		v.Step()
		//fmt.Printf("\n== Step %v ==\n%v\n", v.currentFrameNum, v.DistString())
	}
	thereAgain := v.numStepsToLoc[v.H-1][v.W-1] + 1
	// one more step to go out (move the blizzards)
	v.Step()
	fmt.Println("steps there again: ", thereAgain)

	fmt.Println("total steps: ", (there + back + thereAgain))
}
