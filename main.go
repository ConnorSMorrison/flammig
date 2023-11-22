package main

import (
    "flammig/files"
    "os"
    // "fmt"
    "strings"
    "os/exec"
    "github.com/rivo/tview"
    "github.com/gdamore/tcell/v2"
    "path"
    "path/filepath"
)

func setFiles(fileList *tview.List, absolutePath string, folders []string, fls []string) ([]string, []string) {
    folders, fls = files.GetFiles(absolutePath)
    
    fileList.Clear()
    for _, folder := range folders {
        fileList.AddItem("[blue]" + folder, "", 0, nil)
    }
    for _, file := range fls {
        fileList.AddItem(file, "", 0, nil)
    }

    fileList.SetCurrentItem(0)

    return folders, fls
}

func main() {
    app := tview.NewApplication()

    absolutePath, _ := filepath.Abs(".")

    flex := tview.NewFlex()
    flex.SetBorder(true)
    flex.SetBorderPadding(1, 1, 1, 1)
    flex.SetTitle("Flammig File Manager")
    flex.SetDirection(tview.FlexRow)

    fileList := tview.NewList().SetWrapAround(true).SetSelectedBackgroundColor(tcell.ColorGray)
    folders, fls := files.GetFiles(".")
    flex.AddItem(fileList, 0, 1, true)

    folders, fls = setFiles(fileList, absolutePath, folders, fls)

    app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Key() {
        case tcell.KeyRune:
            switch event.Rune() {
            case 'h':
            // go back
            absolutePath = path.Join(absolutePath, "..")
            folders, fls = setFiles(fileList, absolutePath, folders, fls)
            case 'l':
            // go into a selected folder or xdg-open file
            if len(folders) + len(fls) != 0 {
                main, _ := fileList.GetItemText(fileList.GetCurrentItem())
                filePath := path.Join(absolutePath, main)
                strs := make([]string, 0)
                for _, w := range strings.Split(filePath, "/") {
                    strs = append(strs, strings.Replace(w, "[blue]", "", -1))
                }
                filePath = strings.Join(strs, "/")
                // fmt.Println(filePath)
                absolutePath = filePath
                fl, _ := os.Stat(filePath)
                mode := fl.Mode()
                if mode.IsDir() {
                    absolutePath = filePath
                    folders, fls = setFiles(fileList, absolutePath, folders, fls)
                } else {
                    cmd := exec.Command("gio", "open", filePath)
                    cmd.Run()
                    app.Stop()
                }
            }

            case 'j':
            // go down
            if fileList.GetCurrentItem() + 1 == len(fls) + len(folders) {
                fileList.SetCurrentItem(0)
            } else {
                fileList.SetCurrentItem(fileList.GetCurrentItem() + 1)
            }
            case 'k':
            // go up
            fileList.SetCurrentItem(fileList.GetCurrentItem() - 1)
        }
        }
        return event
    })

    if err := app.SetRoot(flex, true).Run(); err != nil {
        panic(err)
    }
}
