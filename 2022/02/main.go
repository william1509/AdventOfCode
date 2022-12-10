package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const FILENAME = "input.dat"

// Could've used modulos and shit but i'm too lazy
var PART_1 = map[Pair]int{
	{"A", "X"}: 4,
	{"A", "Y"}: 8,
	{"A", "Z"}: 3,

	{"B", "X"}: 1,
	{"B", "Y"}: 5,
	{"B", "Z"}: 9,

	{"C", "X"}: 7,
	{"C", "Y"}: 2,
	{"C", "Z"}: 6,
}

var PART_2 = map[Pair]int{
	{"A", "X"}: 3,
	{"A", "Y"}: 4,
	{"A", "Z"}: 8,

	{"B", "X"}: 1,
	{"B", "Y"}: 5,
	{"B", "Z"}: 9,

	{"C", "X"}: 2,
	{"C", "Y"}: 6,
	{"C", "Z"}: 7,
}

type Pair struct {
	a, b interface{}
}

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

	score := 0

	for scanner.Scan() {
		value := strings.Split(scanner.Text(), " ")
		score += PART_1[Pair{value[0], value[1]}]

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return score
}

func part2() int {
	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	score := 0

	for scanner.Scan() {
		value := strings.Split(scanner.Text(), " ")
		score += PART_2[Pair{value[0], value[1]}]

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return score
}
