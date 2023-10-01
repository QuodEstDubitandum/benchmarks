package utils

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Options struct{
	Context *fiber.Ctx
	N int
	Array []int
	Cores int
}

func MeasurePerformance(f func(opt *Options), opt *Options) error {
	start := time.Now()
	for i:=0; i<5; i++{
		f(opt)
	}
	duration := time.Since(start)
	return opt.Context.SendString(fmt.Sprintf("%v", duration.Milliseconds()/5))
}