package main

import (
	"fmt"
	"log"
	"os"
)

func ListDir(folderName string) {
	entries, err := os.ReadDir(folderName)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fileInfo, err := os.Stat(folderName + "/" + e.Name())
		if err == nil {
			if fileInfo.IsDir() {
				ListDir(folderName + "/" + e.Name())
			} else {
				fmt.Println(folderName + "/" + e.Name())
			}
		}
	}
}

func main() {
	ListDir("./dev")
}
