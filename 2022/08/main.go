package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	var data []string

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	result := 0

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {

			// Tree is at the border
			if i == len(data[0])-1 || i == 0 || j == len(data)-1 || j == 0 {
				result++
				continue
			}

			biggest := 0
			value, _ := strconv.Atoi(string(data[i][j]))

			// Up
			for l := 0; l < i; l++ {
				curr, _ := strconv.Atoi(string(data[l][j]))
				if curr > biggest {
					biggest = curr
				}
			}

			if biggest < value {
				result++
				continue
			}

			// Down
			biggest = 0
			for l := i + 1; l < len(data[i]); l++ {
				curr, _ := strconv.Atoi(string(data[l][j]))
				if curr > biggest {
					biggest = curr
				}
			}

			if biggest < value {
				result++
				continue
			}

			// Left
			biggest = 0
			for l := 0; l < j; l++ {
				curr, _ := strconv.Atoi(string(data[i][l]))
				if curr > biggest {
					biggest = curr
				}
			}

			if biggest < value {
				result++
				continue
			}

			// Right
			biggest = 0
			for l := j + 1; l < len(data); l++ {
				curr, _ := strconv.Atoi(string(data[i][l]))
				if curr > biggest {
					biggest = curr
				}
			}

			if biggest < value {
				result++
				continue
			}
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

	scanner := bufio.NewScanner(file)

	var data []string

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	biggestScore := 0

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {

			// Tree is at the border
			if i == len(data[0])-1 || i == 0 || j == len(data)-1 || j == 0 {
				continue
			}

			value, _ := strconv.Atoi(string(data[i][j]))
			currentScore := 1
			viewingDistance := 1

			// Up
			for l := i - 1; l >= 0; l-- {
				curr, _ := strconv.Atoi(string(data[l][j]))
				if curr >= value {
					break
				}
				viewingDistance++
			}

			currentScore *= viewingDistance

			// Down
			viewingDistance = 0
			for l := i + 1; l < len(data[i]); l++ {
				curr, _ := strconv.Atoi(string(data[l][j]))
				if curr >= value {
					break
				}
				viewingDistance++
			}
			currentScore *= viewingDistance

			// Left
			viewingDistance = 0
			for l := j - 1; l >= 0; l-- {
				curr, _ := strconv.Atoi(string(data[i][l]))
				if curr >= value {
					break
				}
				viewingDistance++
			}
			currentScore *= viewingDistance

			// Right
			viewingDistance = 0
			for l := j + 1; l < len(data); l++ {
				curr, _ := strconv.Atoi(string(data[i][l]))
				if curr >= value {
					break
				}
				viewingDistance++
			}
			currentScore *= viewingDistance

			if currentScore > biggestScore {
				biggestScore = currentScore
			}
		}
	}

	return biggestScore
}

// func loopOver(data []string, s, e, j int) (biggest int) {
// 	for l := s + 1; l < e; l++ {
// 		curr, _ := strconv.Atoi(string(data[l][j]))
// 		if curr > biggest {
// 			biggest = curr
// 		}
// 	}
// 	return
// }
