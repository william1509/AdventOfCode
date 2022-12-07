// This sucks

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const FILENAME = "input"

func main() {
	fmt.Println("Part 1 : ", part1())
	fmt.Println("Part 2 : ", part2())
}

func part1() string {
	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	crates := make([]Stack[rune], 9)

	for scanner.Scan() {

		text := scanner.Text()

		// We reached the empty line after the crates diagram
		if strings.Contains(text, "1") {
			break
		}

		for i, r := range text {
			if r != ' ' && r != '[' && r != ']' {
				crates[i/4].Push(r)
			}
		}

	}
	// Skip empty line
	scanner.Scan()

	for _, item := range crates {
		item = reverse(item)
	}

	for scanner.Scan() {

		text := scanner.Text()
		var n, col1, col2 int

		fmt.Sscanf(text, "move %d from %d to %d", &n, &col1, &col2)
		for i := 0; i < n; i++ {
			crates[col2-1].Push(crates[col1-1].Pop())

		}

	}

	result := ""

	for _, item := range crates {
		result += string(item.Peek())
	}
	return result
}

func part2() string {
	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	crates := make([]Stack[rune], 9)

	for scanner.Scan() {

		text := scanner.Text()

		// We reached the empty line after the crates diagram
		if strings.Contains(text, "1") {
			break
		}

		for i, r := range text {
			if r != ' ' && r != '[' && r != ']' {
				crates[i/4].Push(r)
			}
		}

	}
	// Skip empty line
	scanner.Scan()

	for _, item := range crates {
		item = reverse(item)
	}

	for scanner.Scan() {

		text := scanner.Text()
		var n, col1, col2 int

		fmt.Sscanf(text, "move %d from %d to %d", &n, &col1, &col2)
		crates[col2-1].PushN(crates[col1-1].PopN(n))

	}

	result := ""

	for _, item := range crates {
		result += string(item.Peek())
	}
	return result
}

func reverse(a []rune) (c []rune) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	c = a
	return
}

// https://github.com/Urbansson/advent-of-code/blob/master/2022/day05/main.go

type Stack[C any] []C

func (s *Stack[C]) Push(v C) {
	*s = append(*s, v)
}

func (s *Stack[C]) PushN(v []C) {
	*s = append(*s, v...)
}

func (s *Stack[C]) Pop() C {
	if len(*s) == 0 {
		panic("stack is empty")
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack[C]) PopN(n int) []C {
	if len(*s) < n {
		panic("stack is empty")
	}
	v := (*s)[len(*s)-n:]
	*s = (*s)[:len(*s)-n]
	return v
}

func (s *Stack[C]) Peek() C {
	return (*s)[len(*s)-1]
}
