package utils

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)
type SerializeObject struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
	D string `json:"d"`
	E string `json:"e"`
	F string `json:"f"`
	G string `json:"g"`
	H string `json:"h"`
	I string `json:"i"`
	J string `json:"j"`
	K string `json:"k"`
	L string `json:"l"`
	M string `json:"m"`
	N string `json:"n"`
	O string `json:"o"`
	P string `json:"p"`
	Q string `json:"q"`
}

type Options struct{
	Context *fiber.Ctx
	N int
	Array []int
	Cores int
	DB *sql.DB
	Obj SerializeObject
}

func MeasurePerformance(f func(opt *Options), opt *Options) error {
	start := time.Now()
	for i:=0; i<5; i++{
		f(opt)
	}
	duration := time.Since(start)
	return opt.Context.SendString(fmt.Sprintf("%v", duration.Milliseconds()/5))
}