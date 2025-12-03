package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

// Source - https://stackoverflow.com/a
// Posted by Igor Mikushkin, modified by community. See post 'Timeline' for change history
// Retrieved 2025-12-02, License - CC BY-SA 4.0

func Chunks(s string, chunkSize int) []string {
	if len(s) == 0 {
		return nil
	}
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == chunkSize {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])
	return chunks
}

// Set represents a collection of unique elements.
type Set struct {
	m map[interface{}]struct{}
}

// NewSet initializes a new set with the given elements.
func NewSet(items ...interface{}) *Set {
	s := &Set{
		m: make(map[interface{}]struct{}),
	}
	s.Add(items...)
	return s
}

// Add inserts elements into the set.
func (s *Set) Add(items ...interface{}) {
	for _, item := range items {
		s.m[item] = struct{}{}
	}
}

// Has checks if an element exists in the set.
func (s *Set) Has(item interface{}) bool {
	_, found := s.m[item]
	return found
}

// Remove deletes an element from the set.
func (s *Set) Remove(item interface{}) {
	delete(s.m, item)
}

// Size returns the number of elements in the set.
func (s *Set) Size() int {
	return len(s.m)
}

func main() {
	lines, err := readInput("input2.txt")
	if err != nil {
		log.Fatalf("input readlines error: %s", err)
	}
	idsStr := strings.Split(lines[0], ",")
	key := 0
	key2 := 0

	for _, ids := range idsStr {
		idsRange := strings.Split(ids, "-")
		startingID, _ := strconv.Atoi(idsRange[0])
		endingID, _ := strconv.Atoi(idsRange[1])

		for idInt := startingID; idInt <= endingID; idInt += 1 {
			id := strconv.Itoa(idInt)

			// --- Key 1 Logic (unchanged) ---
			for i := range len(id)/2 + 1 {
				substr := id[0:i]
				if substr == id[i:] {
					// fmt.Println("id match:", substr, id[i:])
					key += idInt
					// fmt.Println("key:", key)
					break
				}
			}

			// --- Key 2 Logic (Fixed) ---
			for j := 1; j <= len(id)/2; j++ {
				// Optimization: If the length isn't divisible by j,
				// it can't be made ONLY of that sequence.
				if len(id)%j != 0 {
					continue
				}

				pattern := id[:j]
				count := len(id) / j

				// Check if the pattern repeated matches the ID
				if strings.Repeat(pattern, count) == id {
					// fmt.Println("id match key2:", id)
					key2 += idInt
					break // Count this ID only once
				}
			}
		}
	}
	fmt.Println("key:", key)
	fmt.Println("key2:", key2)
}
