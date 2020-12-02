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
	count := 0

	file, err := os.Open("day2/input.csv")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for {
		success := scanner.Scan()
		if success == false {
			// False on error or EOF. Check error
			err = scanner.Err()
			if err == nil {
				log.Println("Scan completed and reached EOF")
				break
			} else {
				log.Fatal(err)
				break
			}
		}
		str := scanner.Text()
		if IsValid(str) {
			count++
		}
	}
	fmt.Printf("Valid passwords are: %d", count)
}

func IsValid(pw string) bool {
	str := strings.Split(pw, " ")
	nums := strings.Split(str[0], "-")
	min, _ := strconv.Atoi(nums[0])
	max, _ := strconv.Atoi(nums[1])
	letter := string(str[1][0])
	word := str[2]

	if IsWithinBounds(min, max, letter, word) {
		return true
	} else {
		return false
	}
}

func IsWithinRange(min int, max int, letter string, pw string) bool {
	count := strings.Count(pw, letter)

	if min <= count && count <= max {
		return true
	} else {
		return false
	}
}

func IsWithinBounds(min int, max int, letter string, pw string) bool {
	pos1Valid := strings.Contains(string(pw[min-1]), letter)
	pos2Valid := strings.Contains(string(pw[max-1]), letter)
	if (pos1Valid || pos2Valid) && !(pos1Valid && pos2Valid) {
		return true
	} else {
		return false
	}
}
