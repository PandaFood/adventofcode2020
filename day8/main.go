package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := "day8/input.csv"
	TraverseFile(file)
}

func TraverseFile(filename string) {
	instructions := make([]string, 0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for {
		success := scanner.Scan()
		if success == false {
			err = scanner.Err()
			if err == nil {
				text := scanner.Text()
				instructions = append(instructions, text)
				break
			} else {
				log.Fatal(err)
				break
			}
		}
		text := scanner.Text()
		instructions = append(instructions, text)
	}
	// fmt.Printf("instructions: %#v \n", instructions)

	TrySwap(instructions)
}

func TrySwap(instructions []string) {

	for i, op := range instructions {
		newSet := make([]string, len(instructions))
		copy(newSet, instructions)

		if strings.Contains(op, "jmp") {
			newSet[i] = strings.Replace(newSet[i], "jmp", "nop", 1)
		} else if strings.Contains(op, "nop") {
			newSet[i] = strings.Replace(newSet[i], "nop", "jmp", 1)
		} else if strings.Contains(op, "acc") {
			continue
		}

		// fmt.Printf("instructions: %#v \n", newSet)

		val, finished := Compute(newSet)

		if finished {
			fmt.Printf("Finished with val: %d \n", val)
			break
		}
	}

}

func Compute(instructions []string) (int, bool) {
	finishedInstructions := make([]int, 0)
	accumulator := 0
	index := 0
	finished := false

	for {
		if Contains(finishedInstructions, index) {
			break
		}
		if instructions[index] == "" {
			finished = true
			//fmt.Printf("Finished: %t with val: %d \n", finished, accumulator)
			return accumulator, finished
		}

		op := instructions[index][0:4]
		value, _ := strconv.Atoi(instructions[index][4:])
		finishedInstructions = append(finishedInstructions, index)
		// fmt.Printf("%d: OP: %s %d - %d \n", index, op, value, accumulator)

		if strings.Contains(op, "acc") {
			accumulator += value
			index++
		} else if strings.Contains(op, "jmp") {
			index += value
		} else if strings.Contains(op, "nop") {
			index++
		}
	}

	return accumulator, finished
}

func Contains(list []int, index int) bool {
	for _, item := range list {
		if item == index {
			return true
		}
	}
	return false
}
