package main

import (
	//"github.com/GoesToEleven/go-programming/code_samples/008-ninja-level-twelve/01/dog"
	"fmt"
	"learn/ninja12/dog"
)

type Dog struct {
	name string
	age int
}

func main(){
	mica := Dog{
		name: "mica",
		age: dog.Years(10),
	}
	fmt.Printf("%+v\n",mica)
}