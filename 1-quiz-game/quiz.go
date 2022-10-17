package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

var fileName string = "problems.csv"

func loadData(file string) [][]string {
	fileHandle, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandle.Close()

	lines, err := csv.NewReader(fileHandle).ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return lines
}

func playGame(lines [][]string) {
	var guess string
	correct := 0
	for _, line := range lines {
		fmt.Printf("%s ? ", line[0])
		fmt.Scanf("%s\n", &guess)
		if strings.ToUpper(guess) == strings.ToUpper(line[1]) {
			correct++
		}
	}
	fmt.Printf("You guessed correctly %d out of %d (%.2f%%)", correct, len(lines), 100*float32(correct)/float32(len(lines)))
}

func main() {

	lines := loadData(fileName)
	playGame(lines)
}
