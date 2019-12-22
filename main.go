package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

import (
	"flag"
	"fmt"
)

// import (
// 	"fmt"
// 	"os"
// )

type MyMainWindow struct {
	*walk.MainWindow
}

func getArgs() []string {

	flag.Parse()
	args := flag.Args()
	return args
}
	

func main() {

	args := getArgs()
	var strs string
	for str := range args {
		strs = append(strs, str)
	}
	
	mw := &MyMainWindow{}
	var inTE, outTE *walk.TextEdit
	var results *walk.ListBox
	
	MW := MainWindow{
		AssignTo: &mw.MainWindow,
		Title : "Customer Info",
		MinSize: Size{150,200},
		Size: Size{300,400},
		Layout : VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE, ReadOnly: true},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			ListBox{
				AssignTo: &results,
                Row: 1,
			},
			PushButton{
				Text:"PB0",
			},
			PushButton{
				Text: "PB1",
			},
		},
	}

	MW.Run()
}
	
					
