package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const SECONDS = 100

const HEIGHT = 103

const WIDTH = 101

const INPUT = "input"

func main() {
	f, err := os.Open(INPUT)

	if err != nil {
		log.Fatal(err)
	}

	quadrants := make([]int, 4)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		processLine(scanner.Text(), quadrants)
	}

	log.Println(quadrants)

	result := 1

	for _, val := range quadrants {
		result *= val
	}

	log.Println(result)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	defer f.Close()
}

func processLine(line string, quadrants []int) {
	val := strings.Split(line, " ")
	position := strings.Split(val[0], "=")

	velocity := strings.Split(val[1], "=")

	postitionNums := strings.Split(position[1], ",")

	velocityNums := strings.Split(velocity[1], ",")

	pX, errpX := strconv.Atoi(postitionNums[0])

	pY, errpY := strconv.Atoi(postitionNums[1])

	vX, errvX := strconv.Atoi(velocityNums[0])

	vY, errvY := strconv.Atoi(velocityNums[1])

	if errpX != nil {
		log.Fatal(errpX)
	}
	if errpY != nil {
		log.Fatal(errpY)
	}
	if errvX != nil {
		log.Fatal(errvX)
	}
	if errvY != nil {
		log.Fatal(errvY)
	}

	finalX := (pX + (vX * SECONDS)) % WIDTH

	finalY := (pY + (vY * SECONDS)) % HEIGHT

	if finalY < 0 {
		finalY = HEIGHT + finalY
	}

	if finalX < 0 {
		finalX = WIDTH + finalX
	}

	middlepointHeight := HEIGHT / 2

	middlepointWidht := WIDTH / 2

	if finalX == middlepointWidht || finalY == middlepointHeight {
		return
	}

	if finalX < middlepointWidht && finalY < middlepointHeight {
		quadrants[0]++
	}

	if finalX < middlepointWidht && finalY > middlepointHeight {
		quadrants[1]++
	}

	if finalX > middlepointWidht && finalY < middlepointHeight {
		quadrants[2]++
	}

	if finalX > middlepointWidht && finalY > middlepointHeight {
		quadrants[3]++
	}

}
