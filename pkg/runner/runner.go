package runner

import (
	"akawaka/pkg/config"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Runner struct {
	options *config.Options
}

var extens = []string{}
var keywords = []string{}

func getFileList(rootPath string) []string {

	var files []string
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	/*
		for _, file := range files {
			fmt.Println(file)
		}

	*/
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

func New(options *config.Options) error {
	rootPath := options.DirPath
	keywords = options.Keywords
	extens = options.Extensions

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
	return nil
}
