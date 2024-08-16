package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"sync"
	"encoding/json"
	"bytes"
)

type PostPayload struct {
	Name string `json:"name"`
	Age    int    `json:"age"`
	Address string `json:"address"`
}
func GetApi(url string, wg *sync.WaitGroup, resultChan chan<- string) {
	 defer wg.Done()

	 resp, err := http.Get(url)
	 if err!=nil {
		resultChan <- fmt.Sprintf("error got from %s : %v", url, err)
		return
	 }

	defer resp.Body.Close()
	 // res body

	 body, err := ioutil.ReadAll(resp.Body)

	 if err!= nil{
		resultChan <- fmt.Sprintf("error in bodyparse for %s:%v", url, err)
		return
	 }

	 resultChan <- fmt.Sprintf("response from url %s : %s", url, body)

}

func PostApi(url string, payload PostPayload, wg *sync.WaitGroup,resultChan chan <- string){
	defer wg.Done() // Decrement the counter when the function completes

	// Serialize payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		resultChan <- fmt.Sprintf("Error marshaling payload for %s: %v", url, err)
		return
	}

	// Create a POST request with the JSON payload
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		resultChan <- fmt.Sprintf("Error posting to %s: %v", url, err)
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

func main(){
	apis:= []string{
		"https://dummy.restapiexample.com/api/v1/employee/3",
		"https://dummy.restapiexample.com/api/v1/employee/1",
		"https://dummy.restapiexample.com/api/v1/employee/2",
	}
	postApis:= []string{
		"https://dummy.restapiexample.com/api/v1/create",
		"https://dummy.restapiexample.com/api/v1/create",
		"https://dummy.restapiexample.com/api/v1/create",
	}

	// Define payloads for each POST request
	payloads := []PostPayload{
		{Name: "jaydeep", Age: 25, Address: "RR"},
		{Name: "yug", Age: 10, Address: "RR"},
		{Name: "yahvi", Age: 15, Address: "RR"},
	}

	resultChan := make(chan string, len(apis)*2)

	var wg sync.WaitGroup

	for _, api := range apis {
		wg.Add(1)
		go GetApi(api, &wg, resultChan)
	}

	for i, api := range postApis {
		wg.Add(1)
		go PostApi(api, payloads[i], &wg, resultChan)
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