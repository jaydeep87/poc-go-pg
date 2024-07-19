package topics

import (
	"fmt"
	"sync"
)

func isPrime(num int) bool {

	if num <= 2 {
		return false
	}
	for i:=2; i<num; i++ {
		if num%2 == 0 {
			return false
		}
	}
	return true
}

func GeneratePrime(sendPrim chan <-int, limit int, wg *sync.WaitGroup) {
	fmt.Println("Generate Prime called")

	defer wg.Done()
	for i:=0; i<=limit; i++ {
		if isPrime(i) {
			sendPrim <- i
		}
	}

	close(sendPrim)
}

func PrintPrime(receivedPrime <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for prime := range receivedPrime {
		fmt.Println(prime)
	}
}
