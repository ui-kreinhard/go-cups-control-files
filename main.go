package main

import (
	"io/ioutil"
	"log"
	"github.com/ui-kreinhard/go-cups-control-files/controlFile"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Usage: ./go-cups-control-files PATH_TO_CONTROL_FILE")
	}
	
	fileName := os.Args[1]
	
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("error on reading file", err)
	}
	controlFile.ParseBytes(bytes).PrintContent()
}
