package main

import (
    "lambda/bootstrap"
)

func main() {
    lambda := bootstrap.Set()
    lambda.Start()
}
