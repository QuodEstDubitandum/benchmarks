package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func PrimeNumbers(c *fiber.Ctx){
	p := 2
	n := 100000000
	primes := make([]bool, n+1)

	// exclude 1
	j := n-1

	// sieve of eratosthenes
	for p*p <= n{
		if !primes[p-1]{
			for i:=p*p; i<=n; i+=p{
				if !primes[i-1]{
					j -= 1
				}
				primes[i-1] = true
			} 
		}
		p += 1
	}
	fmt.Printf("Number of prime numbers until %d: %v", n, j)
}