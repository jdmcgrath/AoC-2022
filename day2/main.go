package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	totalPoints, err := processFile("./day2/input.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(totalPoints)
}

func processFile(filename string) (int, error) {
	readFile, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("failed to open file: %w", err)
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)

	totalPoints := 0
	for scanner.Scan() {
		points, err := matchHandler(scanner.Text())
		if err != nil {
			return 0, err
		}
		totalPoints += points
	}
	return totalPoints, nil
}

func matchHandler(matchup string) (int, error) {
	hands := strings.Split(matchup, " ")
	if len(hands) != 2 || len(hands[0]) != 1 || len(hands[1]) != 1 {
		return 0, errors.New("invalid input")
	}

	opponentHand, myHand := rune(hands[0][0]), rune(hands[1][0])
	outcomePoints := getOutcome(opponentHand, myHand)
	shapePoints := getShapePoints(myHand)
	return outcomePoints + shapePoints, nil
}

type pair struct {
	opponentHand, myHand rune
}

func getOutcome(opponentHand, myHand rune) int {
	outcomes := map[pair]int{
		{'A', 'X'}: 3,
		{'A', 'Y'}: 6,
		{'A', 'Z'}: 0,
		{'B', 'X'}: 0,
		{'B', 'Y'}: 3,
		{'B', 'Z'}: 6,
		{'C', 'X'}: 6,
		{'C', 'Y'}: 0,
		{'C', 'Z'}: 3,
	}
	return outcomes[pair{opponentHand, myHand}]
}

func getShapePoints(myHand rune) int {
	shapes := map[rune]int{
		'X': 1,
		'Y': 2,
		'Z': 3,
	}
	return shapes[myHand]
}
