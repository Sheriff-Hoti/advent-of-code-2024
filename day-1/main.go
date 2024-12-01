package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/**
steps:

1. Read file line by line
2. Process each line (processLine function) (based on the format of the saved input this function might change)
3. Convert the string line to number pairs and append each number to their respective slice
4. Sort the slices
5. Compute the absolute difference between the respective slice elements with the same index eg. abs(slice[1]-slice[2])
6. Save the results into a new slice s3
7. Sum the elements of the slice s3. That's the final result

*/

func readFileByLine(processLine func(string) (int, int)) {

	arr1 := make([]int, 1000)
	arr2 := make([]int, 1000)
	counter := 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num1, num2 := processLine(scanner.Text())
		arr1[counter] = num1
		arr2[counter] = num2
		counter++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})

	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})

	arr3 := make([]int, 1000)

	if len(arr1) != len(arr2) {
		log.Panic("the length of two slices is not equal, slice 1:", len(arr1), " ,slice 2:", len(arr2))
	}

	other_counter := len(arr1)

	for i := range other_counter {
		res := arr1[i] - arr2[i]

		if res < 0 {
			arr3[i] = res * (-1)
		} else {
			arr3[i] = res
		}

	}

	result := 0
	for _, val := range arr3 {
		result += val
	}

	log.Printf("Final result: %v", result)

}

func processLine(input string) (int, int) {
	result := strings.Fields(input)

	if len(result) != 2 {
		log.Panic("bad input line, row: ", result)
	}

	num1, err1 := strconv.Atoi(result[0])
	num2, err2 := strconv.Atoi(result[1])

	if err1 != nil || err2 != nil {
		log.Panic(err1, err2)
	}

	return num1, num2

}

func main() {
	start := time.Now()
	readFileByLine(processLine)
	elapsed := time.Since(start)
	log.Printf("The time it took %s", elapsed)
}
