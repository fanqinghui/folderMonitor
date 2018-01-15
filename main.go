package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"strings"
	"time"
)

var FolerPath = "C:\\Users\\Administrator\\Desktop\\考核"
var FileType = ".zip"
var TimeFormat = "2018-01-01 00:00:00"

func main() {

	fileChan := make(chan int, 1)
	ticker := time.NewTicker(time.Minute * 1)

	go func() {
		for {
			//now := time.Now()
			// 计算下一个零点
			//			next := now.Add(time.Hour * 24)
			//			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			//			time.NewTimer(next.Sub(now))
			select {
			case <-ticker.C:
				delZipFile()
			}
		}
	}()
	//阻塞主线程
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
			//			fileTime := file.ModTime().Unix()
			//			fmt.Printf("%s", fileTime)
			fmt.Println(file.Name())
			os.Remove(fileName)
		}
		return nil

	})
}
