package main

import (
    "tfvwr/files"
    "fmt"
)

func main() {
    files := files.GetFiles(".")
    fmt.Println(files)
}
