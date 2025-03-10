package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const DIAGONAL_DISTANCE = 14

const HORIZONTAL_DISTANCE = 10

const INPUT = "input.test"

type Node struct {
	Gcost    int
	Hcost    int
	Fcont    int
	explored bool
}

type Coord struct {
	X int
	Y int
}

func main() {
	f, err := os.Open(INPUT)

	if err != nil {
		log.Fatal(err)
	}

	maze := make([][]string, 0, 1000)

	start := &Coord{}

	end := &Coord{}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		for i, val := range line {
			if val == "S" {
				start = &Coord{X: i, Y: len(maze)}
			}
			if val == "E" {
				end = &Coord{X: i, Y: len(maze)}
			}
		}
		maze = append(maze, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(maze)
	fmt.Println("start:", start)
	fmt.Println("end:", end)

}
