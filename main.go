package main

import (
	"younes.dev/go/fileutils"
)

func main() {
	print(fileutils.ReadTextFile("data/file.txt"))
}