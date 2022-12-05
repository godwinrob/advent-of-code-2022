package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

var mostCalories int
var secondMostCalories int
var thirdMostCalories int
var currentElfTotal int
var currentFoodCalories int

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening input file: %v", err)
	}
	defer input.Close()

	fileScanner := bufio.NewScanner(input)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			if currentElfTotal >= thirdMostCalories {
				if currentElfTotal >= secondMostCalories {
					if currentElfTotal >= mostCalories {
						thirdMostCalories = secondMostCalories
						secondMostCalories = mostCalories
						mostCalories = currentElfTotal
					} else {
						thirdMostCalories = secondMostCalories
						secondMostCalories = currentElfTotal
					}
				} else {
					thirdMostCalories = currentElfTotal
				}
			}

			currentElfTotal = 0
			continue
		} else {
			currentFoodCalories, err = strconv.Atoi(fileScanner.Text())
			if err != nil {
				log.Println(err)
			}
			currentElfTotal += currentFoodCalories
		}
	}

	log.Printf("Most calories: %v", mostCalories)
	log.Printf("Second most calories: %v", secondMostCalories)
	log.Printf("Third most calories: %v", thirdMostCalories)
	log.Printf("top three total: %v", mostCalories+secondMostCalories+thirdMostCalories)
}
