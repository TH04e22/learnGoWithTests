package main

import (
	"os"
	"time"

	"learnGoWithTests/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
