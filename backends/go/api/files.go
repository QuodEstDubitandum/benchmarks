package api

import (
	"benchmarks/utils"
	"fmt"
	"os"
	"sync"
)

// /write
func WriteFileSync(opt *utils.Options){
	for i:=0; i<opt.Cores; i++{
		file, err := os.Create("test.txt")
		if err != nil{
			fmt.Println("Could not create file")
			return
		}

		for i:=0; i<opt.N; i++{
			file.WriteString("Gotta go fast\n")
		}

		file.Close()

		err = os.Remove("test.txt")
		if err != nil {
			fmt.Println("Error deleting the file:", err)
			return
		}
	}
}

// /write-parallel
func WriteFileParallel(opt *utils.Options){
	var wg sync.WaitGroup
	for i:=0; i<opt.Cores; i++{
		temp := i
		wg.Add(1)
		go func(){
			defer wg.Done()
			file, err := os.Create(fmt.Sprintf("test%d.text", temp))
			if err != nil{
				fmt.Println("Could not create file")
				return
			}

			for i:=0; i<opt.N; i++{
				file.WriteString("Gotta go fast\n")
			}

			file.Close()

			err = os.Remove(fmt.Sprintf("test%d.text", temp))
			if err != nil {
				fmt.Println("Error deleting the file:", err)
				return
			}
		}()
	}
	wg.Wait()
}
