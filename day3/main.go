package main

import (
	"bufio"
	"log"
	"os"
)

var (
	matchValue      int
	totalMatchValue int
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening input file: %v", err)
	}
	defer input.Close()
	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	priorities := mapPriorities()

	for fileScanner.Scan() {
		compartment1 := fileScanner.Text()[:len(fileScanner.Text())/2]
		compartment2 := fileScanner.Text()[len(fileScanner.Text())/2:]
		for _, char1 := range compartment1 {
			for _, char2 := range compartment2 {
				if string(char1) == string(char2) {
					matchValue = priorities[string(char1)]
				}
			}
		}
		totalMatchValue += matchValue
		matchValue = 0
	}
	log.Printf("Total match value: %v", totalMatchValue)
}

func mapPriorities() map[string]int {
	return map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
		"A": 27,
		"B": 28,
		"C": 29,
		"D": 30,
		"E": 31,
		"F": 32,
		"G": 33,
		"H": 34,
		"I": 35,
		"J": 36,
		"K": 37,
		"L": 38,
		"M": 39,
		"N": 40,
		"O": 41,
		"P": 42,
		"Q": 43,
		"R": 44,
		"S": 45,
		"T": 46,
		"U": 47,
		"V": 48,
		"W": 49,
		"X": 50,
		"Y": 51,
		"Z": 52,
	}
}
