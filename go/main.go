package main

import (
	api "benchmarks/api"
	utils "benchmarks/utils"

	"github.com/gofiber/fiber/v2"
)

func main(){
	app := fiber.New()

	app.Get("/algorithms/prime", func(c *fiber.Ctx) error {
        return utils.MeasurePerformance(api.PrimeNumbers, c)
    })

	app.Listen(":3001")
}