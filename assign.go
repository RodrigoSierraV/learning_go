package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a string:")

	line, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}

	line = strings.TrimSpace(line)
	line = strings.ToLower(line)

	if strings.HasPrefix(line, "i") && strings.Contains(line, "a") && strings.HasSuffix(line, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
