package files

import (
    "log"
    "os"
    "path/filepath"
    "sort"
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

func getFoldersInPath(path string) ([]string, error) {
    absPath, err := filepath.Abs(path)
    if err != nil {
        return nil, err
    }

    folders := []string{}
    err = filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() && path != absPath {
            folders = append(folders, path)
        }
        return nil
    })

    if err != nil {
        return nil, err
    }

    return folders, nil
}

func separateFilesAndFolders(path string) ([]string, []string) {
    directory, err := os.ReadDir(path)

    if err != nil {
        log.Fatal(err)
    }
    var files []string
    var folders []string

    for _, file := range directory {
        if file.IsDir() {
            folders = append(folders, file.Name())
        } else {
            files = append(files, file.Name())
        }
    }
    sort.Strings(files)
    sort.Strings(folders)
    return files, folders
}

func GetFiles(path string) []string {
    return addAll(separateFilesAndFolders(path))
}
