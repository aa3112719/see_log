package main

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io/ioutil"
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

func (a *App) FileInfo(name string) string {
	fileInfo, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", a.selection, name))
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Sprintf("%s", string(fileInfo))
}
