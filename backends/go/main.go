package main

import (
	api "benchmarks/api"
	utils "benchmarks/utils"

	"github.com/gofiber/fiber/v2"
)


func main(){
	app := fiber.New()

	app.Get("/algorithms/prime", func(c *fiber.Ctx) error {
        return utils.MeasurePerformance(api.PrimeNumbers, &utils.Options{Context: c, N: 10000000})
    })

	app.Get("/algorithms/fast_fibonacci", func(c *fiber.Ctx) error {
        return utils.MeasurePerformance(api.FastFibonacci, &utils.Options{Context: c, N: 1000000})
    })

	app.Get("/algorithms/quicksort", func(c *fiber.Ctx) error {
		size := 1000000
		unsortedArr := make([]int, size)
		for i:=size-1; i>=0; i--{
			unsortedArr[i] = size - i
		}
        return utils.MeasurePerformance(api.Quicksort, &utils.Options{Context: c, QuicksortArray: unsortedArr, N: size})
    })

	app.Listen(":3001")
}