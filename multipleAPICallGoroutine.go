package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// Function to fetch data from a given URL
func fetchData(url string, wg *sync.WaitGroup, resultChan chan<- string) {
	defer wg.Done() // Decrement the counter when the function completes

	// Perform the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		resultChan <- fmt.Sprintf("Error fetching %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resultChan <- fmt.Sprintf("Error reading response body from %s: %v", url, err)
		return
	}

	// Send the result to the result channel
	resultChan <- fmt.Sprintf("Response from %s:\n%s", url, body)
}

func main() {
	// Define the URLs to fetch
	urls := []string{
		"https://dummy.restapiexample.com/api/v1/employee/3",
		"https://dummy.restapiexample.com/api/v1/employee/1",
		"https://dummy.restapiexample.com/api/v1/employee/2",
	}

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Create a channel to collect the results
	resultChan := make(chan string, len(urls))

	// Start a goroutine for each URL
	for _, url := range urls {
		wg.Add(1)
		go fetchData(url, &wg, resultChan)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Close the result channel
	close(resultChan)

	// Print all results
	for result := range resultChan {
		fmt.Println(result)
	}
}
