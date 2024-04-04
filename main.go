package main

import (
	"fmt"
	"os"

	"younes.dev/go/fileutils"
)

func main() {

	rootDir, _ := os.Getwd()

	c, err :=fileutils.ReadTextFile(rootDir + "/data/file.txt")
	if err == nil {
		fmt.Print(c)
	}else{
		fmt.Printf("AHHAHAHA %v", err)
		// panic("AHHAHAHA %v", err)
	}
}