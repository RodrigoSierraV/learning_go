package main

import (
  "fmt";
//   "strings";
//   "strconv";
//   "sort";
)
func main() {
	s := make([]int, 0, 3)
	s = append(s, 100)
	fmt.Println(s, len(s), cap(s))
  }