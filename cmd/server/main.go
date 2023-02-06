package main

import "fmt"

func Run() error {
	fmt.Println("Starting up our application")
	return nil
}

func main() {
	fmt.Println("Paisagens api")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
