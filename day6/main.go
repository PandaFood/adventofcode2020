package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	file := "day6/input.csv"
	TraverseFile(file)
}

func TraverseFile(filename string) {
	finishedReading := false
	text := ""
	totalVotes := 0

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
				finishedReading = true
			} else {
				log.Fatal(err)
				break
			}
		}
		line := scanner.Text()
		//fmt.Printf("%v scanned \n", line)

		if len(strings.TrimSpace(line)) != 0 {
			text += line + "\n"
		} else {
			// Part 1
			//totalVotes += CountVotes(text)

			// Part 2
			totalVotes += AddVotes(text)
			text = ""
		}

		if finishedReading {
			break
		}
	}
	fmt.Printf("The total votes are: %d", totalVotes)
}

// PART 1

func CountVotes(text string) int {
	letters := "abcdefghijklmnopqrstuvwxyz"
	count := 0

	text = strings.ReplaceAll(text, "\n", "")

	for _, letter := range letters {
		count += IncrementOnMatch(text, string(letter))
	}

	return count
}

func IncrementOnMatch(text string, regEx string) int {
	regEx = "(" + regEx + ")"
	matched, _ := regexp.Match(regEx, []byte(text))

	if matched {
		return 1
	} else {
		return 0
	}
}

// PART 2

func AddVotes(text string) int {
	letters := "abcdefghijklmnopqrstuvwxyz"
	count := 0
	text = strings.TrimSuffix(text, "\n")
	votes := strings.Split(text, "\n")
	answers := make([]string, len(votes))

	for i, vote := range votes {
		for _, letter := range letters {
			answers[i] += ReturnOnMatch(vote, string(letter))
		}
	}

	for _, letter := range answers[0] {
		if IsLetterInAllElements(answers, letter) {
			count++
		}
	}

	return count
}

func ReturnOnMatch(text string, letter string) string {
	regEx := "(" + letter + ")"
	matched, _ := regexp.Match(regEx, []byte(text))

	if matched {
		return letter
	} else {
		return ""
	}
}

func IsLetterInAllElements(list []string, letter rune) bool {
	for _, item := range list {
		if !Contains(item, letter) {
			return false
		}
	}
	return true
}

func Contains(list string, item rune) bool {
	for _, li := range list {
		if li == item {
			return true
		}
	}
	return false
}
