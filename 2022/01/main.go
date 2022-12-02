package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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

	var calories = 0
	var maxCalories = 0

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		calories += value

		if err != nil {
			if calories > maxCalories {
				maxCalories = calories
			}
			calories = 0
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return maxCalories
}

func part2() int {
	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	calories := 0
	var data []int

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		calories += value

		if err != nil {
			data = append(data, calories)
			calories = 0
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(data)

	result := 0
	for _, v := range data[len(data)-3:] {
		result += v
	}

	return result
}
