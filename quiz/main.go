package quiz

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func Start() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	workingDir, _ := os.Getwd()
	csvFile, err := os.Open(workingDir + "/quiz/" + *csvFileName)

	if !errors.Is(err, nil) {
		exit(fmt.Sprintf("%s not found in %s/quiz/\nError: %v", *csvFileName, workingDir, err))
	}

	questions, err := csv.NewReader(bufio.NewReader(csvFile)).ReadAll()

	if !errors.Is(err, nil) {
		exit("cannot read questions from file")
	}

	timer := time.NewTimer(time.Duration(*limit) * time.Second)

	var score int

problemLoop:
	for _, question := range questions {
		fmt.Printf("What %s is? ", question[0])
		answerCh := make(chan string)

		go func(){
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- strings.TrimSpace(answer)
		}()

		select {
		case <-timer.C:
			fmt.Print("\nTime's up!\n")
			break problemLoop
		case answer := <-answerCh:
			if answer == strings.TrimSpace(question[1]) {
				score++
			}
		}
	}

	fmt.Printf("Your scored %d out of %d.", score, len(questions))
}

func exit(message string) {
	fmt.Print(message)
	os.Exit(1)
}
