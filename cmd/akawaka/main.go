package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var extens = []string{"go"}
var keywords = []string{"package main"}

func getFileList(rootPath string) []string {

	var files []string
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
	return files
}

func isDir(filepath string) bool {
	s, err := os.Stat(filepath)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func inExtens(fileName string) bool {
	fileSuffix := path.Ext(fileName)
	for _, e := range extens {
		if fileSuffix == "."+e {
			return true
		}
	}
	return false
}

func readFile(fileName string) (string, error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return "", err
	} else {
		content := string(b[:])
		//fmt.Printf("b: %v\n", content)
		return content, nil
	}
}

func search(filePath string) {
	content, err := readFile(filePath)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	for _, keyword := range keywords {
		if strings.Contains(content, keyword) {
			fmt.Printf("[+] %v find keyword: \"%v\"\n", filePath, keyword)
			return
		}
	}
}

func run(rootPath string) {
	files := getFileList(rootPath)
	for _, filePath := range files {
		if isDir(filePath) {
			continue
		} else {
			if inExtens(filePath) {
				search(filePath)
			}
		}
	}
}

func main() {
	fmt.Printf("hello world!\n")
	print(isDir("D:\\Tools\\框架漏洞\\综合扫描器\\afrog\\pocs\\v\\afrog.version"))
	run("D:\\Tools\\框架漏洞\\综合扫描器\\afrog")
}
