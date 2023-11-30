package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	readFile, err := os.Open("./day1input.txt")
	check(err)
	defer func() {
		check(readFile.Close())
	}()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	largestCount := 0
	currentPersonCount := 0
	for fileScanner.Scan() {
		calorieString := fileScanner.Text()
		if calorieString != "" {
			numberOfCalories, err := strconv.Atoi(calorieString)
			check(err)
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
