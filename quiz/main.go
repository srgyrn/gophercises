package quiz

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

func Start() {
	workingDir, _ := os.Getwd()
	csvFile, err := os.Open(workingDir+"/quiz/problems.csv")

	if !errors.Is(err, nil) {
		fmt.Println(err, workingDir)
	}

	data := csv.NewReader(bufio.NewReader(csvFile))

	fmt.Print(data.ReadAll())
}