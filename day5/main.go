package main

import (
	"bufio"
	"fmt"
	aoc "github.com/jdmcgrath/AoC-2022"
	"os"
	"sort"
	"strconv"
	"strings"
)

type StackMap struct {
	stacks map[int][]rune
}

func main() {
	readFile, err := os.Open("./day5/input.txt")
	aoc.Check(err)
	defer func() {
		aoc.Check(readFile.Close())
	}()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sMap := &StackMap{stacks: make(map[int][]rune)}

	// Manually populate the initial state of the stacks
	sMap.stacks[1] = []rune("DBJV")
	sMap.stacks[2] = []rune("PVBWRDF")
	sMap.stacks[3] = []rune("RGFLDCWQ")
	sMap.stacks[4] = []rune("WJPMLNDB")
	sMap.stacks[5] = []rune("HNBPCSQ")
	sMap.stacks[6] = []rune("RDBSNG")
	sMap.stacks[7] = []rune("ZBPMQFSH")
	sMap.stacks[8] = []rune("WLF")
	sMap.stacks[9] = []rune("SVFMR")

	// Parse and execute the moves
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.HasPrefix(line, "move") {
			components := strings.Split(line, " ")
			numberToMove, _ := strconv.Atoi(components[1])
			fromStack, _ := strconv.Atoi(components[3])
			toStack, _ := strconv.Atoi(components[5])
			sMap.doMove(numberToMove, fromStack, toStack)
		}
	}

	// Print the contents of the stacks after all moves have been executed
	sMap.printStacks()
}

func (sm *StackMap) doMove(numberToMove, fromStack, toStack int) {
	if numberToMove > len(sm.stacks[fromStack]) {
		fmt.Println("Attempt to move more elements than available in the stack")
		return
	}
	movingRunes := sm.stacks[fromStack][len(sm.stacks[fromStack])-numberToMove:]
	sm.stacks[fromStack] = sm.stacks[fromStack][:len(sm.stacks[fromStack])-numberToMove]
	sm.stacks[toStack] = appendRuneSliceInReverse(sm.stacks[toStack], movingRunes)
}

func appendRuneSliceInReverse(slice, appendSlice []rune) []rune {
	for i := len(appendSlice) - 1; i >= 0; i-- {
		slice = append(slice, appendSlice[i])
	}

	return slice
}

func (sm *StackMap) printStacks() {
	keys := make([]int, 0, len(sm.stacks))
	for k := range sm.stacks {
		keys = append(keys, k)
	}
	sort.Ints(keys) // Ensure the keys are processed in order
	for _, key := range keys {
		stack := sm.stacks[key]
		if len(stack) > 0 { // Stack is not empty
			top := stack[len(stack)-1]
			fmt.Printf("%s", string(top))
		}
	}
}
