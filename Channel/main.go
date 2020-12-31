package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println(<-c) //이거 한줄 한 번하면 하나만 됨.
	fmt.Println(<-c)
	//fmt.Println(<-c) 넘게 받으면 deadlock 상태가 됨.
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true
}
