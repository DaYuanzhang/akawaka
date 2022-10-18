package runner

import (
	"akawaka/pkg/config"
	"akawaka/pkg/utils"
	"fmt"
)

type Runner struct {
	options *config.Options
}

func New(options *config.Options) error {
	// 验证文件目录路径是否正确
	err := utils.IsValid(options.DirPath)
	if err != nil {
		return err
	}

	// 初始化参数
	utils.Init(options)

	files, _ := utils.GetFileList(options.DirPath)
	fmt.Printf("[!] starting search...\n")
	for _, filePath := range files {
		if utils.IsDir(filePath) {
			continue
		} else {
			if utils.InExtens(filePath) {
				utils.Search(filePath)
			}
		}
	}
	fmt.Printf("[!] search finished...\n")
	return nil
}
