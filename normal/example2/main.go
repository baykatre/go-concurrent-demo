package main

import (
	"fmt"
	"go-concurrency-demo/internal/user"
	"time"
)

func main() {

	now := time.Now()

	var users []user.User

	for i := 0; i < 10; i++ {
		users = append(users, user.GetRandomUser())
	}

	fmt.Printf("Users generated in %s", time.Since(now))
}
