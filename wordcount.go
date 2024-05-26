package main

import (
	"bufio"
	"flag"
	"fmt"
	"hash/adler32"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz"

func generateWord(rnd *rand.Rand, length int) []byte {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rnd.Intn(len(letterBytes))]
	}
	return b
}

func GenerateData(seed int64, words, size int) map[string]int {
	rnd := rand.New(rand.NewSource(seed))
	wordCount := make(map[string]int)
	hasher := adler32.New()
	totalCount := 0
	for i := 0; i < words; i++ {
		word := generateWord(rnd, rnd.Intn(32))
		hasher.Write(word)
		wordTotal := int(hasher.Sum32())
		totalCount += wordTotal
		wordCount[string(word)] = wordTotal
	}

	// Now, adjust the counts to make the total count equal to size
	scale := float64(size) / float64(totalCount)
	for word, count := range wordCount {
		wordCount[word] = int(float64(count) * scale)
	}

	return wordCount
}

func PrintData(data map[string]int, seed int64) {
	rnd := rand.New(rand.NewSource(seed))
	buffer := bufio.NewWriter(os.Stdout)

	for len(data) > 0 {
		for word, count := range data {
			buffer.WriteString(word)

			if count == 1 {
				delete(data, word)
			} else {
				data[word] = count - 1
			}

			if rnd.Intn(100) == 0 {
				buffer.WriteString("\n")
			} else {
				buffer.WriteString(" ")
			}

			break
		}
	}

	buffer.Flush()
}

func validator(data map[string]int) {
	// The format of the input is a list of words and count, one per line, separated by a colon.
	// For example:
	// word1: 100
	// word2: 200
	// The output may be huge, so we will validate it on the fly.

	previousWord := ""

	// Read lines from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line into word and count
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			// Invalid line format, handle error
			continue
		}
		word := strings.TrimSpace(parts[0])
		countStr := strings.TrimSpace(parts[1])
		count, err := strconv.Atoi(countStr)
		if err != nil {
			// Invalid count format
			fmt.Fprintf(os.Stderr, "Invalid count: %s. Expected a number.\n", countStr)
		}

		expectedTotal, ok := data[word]

		if !ok {
			// Word not found in data
			fmt.Fprintf(os.Stderr, "Word not found: %s\n", word)
		}

		if expectedTotal != count {
			// Incorrect count
			fmt.Fprintf(os.Stderr, "Incorrect count for word %s. Expected %d, got %d\n", word, expectedTotal, count)
		}

		// Check if the words are sorted
		if previousWord > word {
			// Words are not sorted
			fmt.Fprintf(os.Stderr, "Words are not sorted. Previous word: %s, current word: %s\n", previousWord, word)
		}

		previousWord = word
	}
}

func main() {
	mode := flag.String("mode", "generator", "Mode: generator or validator")
	seed := flag.Int64("seed", 0, "Seed for random number generator")
	words := flag.Int("words", 10000, "Number of different words to generate")
	size := flag.Int("size", 1000000, "Number of words to generate")
	flag.Parse()

	data := GenerateData(*seed, *words, *size)

	switch *mode {
	case "generator":
		PrintData(data, *seed)
	case "validator":
		validator(data)
	default:
		fmt.Println("Invalid mode. Use generator or validator")
	}
}
