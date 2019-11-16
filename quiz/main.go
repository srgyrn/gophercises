package quiz

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	score int
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
	questions, err := csv.NewReader(bufio.NewReader(csvFile)).ReadAll()

	if !errors.Is(err, nil) {
		fmt.Print("cannot read questions from file")
	}

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