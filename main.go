package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func readCSV(filepath string) ([][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	utf16Reader := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	reader := csv.NewReader(transform.NewReader(bufio.NewReader(file), utf16Reader.NewDecoder()))

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func main() {
	records, err := readCSV("test_file.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, record := range records {
		fmt.Println(record)
	}
}
