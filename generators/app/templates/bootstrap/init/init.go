package main

import (
    "github.com/lambda-platform/lambda/DB/generator/initModel"
    "os"
)

func main() {
    dir, _ := os.Getwd()
    initModel.ModelInit(dir, "<%= projectName %>")
}
