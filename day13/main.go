package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	nodeType string
	rep      string
	value    int
	children []*Node
	special  int
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

func parseData(data []string) [][]string {
	var (
		chunks [][]string
		chunk  []string
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
	return chunks
}

func cmp(a, b int) int {
	if a < b {
		return 1
	}
	if a > b {
		return -1
	}
	return 0
}

func compareNodes(n1 *Node, n2 *Node) int {
	if n1.nodeType == "number" && n2.nodeType == "number" {
		return cmp(n1.value, n2.value)
	}
	if n1.nodeType == "number" {
		n1.nodeType = "list"
		n1.children = []*Node{{nodeType: "number", value: n1.value}}
	}
	if n2.nodeType == "number" {
		n2.nodeType = "list"
		n2.children = []*Node{{nodeType: "number", value: n2.value}}
	}
	smaller := len(n1.children)
	if len(n2.children) < smaller {
		smaller = len(n2.children)
	}

	for i := 0; i < smaller; i++ {
		check := compareNodes(n1.children[i], n2.children[i])
		if check == 0 {
			continue
		}
		return check
	}
	return cmp(len(n1.children), len(n2.children))
}

func parseList(input string) *Node {
	n := &Node{nodeType: "list", children: []*Node{}, rep: input}
	for i := 1; i < len(input); i++ {
		if input[i] == '[' {
			open := 1
			for j := i + 1; j < len(input); j++ {
				if input[j] == '[' {
					open++
				}
				if input[j] == ']' {
					open--
				}
				if open == 0 {
					n.children = append(n.children, parseList(input[i:j+1]))
					i = j
					break
				}
			}
		} else if input[i] == ',' {
			continue
		} else {
			for j := i + 1; j < len(input); j++ {
				if input[j] == ',' || input[j] == ']' {
					n.children = append(n.children, parseNumber(input[i:j]))
					i = j
					break
				}
			}
		}
	}
	return n
}

func parseNumber(input string) *Node {
	n := &Node{nodeType: "number", rep: input}
	n.value, _ = strconv.Atoi(input)
	return n
}

func partOne(data [][]string) int {
	sum := 0

	for i, d := range data {
		first := parseList(d[0])
		second := parseList(d[1])

		if check := compareNodes(first, second); check == 1 {
			sum += i + 1
			continue
		}
	}

	return sum
}

func partTwo(data []string) int {
	var (
		key   int
		nodes []*Node
	)
	for _, d := range data {
		if d == "" {
			continue
		}
		nodes = append(nodes, parseList(strings.TrimSpace(d)))
	}

	nodes = append(nodes, &Node{nodeType: "list", children: []*Node{
		{nodeType: "list", children: []*Node{
			{nodeType: "number", value: 2},
		}},
	}, special: 2})
	nodes = append(nodes, &Node{nodeType: "list", children: []*Node{
		{nodeType: "list", children: []*Node{
			{nodeType: "number", value: 6},
		}},
	}, special: 6})

	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			check := compareNodes(nodes[i], nodes[j])
			if check == -1 {
				nodes[i], nodes[j] = nodes[j], nodes[i]
			}
		}
	}

	sixIndex := 0
	twoIndex := 0
	for i, n := range nodes {
		if n.special == 6 {
			sixIndex = i + 1
		}
		if n.special == 2 {
			twoIndex = i + 1
		}
	}

	key = sixIndex * twoIndex
	return key
}

func main() {
	data := readFile("input.txt")

	inputs := parseData(data)
	fmt.Println(partOne(inputs))
	fmt.Println(partTwo(data))
}
