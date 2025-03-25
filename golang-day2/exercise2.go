package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func getRating(id int, ratings chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Simulate delay
	rating := rand.Intn(5) + 1                                    // Random rating between 1 and 5
	ratings <- rating
}

func main() {
	rand.Seed(time.Now().UnixNano())
	numStudents := 200
	ratings := make(chan int, numStudents)
	var wg sync.WaitGroup

	for i := 0; i < numStudents; i++ {
		wg.Add(1)
		go getRating(i, ratings, &wg)
	}

	go func() {
		wg.Wait()
		close(ratings)
	}()

	totalRating := 0
	count := 0
	for rating := range ratings {
		totalRating += rating
		count++
	}

	averageRating := float64(totalRating) / float64(count)
	fmt.Printf("Average Rating: %.2f\n", averageRating)
}
