package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file := "day7/input.csv"
	TraverseFile(file)
}

func TraverseFile(filename string) {
	bags := make([][]string, 0)

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
		bagLine := strings.Split(scanner.Text(), "contain")
		bags = append(bags, bagLine)
	}

	contained := SearchThroughBags(bags)
	fmt.Printf("The briefcase can be contained in %d other bags \n", len(contained))

}

func SearchThroughBags(bags [][]string) []string {
	baseContained := ListContained(bags, "shiny gold")
	allContained := TraverseBagTree(bags, baseContained)
	allContained = PruneListToUniques(allContained)
	return allContained
}

func TraverseBagTree(allBags [][]string, baseBags []string) []string {
	containedBags := make([]string, 0)
	loops := 0

	for _, bag := range baseBags {
		containedBags = append(containedBags, PruneBagName(bag))
	}
	//fmt.Printf("%#v \n", containedBags)

	for {
		bagsSize := len(containedBags)
		//fmt.Printf("%#v \n", containedBags)
		for _, bag := range containedBags {
			newBags := ListContained(allBags, bag)
			//fmt.Printf("%#v \n", bag)
			for _, newBag := range newBags {
				newBag = PruneBagName(newBag)
				containedBags = append(containedBags, newBag)
			}
		}
		// fmt.Printf("%#v \n", containedBags)
		containedBags = PruneListToUniques(containedBags)

		if bagsSize == len(containedBags) || loops > 10 {
			break
		}
		loops++
	}

	return containedBags
}

func ListContained(bags [][]string, contained string) []string {
	bagsContaining := make([]string, 0)

	for _, bag := range bags {
		if strings.Contains(bag[1], contained) {
			bagsContaining = append(bagsContaining, bag[0])
		}
	}
	return bagsContaining
}

func FindContained(bags [][]string, contained string) int {
	for i, bag := range bags {
		if strings.Contains(bag[1], contained) {
			return i
		}
	}
	return -1
}

func PruneBagName(bag string) string {
	bag = strings.Replace(bag, "bags", "", -1)
	bag = strings.Replace(bag, "bag", "", -1)
	bag = strings.TrimSpace(bag)
	return bag
}

func PruneListToUniques(list []string) []string {
	returnList := make([]string, 0)
	for _, item := range list {
		if !Contains(returnList, item) {
			returnList = append(returnList, item)
		}
	}
	return returnList
}

func Contains(list []string, searched string) bool {
	for _, word := range list {
		if word == searched {
			return true
		}
	}
	return false
}
