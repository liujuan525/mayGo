package main

import (
	"flag"
	"fmt"
)

func main()  {
	var name8 = getTheFlag()
	flag.Parse()
	fmt.Printf("hello, %s\n", *name8)
}

func getTheFlag() *string  {
	return flag.String("name", "everyone", "The greeting object.")
}



