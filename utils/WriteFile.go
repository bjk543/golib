package utils

import (
	"fmt"
	"os"
	"strings"
)

func SaveToFile(filename string, ss []string) {
	WriteFile(filename, "")
	AppendFile(filename, ss)
}
func WriteFile(filename string, s string) {
	path := filename[:strings.LastIndex(filename, "/")]
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.WriteString(s)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
}

func AppendFile(filename string, ss []string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	for _, s := range ss {

		if _, err = f.WriteString(s); err != nil {
			panic(err)
		}
		f.WriteString("\n")
	}

}
