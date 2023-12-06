package utils

import (
	"bufio"
	"os"
)

func ReadFile(path string) []string {
	readFile, err := os.Open(path)
	Check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	readFile.Close()

	return lines
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
