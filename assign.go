package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	namesList := make([]name, 0)
	fmt.Print("Provide name of file: ")
	reader := bufio.NewReader(os.Stdin)
	fileNameBytes, _, _ := reader.ReadLine()
	fileName := string(fileNameBytes)
	fileContentBytes, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Sorry, some issue happened. Couldn't find file")
	} else {
		fileContent := string(fileContentBytes)
		contentLines := strings.Split(fileContent, "\n")
		for _, cLine := range contentLines {
			if len(cLine) > 0 {
				cLineSplit := strings.Split(cLine, " ")
				nameObj := createName(cLineSplit[0], cLineSplit[1])
				namesList = append(namesList, nameObj)
			}
		}
		for _, n := range namesList {
			fmt.Println(n.fname, n.lname)
		}
	}
}

func createName(fname string, lname string) name {
	newName := name{fname: fname, lname: lname}
	if len(fname) > 20 {
		newName.fname = fname[:21]
	}
	if len(lname) > 20 {
		newName.lname = lname[:21]
	}
	return newName
}
