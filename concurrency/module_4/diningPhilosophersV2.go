package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

const numberOfPhilosophers = 5

type Philosopher struct {
	id        int
	leftChop  *sync.Mutex
	rightChop *sync.Mutex
}

func (p Philosopher) pickUpChopsticks() {
	// Randomize which chopstick to pick up first
	if rand.IntN(2) == 0 {
		p.leftChop.Lock()
		p.rightChop.Lock()
	} else {
		p.rightChop.Lock()
		p.leftChop.Lock()
	}
}

func (p Philosopher) putDownChopsticks() {
	p.leftChop.Unlock()
	p.rightChop.Unlock()
}

func (p Philosopher) dine(wg *sync.WaitGroup, permission chan int) {
	for i := 0; i < 3; i++ {
		permission <- 1
		p.pickUpChopsticks()

		fmt.Printf("starting to eat %d\n", p.id)
		time.Sleep(500 * time.Millisecond) // Simulate eating
		fmt.Printf("finishing eating %d\n", p.id)

		p.putDownChopsticks()
		<- permission
	}
	wg.Done()
}


func main() {
	var wg sync.WaitGroup

	// Initialize chopsticks
	chopsticks := make([]*sync.Mutex, numberOfPhilosophers)
	for i := 0; i < numberOfPhilosophers; i++ {
		chopsticks[i] = &sync.Mutex{}
	}

	// Initialize a permission channel of size equal to max philosophers allowed to eat
	permission := make(chan int, 2)
	
	philosophers := make([]Philosopher, numberOfPhilosophers)
	for i := 0; i < numberOfPhilosophers; i++ {
		philosophers[i] = Philosopher{
			id:        i + 1,
			leftChop:  chopsticks[i],
			rightChop: chopsticks[(i+1)%numberOfPhilosophers],
		}
	}

	for _, philosopher := range philosophers {
		wg.Add(1)
		go philosopher.dine(&wg, permission)
	}

	wg.Wait()
}
