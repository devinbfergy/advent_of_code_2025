package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func modOneHundred(number int) int {
	if number%100 == 0 && number != 0 {
		return 0
	} else {
		return number % 100
	}
}

func readInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("input readlines error: %s", err)
	}
	var zeroCount int

	startingSpot := 50
	for i, line := range lines {
		fmt.Println(i, line)
		direction := string(line[0])
		moves, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalf("input readlines error: %s", err)
		}
		switch direction {
		case "L":
			startingSpot = startingSpot - moves
		case "R":
			startingSpot = startingSpot + moves
		}
		startingSpot = modOneHundred(startingSpot)
		if startingSpot == 0 {
			zeroCount = zeroCount + 1
		}
	}

	fmt.Printf("The password is %d\n", zeroCount)
}
