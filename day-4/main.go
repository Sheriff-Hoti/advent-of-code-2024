package main

import (
	"fmt"
	"log"
	"os"
)

/*

88 is the ASCII value for X
77 is the ASCII value for M
65 is the ASCII value for A
83 is the ASCII value for S
10 is the ASCII value for the new line character
*/

func main() {

	byteSlice := readInput()

	length := findLength(byteSlice) + 1

	counter := 0

	for idx := range byteSlice {

		if downCheck(byteSlice, idx, length) {
			counter++
		}

		if downRightCheck(byteSlice, idx, length) {
			counter++
		}

		if rightCheck(byteSlice, idx) {
			counter++
		}

		if downLeftCheck(byteSlice, idx, length) {
			counter++
		}
	}

	fmt.Println("Final result:", counter)

}

func downCheck(slice []byte, index int, length int) bool {
	first, ok1 := get(slice, index)
	second, ok2 := get(slice, index+length)
	third, ok3 := get(slice, index+(2*length))
	fourth, ok4 := get(slice, index+(3*length))

	if !ok1 || !ok2 || !ok3 || !ok4 {
		return false
	}

	if first == byte(88) && second == byte(77) && third == byte(65) && fourth == byte(83) {
		return true
	}

	if first == byte(83) && second == byte(65) && third == byte(77) && fourth == byte(88) {
		return true
	}

	return false
}

func downRightCheck(slice []byte, index int, length int) bool {
	first, ok1 := get(slice, index)
	second, ok2 := get(slice, index+length+1)
	third, ok3 := get(slice, index+(2*length)+2)
	fourth, ok4 := get(slice, index+3+(3*length))

	if !ok1 || !ok2 || !ok3 || !ok4 {
		return false
	}

	if first == byte(88) && second == byte(77) && third == byte(65) && fourth == byte(83) {
		return true
	}

	if first == byte(83) && second == byte(65) && third == byte(77) && fourth == byte(88) {
		return true
	}

	return false
}

func downLeftCheck(slice []byte, index int, length int) bool {
	first, ok1 := get(slice, index)
	second, ok2 := get(slice, index+length-1)
	third, ok3 := get(slice, index+(2*length)-2)
	fourth, ok4 := get(slice, index-3+(3*length))

	if !ok1 || !ok2 || !ok3 || !ok4 {
		return false
	}

	if first == byte(88) && second == byte(77) && third == byte(65) && fourth == byte(83) {
		return true
	}

	if first == byte(83) && second == byte(65) && third == byte(77) && fourth == byte(88) {
		return true
	}

	return false
}

func rightCheck(slice []byte, index int) bool {
	first, ok1 := get(slice, index)
	second, ok2 := get(slice, index+1)
	third, ok3 := get(slice, index+2)
	fourth, ok4 := get(slice, index+3)

	if !ok1 || !ok2 || !ok3 || !ok4 {
		return false
	}

	if first == byte(88) && second == byte(77) && third == byte(65) && fourth == byte(83) {
		return true
	}

	if first == byte(83) && second == byte(65) && third == byte(77) && fourth == byte(88) {
		return true
	}

	return false
}

func get(slice []byte, index int) (byte, bool) {
	if index < 0 || index >= len(slice) {
		return byte(0), false // Return zero value and a boolean indicating the index is invalid
	}
	if slice[index] == byte(10) {
		return byte(0), false
	}
	return slice[index], true // Return the value and true if index is valid
}

func findLength(input []byte) int {
	counter := 0

	for _, val := range input {
		if val == 10 {
			break
		}
		counter++
	}

	return counter
}

func readInput() []byte {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	return content
}
