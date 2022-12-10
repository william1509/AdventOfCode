package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	X := 1
	history := make([]int, 0)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		history = append(history, X)

		if tokens[0] == "addx" {
			value, _ := strconv.Atoi(tokens[1])
			X += value
			history = append(history, X)

		}

	}

	result := 0

	cycles := []int{20, 60, 100, 140, 180, 220}

	// No idea why -2 instead of -1
	for _, item := range cycles {
		fmt.Println(history[item-2])
		result += history[item-2] * item
	}
	return result
}

// Not perfect but good enough
func part2() int {
	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	X := 1
	history := make([]int, 0)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		history = append(history, X)

		if tokens[0] == "addx" {
			value, _ := strconv.Atoi(tokens[1])
			X += value
			history = append(history, X)

		}
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			pos := i*40 + j

			if history[pos]-2 == j || history[pos]-1 == j || history[pos] == j {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	return 0
}
