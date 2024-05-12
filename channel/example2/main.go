package main

import (
	"fmt"
	"go-concurrency-demo/internal/user"
	"time"
)

func main() {

	now := time.Now()

	var users []user.User

	ch := getRandomUsers()

	for i := 0; i < 10; i++ {
		users = append(users, <-ch)
	}

	fmt.Println("time elapsed: ", time.Since(now))
}

func getRandomUsers() <-chan user.User {

	ch := make(chan user.User)

	go func() {
		for i := 0; i < 10; i++ {
			go func() {
				ch <- user.GetRandomUser()
			}()
		}
	}()

	return ch
}
