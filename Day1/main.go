package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var matchmap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"0":     0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func part1() {
	//Open input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Creater a scanner for the file
	scanner := bufio.NewScanner(file)
	sum := 0
	regex := regexp.MustCompile(`\d`)

	// Loop through the lines in file

	for scanner.Scan() {
		line := scanner.Text()
		digits := regex.FindAllString(line, -1)
		// Check if there are numbers
		if len(digits) > 0 {
			// Get first and last digits and concatenate them
			firstDigit := digits[0]
			lastDigit := digits[len(digits)-1]
			twoDigitNumber, _ := strconv.Atoi(firstDigit + lastDigit)
			sum += twoDigitNumber
		}
	}

	// Check for scannig error
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// Print total sum
	fmt.Println("Part one:", sum)
}

func part2() {
	// Open input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a new scanner for the file
	scanner := bufio.NewScanner(file)
	var input []string

	// Loop over all lines in the file and add thenm to the input slice
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	// Check for scanning error
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	sum := 0

	// Loop over all lines in the input slice
	for _, ln := range input {
		first := -1
		last := -1
		minpos := len(ln) + 1
		maxpos := -1

		// Loop over all keys in the matchmap
		for k, v := range matchmap {
			p := strings.Index(ln, k)
			if p == -1 {
				//substring not found
				continue
			}
			// substring found
			if p < minpos {
				first = v
				minpos = p
			}
			// Check if the substring is the last one
			p = strings.LastIndex(ln, k)
			if p > maxpos {
				last = v
				maxpos = p
			}
		}
		sum += 10*first + last
	}
	fmt.Printf("Part two: %v\n", sum)
}

func main() {
	part1()
	part2()
}
