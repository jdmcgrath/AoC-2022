package main

import (
	"bufio"
	aoc "github.com/jdmcgrath/AoC-2022"
	"os"
	"strconv"
	"strings"
)

type elves struct {
	first  string
	second string
}
type elfAreas struct {
	lowerBound int
	upperBound int
}

func main() {
	readFile, err := os.Open("./day4/input.txt")
	aoc.Check(err)
	defer func() {
		aoc.Check(readFile.Close())
	}()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	totalEnclosed := 0
	for fileScanner.Scan() {
		splitString := strings.Split(fileScanner.Text(), ",")
		elves := elves{
			first:  splitString[0],
			second: splitString[1],
		}
		firstElfArea := getBounds(elves.first)
		secondElfArea := getBounds(elves.second)
		if isEnclosed := checkIfBoundsEntirelyEncapsulated(firstElfArea, secondElfArea); isEnclosed {
			totalEnclosed += 1
		}
	}
	println(totalEnclosed)
}

func getBounds(s string) elfAreas {
	boundsString := strings.Split(s, "-")
	lowerBound, err := strconv.Atoi(boundsString[0])
	aoc.Check(err)
	upperBound, err := strconv.Atoi(boundsString[1])
	aoc.Check(err)

	return elfAreas{
		lowerBound,
		upperBound,
	}
}

func checkIfBoundsEntirelyEncapsulated(firstElf, secondElf elfAreas) bool {
	return (firstElf.lowerBound <= secondElf.lowerBound && firstElf.upperBound >= secondElf.upperBound) || (secondElf.lowerBound <= firstElf.lowerBound && secondElf.upperBound >= firstElf.upperBound)
}
