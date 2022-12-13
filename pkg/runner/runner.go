package runner

import (
	"akawaka/pkg/config"
	"akawaka/pkg/utils"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"math/rand"
	"sync"
	"time"
)

type WaitGroupTask struct {
	Key   int
	Value any
}

type Runner struct {
	options *config.Options
}

var lock sync.Mutex

func New(options *config.Options) error {
	// 验证文件目录路径是否正确
	err := utils.IsValid(options.DirPath)
	if err != nil {
		return err
	}

	// 初始化参数
	utils.Init(options)

	files, _ := utils.GetFileList(options.DirPath)
	options.Count = len(files)
	options.CurrentCount = 0
	fmt.Printf("\n[!] starting search...\n")
	defer fmt.Printf("\r\n[!] search finished...\r\n")

	err = excute(options, files)
	return nil
}

func excute(options *config.Options, files []string) error {

	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(options.Thread, func(wgTask any) {
		defer wg.Done()
		filePath := wgTask.(WaitGroupTask).Value.(string)
		//add: check target alive
		if !utils.IsDir(filePath) {
			if options.Is_Filename {
				utils.SearchFilename(filePath)
			} else {
				utils.Search(filePath, options.Verbose)
			}
		}
		outputProcess(options)
	})

	defer p.Release()
	for k, target := range files {
		wg.Add(1)
		_ = p.Invoke(WaitGroupTask{Value: target, Key: k})
	}
	wg.Wait()
	/*
		for _, filePath := range files {
			if utils.IsDir(filePath) {
				continue
			} else {
				if options.Is_Filename {
					utils.SearchFilename(filePath)
				} else {
					utils.Search(filePath)
				}
			}
			options.CurrentCount++
			fmt.Printf("\r%d/%d | %d%% ", options.CurrentCount, options.Count, options.CurrentCount*100/options.Count)
		}
	*/

	return nil
}

func outputProcess(options *config.Options) {
	lock.Lock()
	options.CurrentCount++
	fmt.Printf("\r%d/%d | %d%% ", options.CurrentCount, options.Count, options.CurrentCount*100/options.Count)
	lock.Unlock()
	RandSleep(1)
}

func RandSleep(millisencond int) {
	ms := millisencond + rand.Intn(millisencond)
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
