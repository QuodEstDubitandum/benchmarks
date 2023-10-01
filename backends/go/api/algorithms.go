package api

import (
	"benchmarks/utils"
	"fmt"
	"math/big"
)

// /prime
func PrimeNumbers(opt *utils.Options){
	p := 2
	n := opt.N
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

// /fast-fibonacci
func FastFibonacci(opt *utils.Options){
	n := opt.N
	baseMatrix := [2][2]*big.Int{
		{big.NewInt(1),big.NewInt(1)},
		{big.NewInt(1),big.NewInt(0)},
	}
	powerMatrix(baseMatrix, n)
}

func powerMatrix(matrix [2][2]*big.Int, n int) [2][2]*big.Int{
	if n==1{
		return matrix
	}
	if n%2==1{
		return matrixMult(matrix, powerMatrix(matrix, n-1))
	}
	root := powerMatrix(matrix, n/2)
	return matrixMult(root, root)
}

func matrixMult(a [2][2]*big.Int, b [2][2]*big.Int) [2][2]*big.Int{
	var res [2][2]*big.Int
	for i := range res {
		for j := range res[i] {
			res[i][j] = new(big.Int)
		}
	}
	res[0][0].Mul(a[0][0], b[0][0]).Add(res[0][0], new(big.Int).Mul(a[0][1], b[1][0]))
	res[0][1].Mul(a[0][0], b[0][1]).Add(res[0][1], new(big.Int).Mul(a[0][1], b[1][1]))
	res[1][0].Mul(a[1][0], b[0][0]).Add(res[1][0], new(big.Int).Mul(a[1][1], b[1][0]))
	res[1][1].Mul(a[1][0], b[0][1]).Add(res[1][1], new(big.Int).Mul(a[1][1], b[1][1]))

	return res
}

// /quicksort
func Quicksort(opt *utils.Options){
	baseArray := opt.Array
	
	fmt.Println(divideArray(baseArray)[0])
}

func divideArray(array []int) []int{
	length := len(array)
	if length < 2{
		return array
	}
	// the fact that go rounds integers down if you get a float by dividing an int is really interesting
	pivot := length/2
	left := 0
	
	array[length-1], array[pivot] = array[pivot], array[length-1]

	for i := range array {
		if array[i] < array[length-1] {
			array[i], array[left] = array[left], array[i]
			left++
		}
	}
	array[left], array[length-1] = array[length-1], array[left]

	divideArray(array[:left])
	divideArray(array[left+1:])

	return array
}

// /two-sum
func TwoSum(opt *utils.Options){
	n := opt.N
	hashmap := make(map[int]int, n)
	arr := opt.Array
	count := 0

	for _, number := range arr{
		complement := n - number
		count += hashmap[complement]
		hashmap[number] += 1
	}
	fmt.Println(count)
}