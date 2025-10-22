package data

import (
	"fmt"
	"strings"
)

type Person struct {
	Name string
}

func SearchPerson(users *[]Person, keyword *string) []string {
	defer fmt.Println("Pencarian selesai")

	var results []string

	for _, user := range *users {
		if strings.Contains(strings.ToLower(user.Name), strings.ToLower(*keyword)) {
			results = append(results, user.Name)
		}
	}

	return results
}
