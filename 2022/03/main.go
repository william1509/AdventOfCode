package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	var total rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		firstHalf := line[:len(line)/2]
		secondHalf := line[len(line)/2:]

		total += points(intersec(firstHalf, secondHalf))

	}
	return int(total)
}

func part2() int {
	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var total rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line1 := scanner.Text()

		scanner.Scan()
		line2 := scanner.Text()

		current := intersec(line1, line2)

		scanner.Scan()
		line3 := scanner.Text()

		current = intersec(current, line3)
		total += points(current)
	}
	return int(total)
}

func intersec(a, b string) (c string) {
	m := make(map[rune]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if m[item] {
			m[item] = false
			c += string(item)
		}
	}

	return
}

func points(a string) (c rune) {
	for _, item := range a {
		if item >= 97 {
			c += item - 96
		} else {
			c += item - 38
		}
	}

	return
}
