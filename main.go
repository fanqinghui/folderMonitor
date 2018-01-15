package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"strings"
	"time"
)

var FolerPath = "//data//research"
var FileType = ".zip"
var TimeFormat = "2018-01-01 00:00:00"

func main() {

	fileChan := make(chan int, 1)
	ticker := time.NewTicker(time.Hour * 168) //每个周清理一次

	go func() {
		for {
			//now := time.Now()
			// 计算下一个零点
			//next := now.Add(time.Hour * 24)
			//next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			//time.NewTimer(next.Sub(now))
			select { //就是监听 IO 操作,case 语句里面必须是一个IO操作（面向channel的Io操作）
			case <-ticker.C:
				delZipFile()
			}
		}
	}()
	<-fileChan
}

func delZipFile() {
	fmt.Println("定时清理数据:", time.Now().Format(TimeFormat))
	_, err := ioutil.ReadDir(FolerPath)
	if err != nil {
		fmt.Println(err)
	}
	filepath.Walk(FolerPath, func(fileName string, file os.FileInfo, err error) error {
		if !file.IsDir() && strings.HasSuffix(file.Name(), FileType) {
			fmt.Println(file.Name())
			os.Remove(fileName)
		}
		return nil
	})
}
