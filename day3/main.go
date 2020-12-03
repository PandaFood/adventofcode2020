package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file := "day3/input.csv"

	stepOutputs := []string{
		TraverseFile(file, 1, 1),
		TraverseFile(file, 3, 1),
		TraverseFile(file, 5, 1),
		TraverseFile(file, 7, 1),
		TraverseFile(file, 1, 2),
	}
	var trees [5]int

	for i, son := range stepOutputs {
		trees[i] = strings.Count(son, "#")
	}

	fmt.Printf("The amount of trees are: %#v \n", trees)
	fmt.Printf("The product of trees are: %d \n", trees[0]*trees[1]*trees[2]*trees[3]*trees[4])
}

func TraverseFile(filename string, step int, down int) string {
	globalIndex := step
	steppensSon := ""

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()

	for {
		success := false
		for i := 0; i < down; i++ {
			success = scanner.Scan()
		}
		if success == false {
			err = scanner.Err()
			if err == nil {
				break
			} else {
				log.Fatal(err)
				break
			}
		}
		text := scanner.Text()

		steppensSon += string(text[globalIndex%(len(text))])
		globalIndex += step
	}
	return steppensSon
}
