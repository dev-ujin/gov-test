package main

import (
	"flag"
	"fmt"
)
type Args struct {
	Name string
}
var args = Args{}
func main() {
	flag.StringVar(&args.Name, "name", "Me", "name to print in message")
	flag.Parse()
	fmt.Println("Hello, ", args.Name)
}