/* Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop. Before entering the loop, the program should create an empty
integer slice of size (length) 3. During each pass through the loop, the program prompts the user to
enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice,
and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any
number of integers which the user decides to enter. The program should only quit (exiting the loop)
when the user enters the character ‘X’ instead of an integer. */

package main

import (
	"fmt"
	"strconv"
	"sort"
)

func main() {

	int_slice := make([]int, 0, 3)
	var typed string
	
	for ; true; {
		fmt.Printf("Type an integer or X if you want to stop the program\n")
		fmt.Scan(&typed)	
		new_int, err := strconv.Atoi(typed)

		if typed == "X" {
			break
		}

		if err != nil  {
			fmt.Printf("Non-integer typed\n")
			continue
		}

		int_slice = append(int_slice, new_int)
		sort_and_print(int_slice)
	}
	sort_and_print(int_slice)
	fmt.Printf("End of the program\n")
}

func sort_and_print (int_slice []int) {
	sort.Slice(int_slice, func(i, j int) bool {
		return int_slice[i] < int_slice[j]
	})
	fmt.Println("Your ordered Slice:", int_slice)
}