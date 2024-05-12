package main

import (
	"fmt"
	"go-concurrency-demo/internal/user"
	"time"
)

func main() {

	c := getRandomUsers()

	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(time.Second * 1):
			fmt.Println("timeout")
			return
		}
	}
}

func getRandomUsers() <-chan user.User {

	c := make(chan user.User)

	go func() {
		for i := 0; i < 10; i++ {
			go func() {
				c <- user.GetRandomUser()
			}()
		}
	}()

	return c
}
