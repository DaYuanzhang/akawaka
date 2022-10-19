package utils

import (
	"akawaka/pkg/config"
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var extens = []string{}
var keywords = []string{}

func GetFileList(rootPath string) ([]string, error) {

	var files []string
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		return []string{}, err
	}
	/*
		for _, file := range files {
			fmt.Println(file)
		}

	*/
	return files, err
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

func IsDir(filepath string) bool {
	s, err := os.Stat(filepath)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func InExtens(fileName string) bool {
	fileSuffix := path.Ext(fileName)
	for _, e := range extens {
		if fileSuffix == "."+e {
			return true
		}
	}
	return false
}

func ReadFile(fileName string) (string, error) {
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

func Search(filePath string) {
	content, err := ReadFile(filePath)
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

func IsValid(dirPath string) error {
	_, err := os.Open(dirPath)
	return err
}

func Init(options *config.Options) {
	keywords = options.Keywords
	extens = options.Extensions
}
