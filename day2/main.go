package main

import (
	"bufio"
	"log"
	"os"
)

var (
	totalScore int
	roundScore int
	elfPlay    byte
	myPlay     byte
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening input file: %v", err)
	}
	defer input.Close()

	part1(input)

}

func part1(input *os.File) {

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		elfPlay = fileScanner.Text()[0]
		myPlay = fileScanner.Text()[2]

		if elfPlay == 'A' {
			if myPlay == 'X' {
				// Draw with 1 point
				roundScore = 4
			} else if myPlay == 'Y' {
				// Win with 2 points
				roundScore = 8
			} else {
				// Loss with 3 points
				roundScore = 3
			}
		} else if elfPlay == 'B' {
			if myPlay == 'X' {
				// Loss with 1 point
				roundScore = 1
			} else if myPlay == 'Y' {
				// Draw with 2 points
				roundScore = 5
			} else {
				// Win with 3 points
				roundScore = 9
			}
		} else {
			if myPlay == 'X' {
				// Win with 1 point
				roundScore = 7
			} else if myPlay == 'Y' {
				// Lose with 2 points
				roundScore = 2
			} else {
				// Draw with 3 points
				roundScore = 6
			}
		}
		totalScore += roundScore
		roundScore = 0
	}
	log.Printf("Total part 1 score: %v", totalScore)
}

func part2(input *os.File) {

}
