package main

import (
	"bufio"
	"log"
	"os"
)

type Land struct {
	Area      int
	Perimeter int
}

func main() {
	f, err := os.Open("input.test")

	defer f.Close()

	scanner := bufio.NewScanner(f)

	if err != nil {
		log.Fatal(err)
	}

	matrix := make([][]byte, 0, 1000)

	for scanner.Scan() {
		// do something with a line
		matrix = append(matrix, scanner.Bytes())
	}

	results := map[byte]Land{}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i, slice := range matrix {
		log.Println(slice)
		for j, elem := range slice {
			walls := 0

			//check if the up element is a wall
			upslice, up_ok := get(matrix, i-1)
			if !up_ok {
				walls++
			} else {
				up_elem, up_elem_ok := get(upslice, j)
				if !up_elem_ok {
					log.Panic(up_elem)
				}
				if up_elem != elem {
					walls++
				}
			}

			//get the down element
			downslice, down_ok := get(matrix, i+1)
			if !down_ok {
				walls++
			} else {
				down_elem, down_elem_ok := get(downslice, j)
				if !down_elem_ok {
					log.Panic(down_elem_ok)
				}
				if down_elem != elem {
					walls++
				}
			}

			//check if the left element is a wall
			left_elem, left_ok := get(slice, j-1)

			if !left_ok {
				walls++
			} else {
				if left_elem != elem {
					walls++
				}
			}

			//get the right element

			right_elem, right_ok := get(slice, j+1)

			if !right_ok {
				walls++
			} else {
				if right_elem != elem {
					walls++
				}
			}

			results[elem] = Land{Area: results[elem].Area + 1, Perimeter: results[elem].Perimeter + walls}
		}
	}

	log.Println(results)
}

func get[T any](slice []T, index int) (T, bool) {

	var defa T

	if index < 0 || index >= len(slice) {
		return defa, false
	}

	return slice[index], true
}
