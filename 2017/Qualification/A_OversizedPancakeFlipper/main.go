package main

import "fmt"

func main() {
	fmt.Println(echo("hello."))
}

func echo(name string) string {
	return name
}