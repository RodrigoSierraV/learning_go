/* Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys
“name” and “address”, respectively. Your program should use Marshal() to create a JSON object
from the map, and then your program should print the JSON object. */

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	person_map := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Type your name: ")
    if scanner.Scan() {
        person_map["name"] = scanner.Text()
    }

	fmt.Printf("Type your address: ")
	if scanner.Scan() {
        person_map["address"] = scanner.Text()
    } 

	jsonString, err := json.Marshal(person_map)

	if err != nil {
		fmt.Printf("There was an error, Try again\n")
	} else {
		fmt.Printf("%s\n", jsonString)
	}
}
