package api

import (
	"benchmarks/utils"
	"fmt"
	"io"
	"os"
	"sync"
)

// /write
func WriteFileSync(opt *utils.Options){
	for i:=0; i<opt.Cores; i++{
		file, _ := os.Create("test.txt")

		for i:=0; i<opt.N; i++{
			file.WriteString("Gotta go fast\n")
		}

		file.Close()
		os.Remove("test.txt")
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
			file, _ := os.Create(fmt.Sprintf("test%d.text", temp))

			for i:=0; i<opt.N; i++{
				file.WriteString("Gotta go fast\n")
			}

			file.Close()

			os.Remove(fmt.Sprintf("test%d.text", temp))
		}()
	}
	wg.Wait()
}

// /read
func ReadFileSync(opt *utils.Options){
	for i:=0; i<opt.Cores; i++{
		file, _ := os.Open(fmt.Sprintf("test%d.txt", i))
		io.ReadAll(file)
		file.Close()
	}
}

// /read-parallel
func ReadFileParallel(opt *utils.Options){
	var wg sync.WaitGroup
	for i:=0; i<opt.Cores; i++{
		temp := i
		wg.Add(1)
		go func(){
			defer wg.Done()
			file, _ := os.Open(fmt.Sprintf("test%d.txt", temp))
			io.ReadAll(file)
			file.Close()
		}()
	}
	wg.Wait()
}