package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file := "day11/input.csv"
	TraverseFile(file)
}

func TraverseFile(filename string) {
	seats := make([]string, 0)

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
		row := scanner.Text()
		seats = append(seats, row)
	}

	fmt.Printf("Part 1 \n")
	Part1(seats)
	fmt.Printf("Part 2 \n")
	Part2(seats)

}

func Part2(seats []string) {

	for {
		state := strings.Join(seats, "")
		paddedSeats := GetPaddedSeats(seats)
		newSeating := make([]string, 0)
		//PrintSeats(seats)

		for row, seatingRow := range seats {
			newRow := ""
			for column, seat := range seatingRow {
				surroundingSeats := GetSightedSurroundingSeats(paddedSeats, row+1, column+1)
				newSeat := ChangeState(seat, surroundingSeats, 1, 5)
				newRow += string(newSeat)
			}
			newSeating = append(newSeating, newRow)
		}
		// fmt.Printf("NEW ROUND\n\n")

		if state == strings.Join(newSeating, "") {
			break
		} else {
			seats = newSeating
		}
	}

	occupiedSeats := strings.Count(strings.Join(seats, ""), "#")
	fmt.Printf("Number of occupied seats are: %d \n", occupiedSeats)

}

func Part1(seats []string) {

	for {
		state := strings.Join(seats, "")
		paddedSeats := GetPaddedSeats(seats)
		newSeating := make([]string, 0)
		// PrintSeats(seats)

		for row, seatingRow := range seats {
			newRow := ""
			for column, seat := range seatingRow {
				surroundingSeats := GetImmediateSurroundingSeats(paddedSeats, row+1, column+1)
				newSeat := ChangeState(seat, surroundingSeats, 1, 4)
				newRow += string(newSeat)
			}
			newSeating = append(newSeating, newRow)
		}
		// fmt.Printf("NEW ROUND\n\n")

		if state == strings.Join(newSeating, "") {
			break
		} else {
			seats = newSeating
		}
	}

	occupiedSeats := strings.Count(strings.Join(seats, ""), "#")
	fmt.Printf("Number of occupied seats are: %d \n", occupiedSeats)
}

func SetRow(seat rune, row string, index int) string {
	row = row[0:index] + string(seat) + row[index+1:len(row)]
	return row
}

func GetPaddedSeats(seats []string) []string {
	paddedSeats := make([]string, 0)
	mark := "!"

	paddedSeats = append(paddedSeats, strings.Repeat(mark, len(seats[0])+2))
	for _, row := range seats {
		paddedSeats = append(paddedSeats, mark+row+mark)
	}
	paddedSeats = append(paddedSeats, strings.Repeat(mark, len(seats[0])+2))

	return paddedSeats
}

func ChangeState(seat rune, surrounding string, crowdedLimit int, emptyLimit int) rune {
	if seat == '.' {
		return seat
	} else if seat == '#' && strings.Count(surrounding, "#") >= emptyLimit {
		return 'L'
	} else if seat == 'L' && strings.Count(surrounding, "#") < crowdedLimit {
		return '#'
	} else {
		return seat
	}
}

func GetImmediateSurroundingSeats(seats []string, row int, column int) string {
	surroundingSeats := ""

	surroundingSeats += seats[row-1][column-1 : column+2]
	surroundingSeats += string(seats[row][column-1])
	surroundingSeats += string(seats[row][column+1])
	surroundingSeats += seats[row+1][column-1 : column+2]

	return surroundingSeats
}

func GetSightedSurroundingSeats(seats []string, row int, column int) string {
	i := 0
	characters := []rune{'#', 'L', '!'}
	surroundingSeats := ""
	tl, t, tr, l, r, bl, b, br := "", "", "", "", "", "", "", ""

	for {
		i++
		if tl == "" {
			if Equals(seats[row-i][column-i], characters...) {
				tl = string(seats[row-i][column-i])
				surroundingSeats += tl
			}
		}
		if t == "" {
			if Equals(seats[row-i][column], characters...) {
				t = string(seats[row-i][column])
				surroundingSeats += t
			}
		}
		if tr == "" {
			if Equals(seats[row-i][column+i], characters...) {
				tr = string(seats[row-i][column+i])
				surroundingSeats += tr
			}
		}
		if l == "" {
			if Equals(seats[row][column-i], characters...) {
				l = string(seats[row][column-i])
				surroundingSeats += l
			}
		}
		if r == "" {
			if Equals(seats[row][column+i], characters...) {
				r = string(seats[row][column+i])
				surroundingSeats += r
			}
		}
		if bl == "" {
			if Equals(seats[row+i][column-i], characters...) {
				bl = string(seats[row+i][column-i])
				surroundingSeats += bl
			}
		}
		if b == "" {
			if Equals(seats[row+i][column], characters...) {
				b = string(seats[row+i][column])
				surroundingSeats += b
			}
		}
		if br == "" {
			if Equals(seats[row+i][column+i], characters...) {
				br = string(seats[row+i][column+i])
				surroundingSeats += br
			}
		}

		if len(surroundingSeats) == 8 {
			break
		}
	}

	return surroundingSeats
}

func Equals(base byte, chars ...rune) bool {
	for _, char := range chars {
		if string(base) == string(char) {
			return true
		}
	}
	return false
}

func PrintSeats(seats []string) {
	fmt.Printf("Seats: \n")
	for _, row := range seats {
		fmt.Printf("%s \n", row)
	}
	fmt.Printf("\n")

}
