package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

const QUESTION_INDEX = 0
const ANSWER_INDEX = 1

var problemsFlag = flag.String("problems", "problems.csv", "a csv file in the format of 'question,answer'")

func main() {
	flag.Parse()

	fmt.Println("Welcome to the quiz game!")
	fmt.Println("You will be given a series of questions and you will have to answer them.")
	fmt.Println("Let's start!")

	records := readCsvFile(*problemsFlag)
	totalQuestions := len(records)
	correctAnswers := 0
	for _, record := range records {
		fmt.Println(record[QUESTION_INDEX])
		var answer string
		fmt.Scanln(&answer)
		if compareStrings(answer, record[ANSWER_INDEX]) {
			fmt.Println("Correct!")
			correctAnswers++
		} else {
			fmt.Println("Incorrect!")
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correctAnswers, totalQuestions)
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Unable to read input file"+filePath, err)
		os.Exit(1)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Unable to parse file as CSV for "+filePath, err)
		os.Exit(1)
	}

	return records
}

func normalizeString(s string) string {
	return strings.TrimSpace(strings.ToLower(s))
}

func compareStrings(a, b string) bool {
	return normalizeString(a) == normalizeString(b)
}
