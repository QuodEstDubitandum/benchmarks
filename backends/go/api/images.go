package api

import (
	"benchmarks/utils"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
)

// /decode
func DecodeImage(opt *utils.Options){
	for i:=0; i<opt.N; i++{
		file, err := os.Open("../../public/sample.jpg")
		if err != nil{
			fmt.Println(err)
		}
		defer file.Close()

		jpeg.Decode(file)
	}
}

// /encode
func EncodeImage(opt *utils.Options){
	outputFile, _ := os.Create("../../public/sample.png")
	defer outputFile.Close()

	png.Encode(outputFile, opt.BinaryImage)
}