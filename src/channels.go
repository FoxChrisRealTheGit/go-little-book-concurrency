package main


import (
	"fmt"
	"time"
	"math/rand"
)

func main(){
	//give channel a length to add overflow into queue
	c := make(chan int, 100)
	for i := 0; i<5; i++{
		worker := &Worker{id: i}
		go worker.process(c)
	}

	for{
		c <- rand.Int()
		fmt.Println(len(c))
		time.Sleep(time.Millisecond * 50)
	}
}

type Worker struct{
	id int
}

func (w *Worker) process(c chan int){
	for{
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)

		//to simulate more data than can be handled
		time.Sleep(time.Millisecond * 500)
	}
}