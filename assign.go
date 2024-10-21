package main

import (
	"fmt"
	"sync"
	"time"
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	host  *Host
	num   int
	left  *Chopstick
	right *Chopstick
}

func (p *Philosopher) Eat() {
	if p.num%2 == 0 {
		for i := 0; i < 3; {
			if p.host.allowToEat() {
				p.left.Lock()
				p.right.Lock()
				fmt.Println("starting to eat", p.num)
				time.Sleep(100 * time.Millisecond)
				fmt.Println("finishing eating", p.num)
				p.left.Unlock()
				p.right.Unlock()
				p.host.finishEating()
				i++
			}
		}
	} else {
		for i := 0; i < 3; {
			if p.host.allowToEat() {
				p.right.Lock()
				p.left.Lock()
				fmt.Println("starting to eat", p.num)
				time.Sleep(100 * time.Millisecond)
				fmt.Println("finishing eating", p.num)
				p.right.Unlock()
				p.left.Unlock()
				p.host.finishEating()
				i++
			}
		}
	}
	wg.Done()
}

type Host struct {
	eatingCount int
	mut         sync.Mutex
}

func (h *Host) allowToEat() bool {
	if h.eatingCount < 2 {
		h.mut.Lock()
		h.eatingCount++
		h.mut.Unlock()
		return true
	}
	return false
}

func (h *Host) finishEating() {
	h.mut.Lock()
	h.eatingCount--
	h.mut.Unlock()
}

var wg sync.WaitGroup

func main() {
	host := Host{eatingCount: 0}
	chopsticks := make([]Chopstick, 5)
	philos := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philosopher{host: &host, num: i + 1, left: &chopsticks[i], right: &chopsticks[(i+1)%5]}
	}

	wg.Add(len(philos))
	for i := 0; i < 5; i++ {
		go philos[i].Eat()
	}
	wg.Wait()
}
