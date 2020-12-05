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
	file := "day4/input.csv"

	valid := ScanPair(file)

	fmt.Printf("The amount of valid passports are: %d \n", valid)
}

func ScanPair(filename string) int {
	valid := 0
	text := ""

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

		line := scanner.Text()
		if len(strings.TrimSpace(line)) != 0 {
			text += line + " "
			continue
		} else {
			if IsPassportValid(text) {
				fmt.Printf("%d : %#v \n\n", valid, text)
				valid++
			}
			text = ""
		}
	}
	return valid
}

func IsPassportValid(text string) bool {
	//fmt.Printf("sent in: %v\n", text)
	if len(strings.TrimSpace(text)) == 0 {
		return false
	}
	valid := PropertyValid(text)

	return valid
}

func PropertyValid(text string) bool {
	valid := true
	reBirthYear := regexp.MustCompile(`(byr:(19[2-9][0-9]|200[0-2]) )`)
	reIssueyear := regexp.MustCompile(`(iyr:20(20|1[0-9]) )`)
	reExpirationYear := regexp.MustCompile(`(eyr:20(30|2[0-9]) )`)
	reHeight := regexp.MustCompile(`(hgt:(1([5-8][0-9]|9[0-3])cm)|((59|6[0-9]|7[0-6])in) )`)
	reHairColor := regexp.MustCompile(`(hcl:#[0-9a-f]{6} )`)
	reEyeColor := regexp.MustCompile(`(ecl:((amb)|(blu)|(brn)|(gry)|(grn)|(hzl)|(oth)) )`)
	rePassportID := regexp.MustCompile(`(pid:[0-9]{9} )`)

	validations := []*regexp.Regexp{reBirthYear, reIssueyear, reExpirationYear, reHeight, reHairColor, reEyeColor, rePassportID}

	for _, regex := range validations {
		if !regex.Match([]byte(text)) {
			//fmt.Printf("Regex %#v did not match %#v\n\n", regex.String(), text)
			return false
		}
	}

	return valid
}
