package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type bagType struct {
	name    string
	subBags []bagType
	count   int
}

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

	name := "shiny gold"
	contained := GetContainingBags(bags, name)
	fmt.Printf("The briefcase can be contained in %d other bags \n", len(contained))

	containing := GetContainedBags(bags, name)
	amountContaining := CountRecursiveBags(containing.subBags) - 1
	fmt.Printf("The briefcase contains %d other bags \n", amountContaining)

}

// Part 2
// -------------

func GetContainedBags(bags [][]string, bag string) bagType {
	allContaining := TraverseBagTreeDown(bags, bag)
	return allContaining
}

func ListContaining(bags [][]string, contains string) []string {
	for _, bag := range bags {
		if strings.Contains(bag[0], contains) {
			return ParseContainingName(bag[1])
		}
	}
	return ParseContainingName("")
}

func TraverseBagTreeDown(allBags [][]string, bag string) bagType {
	baseBag := bagType{
		bag,
		make([]bagType, 0),
		0,
	}

	baseBag, _ = RecursiveAddBags(allBags, baseBag)

	fmt.Printf("Bag: %v \n", baseBag)
	return baseBag
}

func CountRecursiveBags(bags []bagType) int {
	count := 1

	for _, bag := range bags {
		count += bag.count * CountRecursiveBags(bag.subBags)
		// fmt.Printf("Bag %s with %d at count %d \n", bag.name, bag.count, count)
	}
	return count
}

func ParseContainingName(bags string) []string {
	names := make([]string, 0)

	for _, name := range strings.Split(bags, ",") {
		name = strings.TrimSpace(name)
		name = strings.TrimSuffix(name, ".")
		name = PruneBagName(name)
		names = append(names, name)
	}
	return names
}

func TrimToBagName(bag string) string {
	re := regexp.MustCompile("(^[0-9]*)")
	parsedBag := string(re.ReplaceAll([]byte(bag), []byte("")))
	parsedBag = strings.TrimSpace(parsedBag)
	return parsedBag
}

func GetBagNumber(bag string) int {
	re := regexp.MustCompile("(^[0-9]*)")
	foundNumber, _ := strconv.Atoi(re.FindString(bag))
	return foundNumber
}

func RecursiveAddBags(allBags [][]string, bag bagType) (bagType, error) {
	subBags := ListContaining(allBags, bag.name)

	for _, bagString := range subBags {
		if bagString == "" {
			return bag, errors.New("End of recursion")
		}
		//fmt.Printf("Bag: %s \n", bagString)

		newBag := bagType{
			TrimToBagName(bagString),
			make([]bagType, 0),
			GetBagNumber(bagString),
		}

		subBag, err := RecursiveAddBags(allBags, newBag)
		if err == nil {
			bag.subBags = append(bag.subBags, subBag)
		}
	}

	return bag, nil
}

// Part 1
// -------------

func GetContainingBags(bags [][]string, bag string) []string {
	baseContained := ListContained(bags, bag)
	allContained := TraverseBagTreeUp(bags, baseContained)
	allContained = PruneListToUniques(allContained)
	return allContained
}

func TraverseBagTreeUp(allBags [][]string, baseBags []string) []string {
	containedBags := make([]string, 0)

	for _, bag := range baseBags {
		containedBags = append(containedBags, PruneBagName(bag))
	}

	for {
		bagsSize := len(containedBags)
		for _, bag := range containedBags {
			newBags := ListContained(allBags, bag)
			for _, newBag := range newBags {
				newBag = PruneBagName(newBag)
				containedBags = append(containedBags, newBag)
			}
		}
		containedBags = PruneListToUniques(containedBags)

		if bagsSize == len(containedBags) {
			break
		}
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
