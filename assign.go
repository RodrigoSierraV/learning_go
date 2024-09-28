package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Write a Bubble Sort program in Go. The program
//should prompt the user to type in a sequence of up to 10 integers. The program
//should print the integers out on one line, in sorted order, from least to
//greatest. Use your favorite search tool to find a description of how the bubble
//sort algorithm works.
func main() {
	fmt.Println("enter your numbers:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	nums := readNumbers(line)
	BubbleSort(nums)
	fmt.Print(nums)
}

func readNumbers(numsAsStr string) []int {
	numsSplitted := strings.Split(numsAsStr, " ")
	nums := make([]int, 0, len(numsSplitted))
	for _, s := range numsSplitted {
		num, _ := strconv.Atoi(s)
		nums = append(nums, num)
	}
	return nums
}

// Swap
//
//A recurring operation in the bubble sort algorithm is
//the Swap operation which swaps the position of two adjacent elements in the slice.
//
//You should write a Swap() function which performs this operation.
//
//Your Swap() function should take two arguments, a slice of integers and an index value i which
//indicates a position in the slice.
//
//The Swap() function should return nothing, but it should swap
//the contents of the slice in position i with the contents in position i+1.
func Swap(stack []int, index int) {
	if index >= len(stack)-1 {
		return
	}
	if stack[index] > stack[index+1] {
		stack[index], stack[index+1] = stack[index+1], stack[index]
	}
}

// BubbleSort
//
//As part of this program, you should write a
//function called BubbleSort() which
//takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
//order.
func BubbleSort(stack []int) {
	for i := 0; i < len(stack)-1; i++ {
		for ii := 0; ii < len(stack)-1; ii++ {
			Swap(stack, ii)
		}
	}
}
