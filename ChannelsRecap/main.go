package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [5]string{"nico", "flynn", "dal", "japanguy", "larry"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println("Received this message: " + <-c)
	}
	fmt.Println("Waiting for messages")
	// resultOne := <-c
	// resultTwo := <-c
	// fmt.Println("Received this message: " + resultOne)
	// fmt.Println("Received this message: " + resultTwo)
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second)
	c <- person + " is sexy"
}
