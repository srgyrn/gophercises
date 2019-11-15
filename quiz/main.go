package quiz

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"sync"
	"strings"
)

var (
	score int
	wrongAnswers int
	wg sync.WaitGroup
	buf *bufio.Reader
)

func Start() {
	workingDir, _ := os.Getwd()
	csvFile, err := os.Open(workingDir+"/quiz/problems.csv")
	wg.Add(1)
	buf = bufio.NewReader(os.Stdin)

	if !errors.Is(err, nil) {
		fmt.Println(err, workingDir)
	}
	questions := csv.NewReader(bufio.NewReader(csvFile))

	go askQuestion(questions)
	wg.Wait()

	fmt.Printf("Your score: %d. Questions answered wrong: %d", score, wrongAnswers)
}

func askQuestion(questions *csv.Reader) {
	defer wg.Done()

	for {
		question, err := questions.Read()

		if !errors.Is(err, nil) {
			break
		}

		fmt.Printf("What %s is? ", question[0])
		answer, _ := buf.ReadString('\n')

		if strings.TrimSpace(answer) == strings.TrimSpace(question[1]) {
			score++
		} else {
			wrongAnswers++
		}
	}
}

func timer() {

}