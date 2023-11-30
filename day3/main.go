package main

import (
	"bufio"
	"fmt"
	aoc "github.com/jdmcgrath/AoC-2022"
	"os"
)

func main() {
	readFile, err := os.Open("./day3/input.txt")
	aoc.Check(err)
	defer func() {
		aoc.Check(readFile.Close())
	}()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	totalPriorities := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		backpack := parseBackpack(line)
		commonItems := backpack.findCommonItems()
		for _, item := range commonItems {
			value := assignValueToRune(item)
			fmt.Printf("Found rune %s, worth %d\n", string(item), value)
			totalPriorities += value
		}
	}
	fmt.Println(totalPriorities)
}

type backpack struct {
	firstCompartment  []rune
	secondCompartment []rune
}

func parseBackpack(line string) backpack {
	runes := []rune(line)
	middle := len(runes) / 2
	return backpack{
		firstCompartment:  runes[:middle],
		secondCompartment: runes[middle:],
	}
}

func (bp *backpack) findCommonItems() []rune {
	set := make(map[rune]bool)
	commonSet := make(map[rune]bool)

	for _, item := range bp.firstCompartment {
		set[item] = true
	}

	var common []rune
	for _, item := range bp.secondCompartment {
		if _, found := set[item]; found {
			if _, isInCommon := commonSet[item]; !isInCommon {
				common = append(common, item)
				commonSet[item] = true
			}
		}
	}
	return common
}

func assignValueToRune(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 'a' + 1)
	} else if r >= 'A' && r <= 'Z' {
		return int(r - 'A' + 27)
	} else {
		return -1
	}
}
