package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()

	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	parsed := make([]int, 0, 1000)

	for scanner.Scan() {
		// do something with a line
		for _, val := range scanner.Bytes() {
			intVal := int16(val) - 48

			parsed = append(parsed, int(intVal))
		}
	}

	final := make([]int, 0, 2000)

	lastValIdx := len(parsed) - 1

	for i, val := range parsed {

		if i > lastValIdx {
			break
		}

		if i%2 == 0 {
			for val > 0 {
				final = append(final, i/2)
				val--
			}
		} else {

			for val > 0 {
				if parsed[lastValIdx] > 0 {
					//do stuff
					final = append(final, lastValIdx/2)

					parsed[lastValIdx]--

				} else {
					lastValIdx -= 2
					final = append(final, lastValIdx/2)

					parsed[lastValIdx]--
				}
				val--

			}
		}
	}

	accumulator := 0

	for idx, val := range final {
		accumulator += idx * val
	}

	log.Print(accumulator)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	log.Printf("The time it took %s", elapsed)

}

// func main() {
// 	f, err := os.Open("input.test")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// remember to close the file at the end of the program
// 	defer f.Close()

// 	scanner := bufio.NewScanner(f)

// 	parsed := make([]byte, 0, 1000)

// 	for scanner.Scan() {
// 		// do something with a line
// 		for i, val := range scanner.Bytes() {
// 			intVal := int16(val) - 48

// 			if i%2 == 0 {
// 				for range intVal {
// 					parsed = append(parsed, byte((i/2)+48))
// 				}
// 			} else {
// 				for range intVal {
// 					parsed = append(parsed, byte(46))
// 				}

// 			}
// 		}
// 	}

// 	lastIdx := len(parsed)

// 	for idx, val := range parsed {
// 		if val == 46 {

// 		}
// 	}

// 	for _, val := range parsed {
// 		log.Print("%v ", int(val))
// 	}
// }
