package main

import (
	"fmt"
	"os"
	"time"
)

func main(){

	var logType string
	var text string
	var output string

	theTime := time.Now().String()

	if len(os.Args) < 2 {
		fmt.Println("./meow-log [info/warning/error] [TEXT] [OUTPUT FILE]")
		return
	} else {
		logType = os.Args[1]
		text = os.Args[2]
		output = os.Args[3]
	}

	file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("could not open file")
		return
	}

	defer file.Close()

	writeContent := "[" + logType + "]	" + "[ " + theTime + " ]		" + text + "\n"

	_, err2 := file.WriteString(writeContent)
	if err2 != nil {
		fmt.Println("could not write into ", output)
		return
	}
}
