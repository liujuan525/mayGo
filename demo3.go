package main

import (
	"flag"
	"fmt"
	"os"
)

var name1 string

func init()  {
	//flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.StringVar(&name1, "name1", "everyone", "The greeting object.")
}

func main()  {
	flag.Parse()
	fmt.Printf("hello, %s\n", name1)
}