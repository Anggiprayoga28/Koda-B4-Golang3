package main

import (
	"Golang3/data"
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("\nSelesai")

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

	if len(results) == 0 {
		fmt.Println("Tidak ada hasil ditemukan.")
		os.Exit(0)
	} else {
		fmt.Printf("Ditemukan %d hasil.\n", len(results))
		os.Exit(0)
	}
}
