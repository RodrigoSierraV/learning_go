package main

import (
	"fmt"
	"sync"
)

var counter int

func increment(wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		counter++ // Race condition here
		fmt.Println("INC counter value:", counter)

	}
	wg.Done()
}

func decrement(wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		counter-- // Race condition here
		fmt.Println("DEC counter value:", counter)

	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go increment(&wg) // First goroutine
	go decrement(&wg) // Second goroutine

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}

/* EXPLANATION

The operation counter++ (increment) and counter-- (decrement) are not simple operations.
They consist of multiple steps:
	Reading the current value of counter.
	Modifying that value (either adding or subtracting).
	Writing the modified value back to memory.

If the two goroutines execute in such a way that they overlap in their operations (which 
is very likely since they run concurrently), we could have a scenario where:

    The incrementing goroutine reads the value of counter and gets, say, 5.
    Before it can update the value, the decrementing goroutine also reads the value of counter,
	sees 5, subtracts 1, and writes 4.
    Now, the incrementing goroutine finally writes its value (which was based on reading 5
	earlier), and overwrites the counter with 6, even though it should have been 4 + 1 = 5.

This race condition leads to unpredictable behavior. The final value of the counter may be
incorrect, and we cannot guarantee what it will be because it depends on the exact timing of
the goroutines. If both goroutines read and write to counter in a conflicting manner, the
result will be unreliable.
*/