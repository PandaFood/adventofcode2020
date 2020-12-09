package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file := "day9/input.csv"
	TraverseFile(file)
}

func TraverseFile(filename string) {
	preamble := 25
	numbers := make([]int, 0)

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
				break
			} else {
				log.Fatal(err)
				break
			}
		}
		row, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, row)
	}

	lastNumber, _ := FindXMAS(numbers, preamble)

	fmt.Printf("The number is: %d \n", lastNumber)

	contNumbers := FindContiguous(numbers, lastNumber)
	sumContNumbers := SumMaxMin(contNumbers)
	fmt.Printf("The second number is: %d from list %#v", sumContNumbers, contNumbers)

}

func FindContiguous(numbers []int, target int) []int {
	// fmt.Printf("%#v \n", numbers)
	for i, first := range numbers {
		for j, second := range numbers[i+1:] {
			// fmt.Printf("%d+%d - %#v \n", first, second, numbers[i:i+j+2])
			first += second
			if first == target {
				return numbers[i : i+j+2]
			} else if first > target {
				break
			}
		}
	}
	return []int{}
}

func SumMaxMin(list []int) int {
	min := list[0]
	max := list[0]
	for _, value := range list {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min + max
}

func FindXMAS(numbers []int, preamble int) (int, int) {
	for i, number := range numbers {
		if i <= preamble {
			continue
		}

		startIndex := i - preamble
		endIndex := i

		if !SumInPreamble(numbers[startIndex:endIndex], number) {
			return number, i
		}
	}

	return 0, 0
}

func SumInPreamble(preamble []int, number int) bool {
	for _, first := range preamble {
		for _, second := range preamble {
			if first+second == number {
				return true
			}
		}
	}
	return false
}
