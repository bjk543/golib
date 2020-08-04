package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func MainReadFile(filename string) string {
	b, err := ioutil.ReadFile(filename) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	// fmt.Println(b) // print the content as 'bytes'

	str := string(b) // convert content to a 'string'

	// fmt.Println(str) // print the content as a 'string'
	return str
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
