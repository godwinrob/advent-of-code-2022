package main

import (
	"bufio"
	"log"
	"os"
)

var (
	totalScore1 int
	totalScore2 int
	roundScore  int
	elfPlay     byte
	myPlay      byte
)

func main() {

	total1 := part1()
	total2 := part2()

	log.Printf("Total part 1 score: %v", total1)
	log.Printf("Total part 2 score: %v", total2)

}

func part1() int {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening input file: %v", err)
	}
	defer input.Close()
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
		totalScore1 += roundScore
		roundScore = 0
	}
	return totalScore1
}

func part2() int {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening input file: %v", err)
	}
	defer input.Close()
	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		elfPlay = fileScanner.Text()[0]
		myPlay = fileScanner.Text()[2]

		if elfPlay == 'A' { // Rock
			if myPlay == 'X' {
				// Lose Scissors with 3 point
				roundScore = 3
			} else if myPlay == 'Y' {
				// Draw Rock with 1 points
				roundScore = 4
			} else {
				// Win Paper with 2 points
				roundScore = 8
			}
		} else if elfPlay == 'B' { // Paper
			if myPlay == 'X' {
				// Loss Rock with 1 point
				roundScore = 1
			} else if myPlay == 'Y' {
				// Draw Paper with 2 points
				roundScore = 5
			} else {
				// Win Scissors with 3 points
				roundScore = 9
			}
		} else if elfPlay == 'C' { // Scissors
			if myPlay == 'X' {
				// Lose paper with 2 point
				roundScore = 2
			} else if myPlay == 'Y' {
				// Draw Scissors with 3 point
				roundScore = 6
			} else {
				// Win Rock with 1 points
				roundScore = 7
			}
		}
		totalScore2 += roundScore
		roundScore = 0
	}
	return totalScore2
}
