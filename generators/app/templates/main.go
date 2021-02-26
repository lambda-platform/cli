package main

import (
    "<%= projectName %>/bootstrap"
)

func main() {
    lambda := bootstrap.Set()
    lambda.Start()
}
