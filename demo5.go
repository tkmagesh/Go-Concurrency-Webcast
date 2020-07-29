package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}
	for {
		c <- rand.Int()
		time.Sleep(time.Millisecond * 5000)
	}
}

type Worker struct {
	id int
}

func (w *Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
	}
}

/*
// what happens if we have more data coming in than we can handle

for {
    data := <-c
    fmt.Printf("worker %d got %d\n", w.id, data)
    time.Sleep(time.Millisecond * 500)
}

c := make(chan int, 100)


//test the growing size of the channel
for {
    c <- rand.Int()
    fmt.Println(len(c))
    time.Sleep(time.Millisecond * 50)
} */
