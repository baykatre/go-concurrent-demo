package main

import (
	"go-concurrency-demo/internal/user"
)

func main() {

	c := fanIn(getRandomUsers(), getRandomUsers())

	for i := 0; i < 10; i++ {
		println(<-c)
	}
}

func fanIn(ch1, ch2 <-chan user.User) chan user.User {

	c := make(chan user.User)

	go func() {
		for {
			select {
			case u := <-ch1:
				c <- u
			case s := <-ch2:
				c <- s
			}
		}
	}()

	return c
}

func getRandomUsers() <-chan user.User {

	c := make(chan user.User)

	go func() {
		for i := 0; i < 10; i++ {
			c <- user.GetRandomUser()
		}
	}()

	return c
}
