package main

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io"
	"io/ioutil"
	"os"
)

// App struct
type App struct {
	ctx       context.Context
	selection string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here

}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) []string {
	selection, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select Dir",
	})
	if err != nil {
		fmt.Println(err)
	}
	files, _ := ioutil.ReadDir(selection)
	var outputlist []string
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("目录:", file.Name())
		} else {
			outputlist = append(outputlist, file.Name())
		}
	}
	a.selection = selection
	return outputlist
}

func (a *App) FileInfo() string {
	selection, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select Log File",
	})
	if err != nil {
		fmt.Println(err)
		return ""
	}
	a.selection = selection
	return selection
}

func revers(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

type FileStrStruct struct {
	selection string   `json:"selection"`
	data      []string `json:"data"`
}

func (a *App) GetFileStr() []string {
	if a.selection == "" {
		return make([]string, 1)
	}
	//打开文件
	file, err := os.Open(a.selection)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	buff := make([]byte, 0, 4096)
	char := make([]byte, 1)
	var outStr []string

	// 查询文件大小
	stat, _ := file.Stat()
	filesize := stat.Size()

	var cursor int64 = 0
	cnt := 0
	for {
		cursor -= 1
		_, _ = file.Seek(cursor, io.SeekEnd)
		_, err = file.Read(char)
		if err != nil {
			panic(err)
		}

		if char[0] == '\n' || char[0] == '\r' {
			if len(buff) > 0 {
				revers(buff)
				// 读取到的行
				outStr = append(outStr, string(buff))
				cnt++
				if cnt == 100 {
					// 超过数量退出
					break
				}
			}
			buff = buff[:0]
		} else {
			buff = append(buff, char[0])
		}
		if cursor == -filesize {
			break
		}
	}

	return outStr
}
