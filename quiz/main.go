package quiz

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

func Start() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	workingDir, _ := os.Getwd()
	csvFile, err := os.Open(workingDir + "/quiz/" + *csvFileName)

	var buf *bufio.Reader
	buf = bufio.NewReader(os.Stdin)
	if !errors.Is(err, nil) {
		exit(fmt.Sprintf("%s not found in %s/quiz/\nError: %v", *csvFileName, workingDir, err))
	}

	questions, err := csv.NewReader(bufio.NewReader(csvFile)).ReadAll()

	if !errors.Is(err, nil) {
		exit("cannot read questions from file")
	}
	var wg sync.WaitGroup
	wg.Add(1)

	var score int
	go func() {
		defer wg.Done()

		for _, question := range questions {
			fmt.Printf("What %s is? ", question[0])
			answer, _ := buf.ReadString('\n')

			if strings.TrimSpace(answer) == strings.TrimSpace(question[1]) {
				score++
			}
		}
	}()
	wg.Wait()

	fmt.Printf("Your scored %d out of %d.", score, len(questions))
}

func exit(message string) {
	fmt.Print(message)
	os.Exit(1)
}
