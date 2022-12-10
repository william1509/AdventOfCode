package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

const FILENAME = "input.dat"

func main() {
	fmt.Println("Part 1 : ", part1())
	fmt.Println("Part 2 : ", part2())
}

func part1() int {
	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	visited := make(map[Pair]bool, 0)
	head := Pair{0, 0}
	tail := Pair{0, 0}

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		var direction Pair
		switch tokens[0] {
		case "R":
			direction = Pair{1, 0}
		case "L":
			direction = Pair{-1, 0}
		case "D":
			direction = Pair{0, -1}
		case "U":
			direction = Pair{0, 1}
		}
		length, _ := strconv.Atoi(tokens[1])
		for i := 0; i < length; i++ {
			oldHead := head
			head = Add(head, direction)

			distance := Distance(head, tail)

			// First iteration
			switch distance {
			case 2:
				tail = Add(tail, direction)
			case 3:
				tail = oldHead
			}
			visited[tail] = true
		}
	}

	return len(visited)
}

func part2() int {
	//Read input file
	input, _ := os.Open(FILENAME)
	scanner := bufio.NewScanner(input)
	defer input.Close()

	tails := make([]Pair, 10)
	visited := make(map[Pair]bool)

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		var direction Pair
		switch tokens[0] {
		case "R":
			direction = Pair{1, 0}
		case "L":
			direction = Pair{-1, 0}
		case "D":
			direction = Pair{0, -1}
		case "U":
			direction = Pair{0, 1}
		}
		length, _ := strconv.Atoi(tokens[1])
		for i := 0; i < length; i++ {
			tails[0] = Add(tails[0], direction)
			for i := 0; i < len(tails)-1; i++ {
				tails[i+1] = setTail(tails[i+1], tails[i])
			}
			visited[tails[9]] = true
		}
	}

	return len(visited)
}

func setTail(tail Pair, head Pair) (newTail Pair) {
	newTail = tail
	offset := Pair{head.x - tail.x, head.y - tail.y}
	a := []Pair{{-2, 1}, {-1, 2}, {0, 2}, {1, 2}, {2, 1}, {2, 2}, {-2, 2}}
	b := []Pair{{1, 2}, {2, 1}, {2, 0}, {2, -1}, {1, -2}, {2, 2}, {2, -2}}
	c := []Pair{{-2, -2}, {2, -1}, {1, -2}, {0, -2}, {-1, -2}, {-2, -1}, {2, -2}}
	d := []Pair{{-2, -2}, {-1, -2}, {-2, -1}, {-2, -0}, {-2, 1}, {-1, 2}, {-2, 2}}

	if slices.Contains(a, offset) {
		newTail.y++
	}
	if slices.Contains(b, offset) {
		newTail.x++
	}
	if slices.Contains(c, offset) {
		newTail.y--
	}
	if slices.Contains(d, offset) {
		newTail.x--
	}

	return
}

type Pair struct {
	x, y int
}

func Add(a, b Pair) Pair {
	return Pair{a.x + b.x, a.y + b.y}
}

func Distance(a, b Pair) int {
	x := Abs(a.x - b.x)
	y := Abs(a.y - b.y)

	if x == 1 && y == 1 {
		return 1
	}

	return x + y

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func printBoard(h, t Pair) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 6; j++ {
			if h.x == j && h.y == i {
				fmt.Print("H")
			} else if t.x == j && t.y == i {
				fmt.Print("T")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func printBoard2(a []Pair) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 6; j++ {
			if slices.Contains(a, Pair{j, i}) {
				fmt.Print("X")

			} else {
				fmt.Print(".")

			}
		}
		fmt.Println()
	}
	fmt.Println()
}
