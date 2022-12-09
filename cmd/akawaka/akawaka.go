package main

import (
	"akawaka/pkg/config"
	"akawaka/pkg/runner"
	"akawaka/pkg/utils"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

var options = &config.Options{}

func main() {
	runner.ShowBanner()

	app := cli.NewApp()
	app.Name = "Akawaka"
	app.Usage = "腚的文件内容搜索小工具"
	app.Version = runner.ShowVersion()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "extension", Aliases: []string{"e"}, Value: "txt", Destination: &options.Extension, Usage: "文件扩展名 eg: -e txt,jsp,asp"},
		&cli.StringFlag{Name: "extension-file", Aliases: []string{"ef"}, Value: "", Destination: &options.Extensions_File, Usage: "文件扩展名文本 eg: -ef extens.txt"},
		&cli.StringFlag{Name: "directory", Aliases: []string{"d"}, Value: utils.GetWd(), Destination: &options.DirPath, Usage: "搜索目录 eg: -d D:\\web"},
		&cli.StringFlag{Name: "keyword", Aliases: []string{"k"}, Value: "", Destination: &options.Keyword, Usage: "搜索关键词 eg: -k keyword1,keyword2"},
		&cli.StringFlag{Name: "keyword-file", Aliases: []string{"kf"}, Value: "", Destination: &options.Keywords_File, Usage: "搜索关键词文本 eg: -kf keywords.txt"},
		&cli.BoolFlag{Name: "is-filename", Aliases: []string{"if"}, Value: false, Destination: &options.Is_Filename, Usage: "搜索文件名 eg: -if true"},
	}

	app.Action = func(c *cli.Context) error {

		// 转换参数
		err := utils.Transform(options)
		if err != nil {
			return err
		}

		// 运行搜索逻辑
		err = runner.New(options)
		return err
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(utils.LogColor.GetColor("Red", "[!] start akawaka failed,"+err.Error()))
	}

}
