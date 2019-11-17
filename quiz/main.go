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

var (
	score int
	wg    sync.WaitGroup
	buf   *bufio.Reader
)

func Start() {
	workingDir, _ := os.Getwd()
	csvFile, err := os.Open(workingDir + "/quissdfz/problems.csv")

	flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	buf = bufio.NewReader(os.Stdin)
	if !errors.Is(err, nil) {
		fmt.Printf("problems.csv not found in %s/quiz/\nError: %v", workingDir, err)
		os.Exit(1)
	}

	questions, err := csv.NewReader(bufio.NewReader(csvFile)).ReadAll()

	if !errors.Is(err, nil) {
		fmt.Print("cannot read questions from file")
	}

	wg.Add(1)

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

func timer() {

}
