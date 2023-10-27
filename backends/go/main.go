package main

import (
	api "benchmarks/api"
	"benchmarks/database"
	utils "benchmarks/utils"
	"database/sql"
	"fmt"
	"image/jpeg"
	"os"
	"runtime"
	"strings"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)


func main(){
	cores := runtime.NumCPU()
	app := fiber.New()

	db, err := sql.Open("sqlite3", "./../../sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// _, err = db.Exec("PRAGMA journal_mode=WAL;")
	// if err != nil {
	// 	panic(err)
	// }

	// only to create the db from the csv file
	// database.Seed(db)
	
	
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

	app.Get("/db/select", func(c *fiber.Ctx) error {
		time := utils.MeasurePerformance(database.SelectQuery, &utils.Options{Context: c, N: 10, DB: db})
		return time
	})

	app.Get("/db/insert", func(c *fiber.Ctx) error {
		time := utils.MeasurePerformance(database.InsertQuery, &utils.Options{Context: c, N: 10, DB: db})
		_, err := db.Exec(`
			DELETE FROM electric_cars
			WHERE vin = "test" AND county = "test"
		`)

		if err != nil{
			fmt.Println(err)
		}
		return time
	})

	app.Get("/db/update", func(c *fiber.Ctx) error {
		time := utils.MeasurePerformance(database.UpdateQuery, &utils.Options{Context: c, N: 10, DB: db})
		return time
	})

	app.Get("/db/delete", func(c *fiber.Ctx) error {
		time := database.DeleteQuery(&utils.Options{Context: c, N: 50, DB: db}) 
		return time
	})

	app.Get("/json/serialize", func(c *fiber.Ctx) error {
		sampleString := "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet."
		time := utils.MeasurePerformance(api.Serialize, &utils.Options{Context: c, N: 10000, Obj: utils.SerializeObject{A: sampleString, B: sampleString, C: sampleString, D: sampleString, E: sampleString, F: sampleString, G: sampleString, H: sampleString, I: sampleString, J: sampleString, K: sampleString, L: sampleString, M: sampleString, N: sampleString, O: sampleString, P: sampleString, Q: sampleString}}) 
		return time
	})

	app.Post("/json/deserialize", func(c *fiber.Ctx) error {
		time := utils.MeasurePerformance(api.Deserialize, &utils.Options{Context: c, N: 10000}) 
		return time
	})

	app.Get("/image/decode", func(c *fiber.Ctx) error {
		time := utils.MeasurePerformance(api.DecodeImage, &utils.Options{Context: c, N: 10}) 
		return time
	})

	app.Get("/image/encode", func(c *fiber.Ctx) error {
		// decoding to get binary data for encoding
		file, err := os.Open("../../public/sample.jpg")
		if err != nil{
			fmt.Println(err)
		}
		defer file.Close()

		img, _ := jpeg.Decode(file)
		time := utils.MeasurePerformance(api.EncodeImage, &utils.Options{Context: c, BinaryImage: img}) 
		return time
	})

	app.Listen(":3001")
}