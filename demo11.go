package main

import "fmt"

var container = []string{"zero", "one", "two"} // 切片

func main()  {
	container := map[int]string{0:"zero", 1:"one", 2:"two"} // 字典
	fmt.Printf("The element is %q\n", container[1])
}


