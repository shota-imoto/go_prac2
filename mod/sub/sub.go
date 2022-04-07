package sub

import "fmt"

func SubFunc() {
	fmt.Println("sub package function")
}

type Sub struct {
	Id   int
	Name string
}

func (sub *Sub) SubMethod() {
	fmt.Println("sub package method")
}
