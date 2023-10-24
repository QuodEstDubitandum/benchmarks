package api

import (
	"benchmarks/utils"
	"encoding/json"
	"fmt"
)

// /serialize
func Serialize(opt *utils.Options){
	var obj utils.SerializeObject
	var jsonString []byte
	for i:=0; i<opt.N; i++{
		obj = opt.Obj
		jsonString, _ = json.Marshal(obj)
	}
	fmt.Println(jsonString[0])
}

// /deserialize
func Deserialize(opt *utils.Options){
	obj := utils.SerializeObject{}
	for i:=0; i<opt.N; i++{
		opt.Context.BodyParser(&obj)
	}
}