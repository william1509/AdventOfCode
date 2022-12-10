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

	result := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ranges := strings.Split(scanner.Text(), ",")

		range1 := strings.Split(ranges[0], "-")
		range2 := strings.Split(ranges[1], "-")

		lb1, _ := strconv.Atoi(range1[0])
		ub1, _ := strconv.Atoi(range1[1])
		lb2, _ := strconv.Atoi(range2[0])
		ub2, _ := strconv.Atoi(range2[1])

		if (lb1 <= lb2 && ub1 >= ub2) || (lb2 <= lb1 && ub2 >= ub1) {
			result++
		}

	}
	return result
}

func part2() int {
	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ranges := strings.Split(scanner.Text(), ",")

		range1 := strings.Split(ranges[0], "-")
		range2 := strings.Split(ranges[1], "-")

		lb1, _ := strconv.Atoi(range1[0])
		ub1, _ := strconv.Atoi(range1[1])
		lb2, _ := strconv.Atoi(range2[0])
		ub2, _ := strconv.Atoi(range2[1])

		if (lb1 <= lb2 && ub1 >= lb2) || (lb2 <= lb1 && ub2 >= ub1) || (lb1 <= ub2 && ub1 >= ub2) {
			result++
		}

	}
	return result
}
