package data

import "strings"

type Person struct {
	Name string
}

func SearchPerson(users []Person, keyword string) []string {
	var results []string

	for _, user := range users {
		if strings.Contains(strings.ToLower(user.Name), strings.ToLower(keyword)) {
			results = append(results, user.Name)
		}
	}

	return results
}
