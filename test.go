package main

import "container/list"


type student struct {
	name *string
	age  *int
}


func main()  {
	s := new(student)
	if s == nil {
		fmt.Println(s)
		return
	}
	fmt.Println("is not nil")
	list.New()
}

