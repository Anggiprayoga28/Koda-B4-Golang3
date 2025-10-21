package main

import (
	"Golang3/data"
	"fmt"
)

func main() {
	users := []data.Person{
		{Name: "Anggi Prayoga", Age: 26, City: "Batam"},
		{Name: "Ari Eka Saputra", Age: 18, City: "Depok"},
		{Name: "Federus Rudi", Age: 19, City: "Aceh"},
		{Name: "Anggi Fitriani", Age: 22, City: "Yogyakarta"},
		{Name: "Sidik Wisnu", Age: 35, City: "Bontang"},
	}

	var keyword string
	fmt.Print("Siapa yang anda cari? ")
	fmt.Scanln(&keyword)

	results := data.SearchPerson(users, keyword)

	if len(results) == 0 {
		fmt.Println("\nTidak ada hasil yang ditemukan.")
	} else {
		fmt.Printf("\nDitemukan %d hasil:\n", len(results))
		for i, person := range results {
			fmt.Printf("%d. %s (%d tahun) - %s\n", i+1, person.Name, person.Age, person.City)
		}
	}
}
