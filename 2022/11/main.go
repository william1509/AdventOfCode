package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	monkeys := make([]Monkey, 0)
	for scanner.Scan() {
		// id, _ := strconv.Atoi(strings.Trim(strings.Split(scanner.Text(), " ")[1], ":"))

		scanner.Scan()
		itemsStr := strings.Split(strings.Split(scanner.Text(), ":")[1], ",")
		itemsInt := make([]int, 0)

		for _, i := range itemsStr {
			j, err := strconv.Atoi(strings.Trim(i, " "))
			if err != nil {
				panic(err)
			}
			itemsInt = append(itemsInt, j)
		}

		scanner.Scan()
		tokens := strings.Split(scanner.Text(), " ")
		operation := tokens[len(tokens)-2]
		constant := tokens[len(tokens)-1]

		scanner.Scan()
		tokens = strings.Split(scanner.Text(), " ")
		test, _ := strconv.Atoi(tokens[len(tokens)-1])

		scanner.Scan()
		tokens = strings.Split(scanner.Text(), " ")
		trueMonkey, _ := strconv.Atoi(tokens[len(tokens)-1])

		scanner.Scan()
		tokens = strings.Split(scanner.Text(), " ")
		falseMonkey, _ := strconv.Atoi(tokens[len(tokens)-1])

		monkeys = append(monkeys, Monkey{operation, constant, test, trueMonkey, falseMonkey, itemsInt})

		scanner.Scan()
	}

	inspections := make([]int, len(monkeys))
	for i := 0; i < 20; i++ {
		for j := 0; j < len(monkeys); j++ {
			for k := 0; k < len(monkeys[j].items); k++ {
				value, err := strconv.Atoi(monkeys[j].constant)

				if err != nil {
					monkeys[j].items[k] *= monkeys[j].items[k]
				} else if monkeys[j].operation == "*" {
					monkeys[j].items[k] *= value
				} else if monkeys[j].operation == "+" {
					monkeys[j].items[k] += value
				} else {
					panic(0)
				}

				monkeys[j].items[k] /= 3

				if monkeys[j].items[k]%monkeys[j].test == 0 {
					monkeys[monkeys[j].trueMonkey].items = append(monkeys[monkeys[j].trueMonkey].items, monkeys[j].items[k])
				} else {
					monkeys[monkeys[j].falseMonkey].items = append(monkeys[monkeys[j].falseMonkey].items, monkeys[j].items[k])
				}
				inspections[j]++
			}
			monkeys[j].items = make([]int, 0)
		}
	}
	sort.Ints(inspections)

	return inspections[len(inspections)-2] * inspections[len(inspections)-1]
}

func part2() int {
	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	monkeys := make([]Monkey, 0)
	for scanner.Scan() {

		scanner.Scan()
		itemsStr := strings.Split(strings.Split(scanner.Text(), ":")[1], ",")
		itemsInt := make([]int, 0)

		for _, i := range itemsStr {
			j, err := strconv.Atoi(strings.Trim(i, " "))
			if err != nil {
				panic(err)
			}
			itemsInt = append(itemsInt, j)
		}

		scanner.Scan()
		tokens := strings.Split(scanner.Text(), " ")
		operation := tokens[len(tokens)-2]
		constant := tokens[len(tokens)-1]

		scanner.Scan()
		tokens = strings.Split(scanner.Text(), " ")
		test, _ := strconv.Atoi(tokens[len(tokens)-1])

		scanner.Scan()
		tokens = strings.Split(scanner.Text(), " ")
		trueMonkey, _ := strconv.Atoi(tokens[len(tokens)-1])

		scanner.Scan()
		tokens = strings.Split(scanner.Text(), " ")
		falseMonkey, _ := strconv.Atoi(tokens[len(tokens)-1])

		monkeys = append(monkeys, Monkey{operation, constant, test, trueMonkey, falseMonkey, itemsInt})

		scanner.Scan()
	}

	inspections := make([]int, len(monkeys))

	mod := 1
	for i := range monkeys {
		mod *= monkeys[i].test
	}

	for i := 0; i < 10000; i++ {
		for j := 0; j < len(monkeys); j++ {
			for k := 0; k < len(monkeys[j].items); k++ {
				value, err := strconv.Atoi(monkeys[j].constant)

				if err != nil {
					monkeys[j].items[k] *= monkeys[j].items[k]
				} else if monkeys[j].operation == "*" {
					monkeys[j].items[k] *= value
				} else if monkeys[j].operation == "+" {
					monkeys[j].items[k] += value
				} else {
					panic(0)
				}

				// makes sure that all the numbers stay within this range, without breaking the rest of the logic
				monkeys[j].items[k] %= mod

				if monkeys[j].items[k]%monkeys[j].test == 0 {
					monkeys[monkeys[j].trueMonkey].items = append(monkeys[monkeys[j].trueMonkey].items, monkeys[j].items[k])
				} else {
					monkeys[monkeys[j].falseMonkey].items = append(monkeys[monkeys[j].falseMonkey].items, monkeys[j].items[k])
				}
				inspections[j]++
			}
			monkeys[j].items = make([]int, 0)
		}
	}
	sort.Ints(inspections)

	return inspections[len(inspections)-2] * inspections[len(inspections)-1]
}

type Monkey struct {
	operation   string
	constant    string
	test        int
	trueMonkey  int
	falseMonkey int
	items       []int
}
