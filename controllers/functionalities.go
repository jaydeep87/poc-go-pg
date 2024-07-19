package controllers

import (
	"fmt"
	"sync"
	"net/http"
	topics "github.com/jaydeep87/poc-go-pg/topics"
	"github.com/gin-gonic/gin"
)

func RunGoroutinesWithChan(c *gin.Context) {
	limit :=100	
	primeNoChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go topics.GeneratePrime(primeNoChan, limit, &wg)
	// go topics.PrintPrime(primeNoChan, &wg)
	
	// Receive results from the results channel
	var allPrimes []int
	for result := range primeNoChan {
		allPrimes = append(allPrimes, result)
		fmt.Println("Received result:", result)
	}
	fmt.Println(allPrimes);
	
	c.JSON(http.StatusOK, gin.H{
		"sc":  http.StatusOK,
		"sm": "success",
		"primes": allPrimes,
	})
	return
	wg.Wait()

}
func GoRoutineChanCommunication(c *gin.Context) {
	limit :=100	
	primeNoChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go topics.GeneratePrime(primeNoChan, limit, &wg)
	go topics.PrintPrime(primeNoChan, &wg)
	
	// Receive results from the results channel
	wg.Wait()
	
	c.JSON(http.StatusOK, gin.H{
		"sc":  http.StatusOK,
		"sm": "prime no generated",
	})
	// return

}