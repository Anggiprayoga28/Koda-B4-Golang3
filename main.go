package main

import (
	"Golang3/data"
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Selesai")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
			os.Exit(1)
		}
	}()

	users := []data.Person{
		{Name: "John Doe"},
		{Name: "Jane Smith"},
		{Name: "John Carter"},
		{Name: "Alice Johnson"},
		{Name: "Bob Wilson"},
	}

	var keyword string
	fmt.Print("Siapa yang anda cari? ")
	fmt.Scanln(&keyword)

	results := data.SearchPerson(&users, &keyword)

	fmt.Println(results)

	os.Exit(0)
}
