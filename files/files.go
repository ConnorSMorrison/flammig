package files

import (
    "log"
    "os"
    "sort"
    "strings"
)

func contains(sli []string, element string) bool {
    for _, el := range sli {
        if el == element { return true }
    }
    return false
}

func addAll(sli1 []string, sli2 []string) []string {
    var newSli []string
    for _, el := range sli1 {
        newSli = append(newSli, el)
    }
    for _, el := range sli2 {
        newSli = append(newSli, el)
    }
    return newSli
}

func separateFilesAndFolders(path string, showHidden bool) ([]string, []string) {
    directory, err := os.ReadDir(path)

    if err != nil {
        log.Fatal(err)
    }
    var files []string
    var folders []string

    for _, file := range directory {
        if strings.HasPrefix(file.Name(), ".") && !showHidden { continue }

        if file.IsDir() {
            folders = append(folders, file.Name())
        } else {
            files = append(files, file.Name())
        }
    }

    sort.Strings(files)
    sort.Strings(folders)
    return folders, files
}

func GetFiles(path string, showHidden bool) ([]string, []string) {
    return separateFilesAndFolders(path, showHidden)
}
