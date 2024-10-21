package main

import (
	"fmt"
	// "slices"
	"sync"
)

type chopMU struct {
	mu sync.Mutex
	rightChopStick int
	leftChopStick int
}
var phils = []int{1, 2, 3, 4, 5}
var chopSticks = []int{0,0,0,0,0}

func eat(phil int, wg *sync.WaitGroup, permission chan int) []int {

	var c1, c2 int
	for i := 0; i < 3; i++ {
		permission <- 1
		cs.mu.Lock()
		fmt.Println("starting to eat ", phil)
		c1 = cs.chopSticks[phil - 1]
		c2 = cs.chopSticks[(phil)%5]
		cs.mu.Unlock()
		fmt.Println("finishing eating ", phil)
		<- permission
	}
	wg.Done()
	return []int{c1, c2}
}


func main() {
	var wg sync.WaitGroup
	permission := make(chan int, 2)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go eat(phils[i], &wg, permission)
	}

	wg.Wait()

}
