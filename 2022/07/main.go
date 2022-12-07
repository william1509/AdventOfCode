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
	var currentNode *Node = &Node{"/", nil, make([]*File, 0), make([]*Node, 0)}

	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		tokens := strings.Split(text, " ")
		if tokens[0] == "$" {
			switch tokens[1] {
			case "cd":
				if tokens[2] == ".." {
					currentNode = currentNode.parent
					continue
				}
				if tokens[2] == "/" {
					currentNode = getRoot(currentNode)
					continue
				}

				var alreadyDiscoveredNode *Node
				for _, item := range currentNode.dir {
					if item.name == tokens[2] {
						alreadyDiscoveredNode = item
						break
					}
				}

				if alreadyDiscoveredNode != nil {
					currentNode = alreadyDiscoveredNode
					break
				}

				currentNode = &Node{tokens[2], currentNode, make([]*File, 0), make([]*Node, 0)}
			case "ls":
				break
			default:
				break
			}

		} else {
			switch tokens[0] {
			case "dir":
				currentNode.dir = append(currentNode.dir, &Node{tokens[1], currentNode, make([]*File, 0), make([]*Node, 0)})
			default:
				value, _ := strconv.Atoi(tokens[0])
				currentNode.files = append(currentNode.files, &File{value, tokens[1]})
			}
		}
	}
	result := getSmallDirSize(getRoot(currentNode))

	return result
}

func part2() int {
	var currentNode *Node = &Node{"/", nil, make([]*File, 0), make([]*Node, 0)}

	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		tokens := strings.Split(text, " ")
		if tokens[0] == "$" {
			switch tokens[1] {
			case "cd":
				if tokens[2] == ".." {
					currentNode = currentNode.parent
					continue
				}
				if tokens[2] == "/" {
					currentNode = getRoot(currentNode)
					continue
				}

				var alreadyDiscoveredNode *Node
				for _, item := range currentNode.dir {
					if item.name == tokens[2] {
						alreadyDiscoveredNode = item
						break
					}
				}

				if alreadyDiscoveredNode != nil {
					currentNode = alreadyDiscoveredNode
					break
				}

				currentNode = &Node{tokens[2], currentNode, make([]*File, 0), make([]*Node, 0)}
			case "ls":
				break
			default:
				break
			}

		} else {
			switch tokens[0] {
			case "dir":
				currentNode.dir = append(currentNode.dir, &Node{tokens[1], currentNode, make([]*File, 0), make([]*Node, 0)})
			default:
				value, _ := strconv.Atoi(tokens[0])
				currentNode.files = append(currentNode.files, &File{value, tokens[1]})
			}
		}
	}

	missingSpace := 30000000 - (70000000 - getDirSize(getRoot(currentNode)))
	result := lookForSmallestDir(getRoot(currentNode), missingSpace, 9999999999)
	return result
}

func getDirSize(n *Node) int {
	fileSize := 0
	for _, item := range n.files {
		fileSize += item.size
	}

	dirSize := 0
	for _, item := range n.dir {
		dirSize += getDirSize(item)
	}

	return fileSize + dirSize
}

func getSmallDirSize(n *Node) (size int) {
	size = 0
	curr := getDirSize(n)

	if curr < 100000 {
		size += curr
	}
	for _, item := range n.dir {
		size += getSmallDirSize(item)
	}

	return
}

func lookForSmallestDir(n *Node, requiredSpace int, bestDir int) int {
	curr := getDirSize(n)

	if curr > requiredSpace && curr < bestDir {
		bestDir = curr
	}
	for _, item := range n.dir {
		bestDir = lookForSmallestDir(item, requiredSpace, bestDir)
	}

	return bestDir
}

func getRoot(c *Node) *Node {
	current := c
	for {
		if current.parent == nil {
			return current
		}
		current = current.parent
	}
}

type Node struct {
	name   string
	parent *Node
	files  []*File
	dir    []*Node
}

type File struct {
	size int
	name string
}
