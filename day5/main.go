package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file := "day5/input.csv"
	TraverseFile(file)
}

func TraverseFile(filename string) {
	maxSeatID := 0
	seats := make([]int, 1)

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
		text := scanner.Text()

		if len(strings.TrimSpace(text)) == 0 {
			return
		}

		row := GetRow(text[0:7])
		column := GetColumn(text[7:10])
		seatID := row*8 + column
		seats = append(seats, seatID)
		//fmt.Printf("Seating: %d - %d - ID:%v\n", row, column, seatID)
		//fmt.Printf("--------- \n")

		if maxSeatID < seatID {
			maxSeatID = seatID
			//fmt.Printf("New highest seatID: %v \n", maxSeatID)
		}
	}
	//fmt.Printf("Seats: %#v \n", seats)
	mySeat := FindMySeat(seats)
	fmt.Printf("My seat is: %d \n", mySeat)
}

func GetRow(row string) int {
	lowIndex := 0
	highIndex := 127

	for _, letter := range row {
		if string(letter) == "B" {
			lowIndex += int(float32(((highIndex + 1) - lowIndex) / 2))
		} else if string(letter) == "F" {
			highIndex -= int(float32(((highIndex + 1) - lowIndex) / 2))
		}
		//fmt.Printf("Low: %d High: %d \n", lowIndex, highIndex)
	}
	return highIndex
}

func GetColumn(column string) int {
	lowIndex := 0
	highIndex := 7

	for _, letter := range column {
		if string(letter) == "R" {
			lowIndex += int(float32(((highIndex + 1) - lowIndex) / 2))
		} else if string(letter) == "L" {
			highIndex -= int(float32(((highIndex + 1) - lowIndex) / 2))
		}
		//fmt.Printf("Low: %d High: %d \n", lowIndex, highIndex)
	}
	return highIndex
}

func FindMySeat(list []int) int {
	for _, seatID := range list {
		if !IsIncluded(seatID+1, list) && IsIncluded(seatID+2, list) {
			return seatID + 1
		}
	}
	return 0
}

func IsIncluded(id int, list []int) bool {
	for _, val := range list {
		if val == id {
			return true
		}
	}
	return false
}
