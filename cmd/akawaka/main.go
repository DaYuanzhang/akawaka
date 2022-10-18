package main

import (
	"akawaka/pkg/config"
	"akawaka/pkg/runner"
	"akawaka/pkg/utils"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

var options = &config.Options{}

func main() {

	app := cli.NewApp()
	app.Name = "Akawaka"
	app.Usage = "hello world"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "extension", Aliases: []string{"e"}, Value: "txt", Destination: &options.Extens, Usage: "文件扩展名 eg:-e txt,jsp,asp"},
		&cli.StringFlag{Name: "directory", Aliases: []string{"d"}, Value: utils.GetWd(), Destination: &options.DirPath, Usage: "搜索目录 eg: -d D:\\web"},
		&cli.StringFlag{Name: "keyword", Aliases: []string{"k"}, Value: "", Destination: &options.Keyword, Usage: "搜索关键词 eg: -k keyword1,keyword2"},
		&cli.StringFlag{Name: "keyword-file", Aliases: []string{"kf"}, Value: "", Destination: &options.Keywords_File, Usage: "搜索关键词文本 eg: -kf keywords.txt"},
	}

	app.Action = func(c *cli.Context) error {
		utils.ShowBanner2()
		/*
			if len(options.Keyword) == 0 || len(options.Keywords_File) == 0 {
				return errors.New("no keywords")
			}

		*/
		if len(options.Keyword) != 0 {
			if strings.Contains(options.Keyword, ",") {
				options.SetKeywords()
			} else {
				options.Keywords = append(options.Keywords, options.Keyword)
			}
			if strings.Contains(options.Extens, ",") {
				options.SetExtensions()
			} else {
				options.Extensions = append(options.Extensions, options.Extens)
			}
			err := runner.New(options)
			return err
		} else if len(options.Keywords_File) != 0 {
			var err error
			options.Keywords, err = utils.ReadArrFromTxt(options.Keywords_File)
			if err != nil {
				return err
			}
			if strings.Contains(options.Extens, ",") {
				options.SetExtensions()
			} else {
				options.Extensions = append(options.Extensions, strings.TrimSpace(options.Extens))
			}
			fmt.Println(options.Keywords)
			err = runner.New(options)
			return err
		}

		return nil
	}

	utils.ShowBanner()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("start afrog failed, ", err.Error())
	}

	/*
		utils.ShowBanner()
		fmt.Printf("hello world!\n")
		print(isDir("D:\\Tools\\框架漏洞\\综合扫描器\\afrog\\pocs\\v\\afrog.version"))
		run("D:\\Tools\\框架漏洞\\综合扫描器\\afrog")

	*/

}
