package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// colors:= [3]string{"red", "green", "blue"}
var MaxCubes = map[string]int{
	"green": 13,
	"red":   12,
	"blue":  14,
}

func part1() {
	//Open input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//New scanner
	scanner := bufio.NewScanner(file)

	totalId := 0

	// Split line into game and trys
	for scanner.Scan() {
		line := scanner.Text()
		totalId += checkMaxCubes(line)
		fmt.Println(totalId)
	}

}

func checkMaxCubes(line string) int {
	re := regexp.MustCompile("[0-9]+")
	re2 := regexp.MustCompile("red|green|blue")
	lineParts := strings.Split(line, ":")
	gameId, err := strconv.Atoi(re.FindString(lineParts[0]))
	if err != nil {
		os.Exit(2)
	}
	gameTrys := strings.Split(lineParts[1], ";")
	for _, ln := range gameTrys {
		cubes := re.FindAllString(ln, -1)
		cubeColor := re2.FindAllString(ln, -1)
		fmt.Println(ln)
		fmt.Println(cubes)
		fmt.Println(cubeColor)
		for i := range cubes {
			cube, _ := strconv.Atoi(cubes[i])
			if cube > MaxCubes[cubeColor[i]] {
				fmt.Println("dejo de ser SME")
				return 0
			}
		}
	}
	return gameId
}

func part2() {
	totalPower := 0
	//Open input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//New scanner
	scanner := bufio.NewScanner(file)

	// Split line into game and trys
	for scanner.Scan() {
		line := scanner.Text()
		totalPower += checkMinimumCubes(line)
	}
	fmt.Println(totalPower)
}

func checkMinimumCubes(line string) int {
	maxGreen, maxRed, maxBlue := 0, 0, 0
	re := regexp.MustCompile("[0-9]+")
	re2 := regexp.MustCompile("red|green|blue")
	lineParts := strings.Split(line, ":")
	gameTrys := strings.Split(lineParts[1], ";")
	for _, ln := range gameTrys {
		cubes := re.FindAllString(ln, -1)
		cubeColor := re2.FindAllString(ln, -1)
		fmt.Println(ln)
		fmt.Println(cubes)
		fmt.Println(cubeColor)
		for i := range cubes {
			cube, _ := strconv.Atoi(cubes[i])
			if cubeColor[i] == "green" {
				if cube > maxGreen {
					maxGreen = cube
				}
			} else if cubeColor[i] == "red" {
				if cube > maxRed {
					maxRed = cube
				}
			} else if cubeColor[i] == "blue" {
				if cube > maxBlue {
					maxBlue = cube
				}
			}
		}
	}
	return maxBlue * maxGreen * maxRed
}

func main() {
	part1()
	part2()
}
