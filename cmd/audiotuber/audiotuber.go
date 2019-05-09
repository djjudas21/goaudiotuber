package main

import "fmt"
import "flag"

func main() {
	wordPtr := flag.String("word", "foo", "a string")
	flag.Parse()
	fmt.Println("word:", *wordPtr)
}
