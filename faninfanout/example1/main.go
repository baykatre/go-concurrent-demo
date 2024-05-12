package main

import (
	"fmt"
	"go-concurrency-demo/internal/user"
	"time"
)

func main() {

	now := time.Now()

	var users []user.User

	ch := fanIn(getRandomUser(), getRandomUser())

	for i := 0; i < 2; i++ {
		users = append(users, <-ch)
	}

	fmt.Println(len(users))
	fmt.Printf("time elapsed: %v\n", time.Since(now))
}

func fanIn(ch1, ch2 <-chan user.User) <-chan user.User {

	ch := make(chan user.User)

	go func() {
		for {
			ch <- <-ch1
		}
	}()

	go func() {
		for {
			ch <- <-ch2
		}
	}()

	return ch
}

func getRandomUser() <-chan user.User {

	ch := make(chan user.User)

	go func() {
		ch <- user.GetRandomUser()
	}()

	return ch
}
