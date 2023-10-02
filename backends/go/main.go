package main

import (
	api "benchmarks/api"
	utils "benchmarks/utils"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/gofiber/fiber/v2"
)


func main(){
	cores := runtime.NumCPU()
	fmt.Println(cores)
	app := fiber.New()

	app.Get("/algorithms/prime", func(c *fiber.Ctx) error {
        return utils.MeasurePerformance(api.PrimeNumbers, &utils.Options{Context: c, N: 10000000})
    })

	app.Get("/algorithms/fast-fibonacci", func(c *fiber.Ctx) error {
        return utils.MeasurePerformance(api.FastFibonacci, &utils.Options{Context: c, N: 1000000})
    })

	app.Get("/algorithms/quicksort", func(c *fiber.Ctx) error {
		size := 1000000
		unsortedArr := make([]int, size)
		for i:=size-1; i>=0; i--{
			unsortedArr[i] = size - i
		}
        return utils.MeasurePerformance(api.Quicksort, &utils.Options{Context: c, Array: unsortedArr, N: size})
    })

	app.Get("/algorithms/two-sum", func(c *fiber.Ctx) error {
		size := 1000000
		arr := make([]int, size)
		for i:=0; i<size; i++{
			arr[i] = i+1
		}
        return utils.MeasurePerformance(api.TwoSum, &utils.Options{Context: c, Array: arr, N: size})
    })

	app.Get("/files/write", func(c *fiber.Ctx) error {
		return utils.MeasurePerformance(api.WriteFileSync, &utils.Options{Context: c, N: 10000, Cores: cores})
	})

	app.Get("/files/write-parallel", func(c *fiber.Ctx) error {
		return utils.MeasurePerformance(api.WriteFileParallel, &utils.Options{Context: c, N: 10000, Cores: cores})
	})

	app.Get("/files/read", func(c *fiber.Ctx) error {
		// creating files
		n := 1000000
		data := strings.Repeat("Gotta go fast\n", n)
		for i:=0; i<cores; i++{
			file, _ := os.Create(fmt.Sprintf("test%d.txt", i))
			file.WriteString(data)
			file.Close()
		}
		// measuring performance of read
		time := utils.MeasurePerformance(api.ReadFileSync, &utils.Options{Context: c, N: n, Cores: cores})
		// deleting files
		for i:=0; i<cores; i++{
			os.Remove(fmt.Sprintf("test%d.txt", i))
		}
		return time
	})

	app.Get("/files/read-parallel", func(c *fiber.Ctx) error {
		// creating files
		n := 1000000
		data := strings.Repeat("Gotta go fast\n", n)
		for i:=0; i<cores; i++{
			file, _ := os.Create(fmt.Sprintf("test%d.txt", i))
			file.WriteString(data)
			file.Close()
		}
		// measuring performance of read
		time := utils.MeasurePerformance(api.ReadFileParallel, &utils.Options{Context: c, N: n, Cores: cores})
		// deleting files
		for i:=0; i<cores; i++{
			os.Remove(fmt.Sprintf("test%d.txt", i))
		}
		return time
	})

	app.Listen(":3001")
}