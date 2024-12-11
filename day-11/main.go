package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const BLINKS = 25

func main() {

	content, err := os.ReadFile("input")

	if err != nil {
		log.Panic(err)
	}

	stones := make([]int, 0, 1000)

	for _, val := range strings.Split(strings.TrimSpace(string(content)), " ") {
		intVal, intError := strconv.Atoi(val)

		if intError != nil {
			log.Panic(intError)
		}
		stones = append(stones, intVal)
	}

	fmt.Println(stones)

	for i := BLINKS - 1; i >= 0; i-- {
		temp := make([]int, 0, 1000)

		for _, val := range stones {
			// fmt.Println(val)
			if val == 0 {
				temp = append(temp, 1)
			} else if hasEvenDigits(val) {
				strVal := strconv.Itoa(val)
				strLen := len(strVal) / 2

				first := strVal[:strLen]
				second := strVal[strLen:]

				num1, err1 := strconv.Atoi(first)
				num2, err2 := strconv.Atoi(second)

				if err1 != nil || err2 != nil {
					log.Panic(err1, err2)
				}

				temp = append(temp, num1, num2)
			} else {
				temp = append(temp, val*2024)
			}
		}

		stones = temp

		fmt.Println(i)
	}

	fmt.Println(len(stones))

}

func hasEvenDigits(input int) bool {
	return len(strconv.Itoa(input))%2 == 0
}
