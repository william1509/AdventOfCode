package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const FILENAME = "input"

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
			case 0, 1:
				break
			case 2:
				tail = Add(tail, direction)
				break
			case 3:
				tail = oldHead
				break
			}
			visited[tail] = true
		}
	}

	return len(visited)
}

func part2() int {
	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	visited := make(map[Pair]bool, 0)
	head := Pair{0, 0}
	tail := make([]Pair, 9)

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

			for i, item := range tail {
				distance := Distance(head, tail)

				// First iteration
				switch distance {
				case 0, 1:
					break
				case 2:
					tail = Add(tail, direction)
					break
				case 3:
					tail = oldHead
					break
				}
			}

			distance := Distance(head, tail)

			// First iteration
			switch distance {
			case 0, 1:
				break
			case 2:
				tail = Add(tail, direction)
				break
			case 3:
				tail = oldHead
				break

			}
			visited[tail] = true
		}
	}

	return len(visited)
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
