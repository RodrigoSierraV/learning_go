// Write a program which prompts the user to enter a string.
// The program searches through the entered string for the characters ‘i’,
// ‘a’, and ‘n’. The program should print “Found!” if the entered string
// starts with the character ‘i’, ends with the character ‘n’, and
// contains the character ‘a’. The program should print “Not Found!”
// otherwise. The program should not be case-sensitive, so it does not
// matter if the characters are upper-case or lower-case.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
   )

func main() {
	var enter_string string

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a string:\n")
	enter_string, _ = reader.ReadString('\n')
	enter_string = strings.ToLower(strings.TrimSpace(enter_string))

	if (strings.HasPrefix(enter_string, "i") && strings.HasSuffix(enter_string, "n") && strings.Contains(enter_string, "a")) {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}

}
