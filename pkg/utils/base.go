package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ShowBanner() {
	fmt.Println("hello world")
}

func ShowBanner2() {
	fmt.Println("hello world")
}

func GetWd() string {
	dir, _ := os.Getwd()
	return dir
}

func ReadArrFromTxt(fileName string) ([]string, error) {
	var err error
	var arr = []string{}

	f, err := os.Open(fileName)
	if err != nil {
		return arr, err
	}

	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if len(line) != 0 {
			arr = append(arr, strings.TrimSpace(line))
		}
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
	}
	return arr, nil
}
