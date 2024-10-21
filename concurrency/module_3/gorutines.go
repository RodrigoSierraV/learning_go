/* Write a Bubble Sort program in Go. The program
Write a program to sort an array of integers. The program should partition the array into 4 parts,
each of which is sorted by a different goroutine. Each partition should be of approximately equal
size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array. 

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of
the array should print the subarray that it will sort. When sorting is complete, the main
goroutine should print the entire sorted list.
*/
package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"os"
)

func main() {
	intSlice := make([]int, 0, 10)
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Type a sequence of up to 10 integers, space separated: '1 2 3'\n")
	enterString, _ := reader.ReadString('\n')
	stringSlice := strings.Split(strings.TrimSpace(enterString), " ")

	for _, v := range stringSlice {
		newInt, err := strconv.Atoi(v)

		if err != nil {
			continue
		}

		intSlice = append(intSlice, newInt)
	}

	size := int(math.Ceil(float64(len(intSlice) / 4)))
	remainder := len(intSlice) % 4
	var wg sync.WaitGroup
	init := 0

	for i := 0; i < 4; i++ {
		if i == 3 {
			size += remainder
		}
		wg.Add(1)
		go SortArray(intSlice[init:(init + size)], &wg)
		init += size
	}

	wg.Wait()
	BubbleSort(intSlice)
	fmt.Println(intSlice)
}

func SortArray(s []int, wg *sync.WaitGroup) {
	fmt.Println("SUBARRAY: ",s)
	BubbleSort(s)
	wg.Done()
}

func BubbleSort(s []int) {
	for i := 0; i < len(s); i++ {
		for j := 1; j < len(s); j++ {
			if s[j - 1] > s[j] {
				Swap(s, j - 1)
			}
		}
	}
}

func Swap(s []int, i int) {
	s[i], s[i + 1] = s[i + 1], s[i]
}
