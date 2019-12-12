package fileReader

import (
	"bufio"
	"log"
	"os"
)

func ReadLines(fileName string) []string {

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	return txtlines
}

//ReadLine returns a single line from the read file at the index point
func ReadLine(fileName string, index int) string {
	lines := ReadLines(fileName)
	return lines[index]
}
