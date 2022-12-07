package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	scanner.Scan()

	text := scanner.Text()

	nDifferentLetters := 4

	for i, _ := range text {

		buff := make(map[byte]bool)
		for j := 0; j < nDifferentLetters; j++ {
			buff[text[i+j]] = true
		}

		if len(buff) == nDifferentLetters {
			return i + nDifferentLetters
		}
	}

	return 0
}

func part2() int {
	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	text := scanner.Text()

	nDifferentLetters := 14

	for i, _ := range text {

		buff := make(map[byte]bool)
		for j := 0; j < nDifferentLetters; j++ {
			buff[text[i+j]] = true
		}

		if len(buff) == nDifferentLetters {
			return i + nDifferentLetters
		}
	}

	return 0
}
