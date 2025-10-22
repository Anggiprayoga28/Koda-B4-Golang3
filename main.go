package main

import (
	"Golang3/data"
	"fmt"
)

func main() {
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
}
