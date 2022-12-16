# akawaka
一个简单的文件内容搜索工具

支持指定后缀名搜索文件内容

可用于应急响应、内网终端信息收集、代码审计等

```
        __                                      __
       /\ \                                    /\ \
   __  \ \ \/'\      __     __  __  __     __  \ \ \/'\      __
 /'__`\ \ \ , <    /'__`\  /\ \/\ \/\ \  /'__`\ \ \ , <    /'__`\
/\ \L\.\_\ \ \\`\ /\ \L\.\_\ \ \_/ \_/ \/\ \L\.\_\ \ \\`\ /\ \L\.\_
\ \__/.\_\\ \_\ \_\ \__/.\_\\ \___x___/'\ \__/.\_\\ \_\ \_\ \__/.\_\
 \/__/\/_/ \/_/\/_/\/__/\/_/ \/__//__/   \/__/\/_/ \/_/\/_/\/__/\/_/

Author: dean
Version: 0.0.5

NAME:
   Akawaka - 腚的文件内容搜索小工具

USAGE:
   Akawaka [global options] command [command options] [arguments...]

VERSION:
   0.0.5

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --extension value, -e value         文件扩展名 eg: -e txt,jsp,asp
   --extension-file value, --ef value  文件扩展名文本 eg: -ef extens.txt
   --directory value, -d value         搜索目录 eg: -d D:\web (default: "D:\\Tools\\myTools\\Go\\akawaka\\cmd\\akawaka\\akawaka_releases_v0.0.5")
   --keyword value, -k value           搜索关键词 eg: -k keyword1,keyword2
   --keyword-file value, --kf value    搜索关键词文本 eg: -kf keywords.txt
   --is-filename, --if                 搜索文件名 eg: -if true (default: false)
   --verbose, -m                       详细输出 eg: -v true (default: false)
   --Thread value, -t value            线程 eg: -t 10 (default: 10)
   --help, -h                          show help (default: false)
   --version, -v                       print the version (default: false)
   ```

![image](https://user-images.githubusercontent.com/67625626/206980323-1d851182-8571-4418-8cad-3d7355036acf.png)

