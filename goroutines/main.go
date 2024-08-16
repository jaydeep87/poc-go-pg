// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func isPrime(num int) bool {

	if num <= 2 {
		return false
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func GenratePrime(limit int, resultChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i:=0; i<=limit; i++ {
		if isPrime(i) {
			resultChan <- i
		}
	}
	close(resultChan)

}

func PrintPrime(resultChan <-chan int, wg *sync.WaitGroup){
	defer wg.Done()

	for primeNo := range resultChan {
		fmt.Println(primeNo)
	}

}


func main() {
	fmt.Println("Enter limit: ")

	// var then variable name then variable type
	var limit int

	// Taking input from user
	fmt.Scanln(&limit)
	resultChan := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go GenratePrime(limit, resultChan, &wg)
	go PrintPrime(resultChan, &wg)
	wg.Wait()
	// fmt.Println(limit)
	// fmt.Println(isPrime(5))


}
