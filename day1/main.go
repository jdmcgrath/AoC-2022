package main

import (
	"bufio"
	"fmt"
	"github.com/jdmcgrath/AoC-2022"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("./day1/input.txt")
	aoc.Check(err)
	defer func() {
		aoc.Check(readFile.Close())
	}()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	largestCount := 0
	currentPersonCount := 0
	for fileScanner.Scan() {
		calorieString := fileScanner.Text()
		if calorieString != "" {
			numberOfCalories, err := strconv.Atoi(calorieString)
			aoc.Check(err)
			currentPersonCount += numberOfCalories
		} else {
			if currentPersonCount > largestCount {
				largestCount = currentPersonCount
			}
			currentPersonCount = 0
		}
	}

	if currentPersonCount > largestCount {
		largestCount = currentPersonCount
	}

	fmt.Println(largestCount)
}
