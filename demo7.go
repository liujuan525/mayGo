package main

import (
	"flag"
	"fmt"
)

func main()  {
	// 方法一
	//var name string
	//flag.StringVar(&name, "name", "everyone", "The greeting object.")
	//flag.Parse()
	//fmt.Printf("hello, %s\n", name)

    // 方法二
    var name = flag.String("name", "everyone", "The greeting object.")
	fmt.Printf("hello, %s\n", *name)

}





