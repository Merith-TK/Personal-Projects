package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/h2non/filetype"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/reujab/wallpaper"
)

// Set empty variable
var IMAGE = ""

//Main Process/window
func main() {
	fmt.Println("Wallpaper v2.1")
	var textEdit *walk.TextEdit
	MainWindow{
		Title:   "Wallpaper Tool",
		MinSize: Size{320, 240},
		Layout:  VBox{},

		// Listen for file being drag'n'dropped
		OnDropFiles: func(files []string) {
			textEdit.SetText(strings.Join(files, "\r\n"))
			//set 'wallpaper' to the filepath of said image
			IMAGE = strings.Join(files, "\r\n")

			//check what type of file the provided file is
			buf, err := ioutil.ReadFile(IMAGE)
			if err != nil {
				panic(err)
			}

			// See if file is a image
			if filetype.IsImage(buf) {
				textEdit.SetText(strings.Join(files, "\r\n"))
				wallpaper.SetFromFile(IMAGE)
			} else {
				fmt.Println("The File is not a Image")
				textEdit.SetText("The File is not a Image")
			}
		},
		Children: []Widget{
			TextEdit{
				AssignTo: &textEdit,
				ReadOnly: true,
				Text:     "Drag your desired wallpaper here.",
			},
		},
	}.Run()
}
