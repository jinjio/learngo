package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [4]string{"nico", "jio", "dal", "larry"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println("wating for", i)
		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is sexy"
}
