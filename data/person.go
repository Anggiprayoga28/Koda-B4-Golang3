package data

import "strings"

type Person struct {
	Name string
	Age  int
	City string
}

func SearchPerson(users []Person, keyword string) []Person {
	var results []Person

	for _, user := range users {
		if strings.Contains(strings.ToLower(user.Name), strings.ToLower(keyword)) {
			results = append(results, user)
		}
	}

	return results
}
