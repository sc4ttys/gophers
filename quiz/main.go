package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//fmt.Print("Starting quiz:")
	var score int
	var usera string
	q, a := readCsv()
	for i, question := range q {
		fmt.Print(question, "=")
		fmt.Scan(&usera)
		if usera == a[i] {
			score += 1
		}
	}
	fmt.Print("Final score: ", score)
}

func readCsv() (questions []string, answers []string) {
	csvFile, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalln("Couldn't open the file", err)
	}

	questions = []string{}
	answers = []string{}

	rder := csv.NewReader(csvFile)
	for {
		record, err := rder.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Error readng the file", err)
		}
		questions = append(questions, record[0])
		answers = append(answers, record[1])
	}
	return questions, answers
}
