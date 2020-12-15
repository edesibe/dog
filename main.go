package main

import (
	"fmt"
	"learn/ninja12/dog" // path should be relative path from GOPATH/src
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