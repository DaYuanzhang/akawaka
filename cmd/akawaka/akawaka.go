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
	app.Usage = "阿军的文件内容搜索小工具"
	app.Version = runner.ShowVersion()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "extension", Aliases: []string{"e"}, Value: "txt", Destination: &options.Extens, Usage: "文件扩展名 eg:-e txt,jsp,asp"},
		&cli.StringFlag{Name: "directory", Aliases: []string{"d"}, Value: utils.GetWd(), Destination: &options.DirPath, Usage: "搜索目录 eg: -d D:\\web"},
		&cli.StringFlag{Name: "keyword", Aliases: []string{"k"}, Value: "", Destination: &options.Keyword, Usage: "搜索关键词 eg: -k keyword1,keyword2"},
		&cli.StringFlag{Name: "keyword-file", Aliases: []string{"kf"}, Value: "", Destination: &options.Keywords_File, Usage: "搜索关键词文本 eg: -kf keywords.txt"},
	}

	app.Action = func(c *cli.Context) error {

		err := utils.Transform(options)
		if err != nil {
			return err
		}

		err = runner.New(options)
		return err
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("[!] start akawaka failed,", err.Error())
	}

}