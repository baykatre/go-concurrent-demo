package main

import (
	"fmt"
	"go-concurrency-demo/internal/user"
	"sync"
	"time"
)

func main() {

	now := time.Now()

	var users []user.User

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			users = append(users, getRandomUser(&wg))
		}()
	}

	wg.Wait()
	fmt.Println(len(users))
	fmt.Println(users)
	fmt.Printf("time elapsed: %v\n", time.Since(now))
}

func getRandomUser(wg *sync.WaitGroup) user.User {

	defer wg.Done()

	randomUser := user.GetRandomUser()
	return randomUser
}
