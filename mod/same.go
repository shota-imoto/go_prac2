package main

import "fmt"

func same_func() {
	fmt.Println("same package function")
}

type Same struct {
	Id   int
	Name string
}

func (same *Same) same_method() {
	fmt.Println("same package method")
}
