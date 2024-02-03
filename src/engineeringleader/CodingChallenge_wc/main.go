package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var command string

	flag.StringVar(&command, "c", "Text.txt", "Please provide a file to calculate the byte size")
	flag.StringVar(&command, "w", "Text.txt", "Please provide a file to calculate the number of words")
	flag.StringVar(&command, "m", "Text.txt", "Please provide a file to calculate the number of characters")

	flag.Parse()
	FileCalc(command)
	CountChars(command)
}

func FileCalc(file string) {
	var wordCounts int

	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "file: %v\n", err)
	}
	fmt.Printf("The file has the size of %d bytes \n", len(data))

	for _, line := range strings.Split(string(data), "\n") {
		wordCounts += len(strings.Fields(line))

	}

	fmt.Printf("The file has %d words.\n", wordCounts)

}

func CountChars(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Initialize a variable to store the total character count
	charCount := 0

	// Loop through each line
	for scanner.Scan() {
		// Get the length of the current line and add it to the total count
		line := scanner.Text()
		charCount += len(line)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Print the result
	fmt.Println("The file has", charCount, "characters.")
}
