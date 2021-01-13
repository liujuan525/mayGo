package main

import (
	"flag"
)

var name2 string

func init()  {
	flag.StringVar(&name2, "name2", "everyone", "The greeting object.")
}

func main()  {
	flag.Parse()
	hello(name2)
}





