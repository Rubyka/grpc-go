package main

import (
	"fmt"
)

func main() {
	doSimple();
	fmt.Println("vim-go")
}

func doSimple(){
	var sm = simple.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}
	fmt.Println(sm)
}
