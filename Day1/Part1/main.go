package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

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

func main() {
	part1()
}
