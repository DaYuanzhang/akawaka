package runner

import "fmt"

func ShowBanner() {
	banner := `
NAME:
   Akawaka - 阿军的文件内容搜索小工具

USAGE:
   Akawaka [global options] command [command options] [arguments...]

VERSION:
   %s

COMMANDS:
   help, h  Shows a list of commands or help for one command

`
	fmt.Printf(banner, ShowVersion())
}

func ShowVersion() string {
	return "0.0.1"
}
