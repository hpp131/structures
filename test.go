package main

import (
	// "container/list"
	"fmt"
)


type student struct {
	name *string
	age  *int
}


func main()  {
	var a student
	s := new(student)
	if s == nil {
		fmt.Println(s)
		return
	}
	fmt.Println("is not nil")
	// list.New()
}

