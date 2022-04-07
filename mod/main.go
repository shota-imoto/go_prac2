package main

import "github.com/shota-imoto/mod/sub"

func main() {
	same_func()
	same := Same{Id: 1, Name: "OK"}
	same.same_method()
	sub.SubFunc()
	obj := sub.Sub{Id: 2, Name: "SUB"}
	obj.SubMethod()
}
