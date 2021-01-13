package main

import (
	"flag"
	"mayGo/q2/lib"
)

var name2 string

func init()  {
	flag.StringVar(&name2, "name2", "everyone", "The greeting object.")
}

func main()  {
	flag.Parse()
	lib.Hello(name2)
}





