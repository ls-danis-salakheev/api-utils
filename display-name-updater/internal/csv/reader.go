package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func LoadCsvLines(fileName string) (csvLines [][]string, errOccurred bool) {
	file, err := os.Open(fileName)
	defer closeFile(file)
	if err != nil {
		fmt.Println("Could load a file:", err)
		return nil, true
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Could not read lines from the csv file:", err)
		return nil, true
	}
	return lines, false
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Println("Could not close a file:", err)
	}
}
