package main

import (
	walkMain "github.com/lxn/walk"
	walk "github.com/lxn/walk/declarative"

	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/ncruces/zenity"
)

var (
	globalFiles []string
)

func writeCrashLog(err string) {
	currDir, _ := os.Getwd()
	path := path.Join(currDir, fmt.Sprintf("costCalculatorCrash%v.txt", time.Now().Unix()))
	content := fmt.Sprintf("Crash Log: %v\n%v", time.Now().Format(time.RFC850), err)
	os.WriteFile(path, []byte(content), 0644)
}

func panicAndLog(err string) {
	writeCrashLog(err)
	panic(err)
}

func fatalDialog(err string) {
	zenity.Error(fmt.Sprintf("%v", err), zenity.Title("Fatal Error"))
}

func fatalDialogAndLog(err string) {
	fatalDialog(err)
	panicAndLog(err)
}

// scanFolder returns a slice of all valid file paths in the given folder
func scanFolder(path string) ([]string, error) {
	var validPaths []string
	pattern := regexp.MustCompile(`[ _A-Za-z0-9]+_[0-9]*[.0-9]+[a-z]{3}`)
	var walkFunc = func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if pattern.MatchString(info.Name()) { // add only if valid file name
			//fmt.Printf("Matched: %s\n", info.Name())

			if !info.IsDir() { // do not add directory
				validPaths = append(validPaths, path)
			}
		}
		return nil
	}
	err := filepath.Walk(path, walkFunc)
	if err != nil {
		return nil, err
	}
	return validPaths, nil
}

// calculate total of all names given in the array
func costTotal(paths *[]string, textbox *walkMain.TextEdit) float64 {
	total := 0.0
	textbox.SetText("")
	for _, filename := range *paths {
		cost,err := calcCost(filename, "_")
		if err != nil {
			fatalDialogAndLog(err.Error())
		}
		total += cost
		_, name := filepath.Split(filename)
		textbox.AppendText(fmt.Sprintf("%s **Cost: %.2f**\r\n", name, cost))
	}
	return total
}

// calcCost returns the cost of a single name
func calcCost(filename string, sep string) (float64,error) {
	filename = strings.TrimRight(filename, filepath.Ext(filename))
	split := strings.Split(filename, sep)
	end := split[len(split)-1]
	cost, err := strconv.ParseFloat(end, 64)
	if err != nil {
		return 0.0, fmt.Errorf("calcCost>strconv.ParseFloat: %v", err)
	}
	return cost,nil
}

func openFolderDialog(textBox *walkMain.TextEdit) {
	selectedFolder, err := zenity.SelectFile(
		zenity.Title("Select Folder"),
		zenity.Directory())
	if err == zenity.ErrCanceled { // cancel then return
		return
	} else if err != nil { // any other error crash
		panicAndLog("openFolderDialog>zenitySelectFile" + err.Error())
	}
	err = zenity.Question(fmt.Sprintf("Selected Folder: %s", selectedFolder), zenity.Title("Confirm"))
	if err == zenity.ErrCanceled { // also cancel if wrong folder selected
		return
	} else if err != nil {
		panicAndLog("openFolderDialog>zenityQuestion" + err.Error())
	}

	files,err := scanFolder(selectedFolder)
	if err != nil {
		panicAndLog(fmt.Sprint("openFolderDialog>scanFolder: %v", err))
	}

	textBox.SetText("")
	for _, file := range files {
		if !slices.Contains(globalFiles, file) {
			globalFiles = append(globalFiles, file) // add any new files
		}
	}

	for _, file := range globalFiles {
		textBox.AppendText(file + "\r\n")
	}
}

func main() {
	var inputFileBox *walkMain.TextEdit
	defer func() {
		if err := recover(); err != nil {
			panicAndLog(fmt.Sprintf("main>panic: %v", err))
		}
	}()

	walk.MainWindow{
		Title:   "Cost Calculator",
		MinSize: walk.Size{Width: 600, Height: 400},
		Size:    walk.Size{Width: 640, Height: 360},
		Layout:  walk.VBox{},
		Children: []walk.Widget{walk.TextLabel{Text: "Input Files"},
			walk.HSplitter{
				Children: []walk.Widget{
					walk.TextEdit{AssignTo: &inputFileBox, ReadOnly: true, Text: "Use add folder button to add files", Font: walk.Font{PointSize: 10}},
				},
			},
			walk.HSplitter{Children: []walk.Widget{
				walk.PushButton{
					Text: "Add Folder",
					OnClicked: func() {
						go openFolderDialog(inputFileBox)
					},
				},
				walk.PushButton{
					Text: "Calculate",
					OnClicked: func() {
						cost := costTotal(&globalFiles, inputFileBox)
						zenity.Info(fmt.Sprintf("Total Cost: %.2f", cost), zenity.Title("Total Cost"))
					},
				}, walk.PushButton{
					Text: "Clear files",
					OnClicked: func() {
						inputFileBox.SetText("Use select folder button to add files")
						globalFiles = []string{}
					}},
			}},
		},
	}.Run()
}
