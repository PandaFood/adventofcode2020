package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file := "dayx/input.csv"
	TraverseFile(file)
}

func TraverseFile(filename string) {

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

	}

}
