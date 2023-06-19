package main

import "fmt"

// Run is going to be responsible for
// the instantiation and start up of
// our go application
func Run() error {
	fmt.Println("Starting up the app")
	return nil
}

func main() {
	fmt.Println("Go REST API Running")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
