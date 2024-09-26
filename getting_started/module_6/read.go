/* Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names. Each line of the text
file has a first name and a last name, in that order, separated by a single space on the line.
Your program will define a name struct which has two fields, fname for the first name, and lname
for the last name. Each field will be a string of size 20 (characters). Your program should
prompt the user for the name of the text file. Your program will successively read each line
of the text file and create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read from the file,
your program will have a slice containing one struct for each line in the file. After reading
all lines from the file, your program should iterate through your slice of structs and print
the first and last names found in each struct. */

package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)


func main() {
	var file_name string

	type name struct {
		fname string
		lname string
	}

	slice_name := make([]name, 0)
	fmt.Printf("Type the file name: ")
	fmt.Scan(&file_name)

	// Open the file
    file, err := os.Open(file_name)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    // Create a buffered reader
    reader := bufio.NewReader(file)

    // Read and print lines
    for {
		var temp_struct name
        line, err := reader.ReadString('\n')
        if err != nil {
            break
        }
		stringSlice := strings.Split(line, " ")
		temp_struct.fname = stringSlice[0]
		temp_struct.lname = stringSlice[1]

		slice_name = append(slice_name, temp_struct)
    }

	for _, v := range slice_name {
		fmt.Printf("First name: %s, Last name: %s", v.fname, v.lname)
	}
}
