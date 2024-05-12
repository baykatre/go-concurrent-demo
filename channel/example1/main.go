package main

import (
	"fmt"
	"go-concurrency-demo/internal/user"
	"time"
)

func main() {

	now := time.Now()

	var users []user.User

	ch := make(chan user.User)
	getRandomUsers(ch)

	for i := 0; i < 10; i++ {
		users = append(users, <-ch)
	}

	fmt.Println("time elapsed: ", time.Since(now))
}

func getRandomUsers(ch chan user.User) {

	for i := 0; i < 10; i++ {
		go func() {
			ch <- user.GetRandomUser()
		}()
	}
}
