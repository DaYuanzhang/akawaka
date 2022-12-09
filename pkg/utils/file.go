package utils

import (
	"akawaka/pkg/config"
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var extens = []string{}
var keywords = []string{}

/*
遍历目录文件名字，返回字符串数组
*/
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

/*
获取当前目录
*/
func GetWd() string {
	dir, _ := os.Getwd()
	return dir
}

/*
从文本中读出keywords并转换成字符串数组，每行一个元素
*/
func ReadArrFromTxt(fileName string) ([]string, error) {
	var err error
	var arr = []string{}
	fileSuffix := path.Ext(fileName)
	println("fileSuffix: ", fileSuffix, "\n")
	if fileSuffix != ".txt" {
		return arr, errors.New("file only supports txt")
	}

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

/*
判断是否为目录
*/
func IsDir(filepath string) bool {
	s, err := os.Stat(filepath)
	if err != nil {
		return false
	}
	return s.IsDir()
}

/*
判断后缀名是否在指定范围
*/
func InExtens(fileName string) bool {
	fileSuffix := path.Ext(fileName)
	for _, e := range extens {
		if fileSuffix == "."+e {
			return true
		}
	}
	return false
}

/*
读取文件内容
*/
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

/*
匹配内容关键字
*/
func Search(filePath string) {
	content, err := ReadFile(filePath)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	var msg string
	for _, keyword := range keywords {
		if strings.Contains(strings.ToLower(content), strings.ToLower(keyword)) {
			// msg = LogColor.GetColor("low", "[+] "+filePath+" find keyword: \""+keyword+"\":\n")
			msg = "[+] " + filePath + " find keyword: \"" + keyword + "\":\n"
			fmt.Printf(msg)
			outputContext(strings.ToLower(content), strings.ToLower(keyword))
		}
	}

}

/*
输出上下文, 并去除前后的换行符回车符
*/
func outputContext(content string, keyword string) {
	index := strings.Index(content, keyword)
	const (
		intel int = 10
		blank     = "\r\n"
	)

	content_len := len(content)
	keyword_len := len(keyword)

	if index != -1 {
		s := 0
		d := 0
		if index >= intel {
			s = index - intel
		} else {
			s = 0
		}
		if (content_len - index - keyword_len) >= intel {
			d = index + keyword_len + intel
		} else {
			d = content_len
		}

		if strings.Contains(content[s:index], blank) {
			s = s + strings.LastIndex(content[s:index], blank) + len(blank)
		}

		if strings.Contains(content[index+keyword_len:d], blank) {
			d = index + keyword_len + strings.Index(content[index+keyword_len:d], blank)
		}

		println(content[s:index] + LogColor.GetColor("Green", keyword) + content[index+keyword_len:d])
		next_content := content[index+keyword_len:]
		outputContext(next_content, keyword)
	}
}

/*
目录是否有效
*/
func IsValid(dirPath string) error {
	_, err := os.Open(dirPath)
	return err
}

/*
初始化参数
*/
func Init(options *config.Options) {
	keywords = options.Keywords
	extens = options.Extensions
}
