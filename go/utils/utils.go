package utils

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type MyFuncType func(a int, b int) int

func MeasurePerformance(f func(c* fiber.Ctx), c* fiber.Ctx) error {
	start := time.Now()
	f(c)
	duration := time.Since(start)
	return c.SendString(fmt.Sprintf("%v", duration.Milliseconds()))
}