package main

import "fmt"

func main() {
	var s []byte
	s = append(s, []byte("Hello, Debugging!")...)
	fmt.Println(string(s))
}
