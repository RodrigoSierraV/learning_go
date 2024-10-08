package main

import "fmt"
import "sync"

var Wait sync.WaitGroup
var Counter = 0

func Add(id int, value int) {
	for i := 0; i < value; i++ {
		Counter += 1
		fmt.Printf("id #%d = %d\n", id, Counter)
		if Counter >= id*2 {
			break
		}
	}
}
func Add2(id int, value int) {
	for i := 0; i < value; i++ {
		Counter += 1
		fmt.Printf("id #%d = %d\n", id, Counter)
		if Counter >= id*2 {
			break
		}
	}
	Wait.Done()
}

func main() {
	fmt.Println("Normal Call")
	Counter = 0
	for i := 1; i <= 10; i++ {
		Add(i, i)
	}
	fmt.Printf("Final counter : %d\n\n", Counter)

	Counter = 0
	fmt.Println("Concurrency Call")
	for i := 1; i <= 10; i++ {
		Wait.Add(1)
		go Add2(i, i)
	}
	Wait.Wait()
	fmt.Printf("Final counter : %d\n", Counter)
}

/**
Race condition is when multiple threads access the same resource(s) and one or more of the threads modify / change the value of the resource(s)

In this example, the normal call (sequential) :
- the sequence is known
- The state of each steps is known (e.g. the value of Counter always increased by 1)
- the result can be predict and always the same

While the concurrency call :
- The sequence is unknown
- the state in each steps is unknown (e.g. the value of the Counter is not always increasing by 1)
- The result is unknown and unpredictable
**/
