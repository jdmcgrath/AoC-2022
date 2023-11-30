package main

import (
	"reflect"
	"testing"
)

func Test_doMove(t *testing.T) {
	tests := []struct {
		name          string
		initialStacks map[int][]rune
		moveArgs      struct {
			numberToMove, fromStack, toStack int
		}
		expectedStacks map[int][]rune
	}{
		{
			name: "Test 1",
			initialStacks: map[int][]rune{
				1: {'a', 'b', 'c'},
				2: {'d', 'e', 'f'},
			},
			moveArgs: struct {
				numberToMove, fromStack, toStack int
			}{
				numberToMove: 1,
				fromStack:    1,
				toStack:      2,
			},
			expectedStacks: map[int][]rune{
				1: {'a', 'b'},
				2: {'d', 'e', 'f', 'c'},
			},
		},
		// add more test cases as needed.
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			stackMapInstance := &StackMap{
				stacks: test.initialStacks,
			}

			stackMapInstance.doMove(test.moveArgs.numberToMove, test.moveArgs.fromStack, test.moveArgs.toStack)

			if !reflect.DeepEqual(test.expectedStacks, stackMapInstance.stacks) {
				t.Errorf("doMove() = %v, want %v", stackMapInstance.stacks, test.expectedStacks)
			}
		})
	}
}
