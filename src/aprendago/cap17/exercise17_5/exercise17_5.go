/*
Aulas sobre aplicações com JSON:
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=113
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=114
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=115
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=116
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=117
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=118
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=119
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=120
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=121
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=122
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=123
https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=124

OBS.: Os structs que serão exportados para JSON precisam começar com letra maiúscula para serem visíveis às funções que os manipularão, caso contrário não funcionará corretamente conforme explicado nesta aula.
 https://www.youtube.com/watch?v=-tU2PSY8F5w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=114
 https://stackoverflow.com/a/21825458/8507770
 https://narkive.com/8BNGRCKN.2
 https://stackoverflow.com/questions/43444248/why-is-the-format-string-of-struct-field-always-lower-case
 https://eager.io/blog/go-and-json/
 https://gobyexample.com/json

 Method chaining in go: https://golangbyexample.com/method-chaining-go/
*/

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type age []user

func (p age) Len() int {
	return len(p)
}

func (p age) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func (p age) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type surname []user

func (p surname) Len() int {
	return len(p)
}

func (p surname) Less(i, j int) bool {
	return p[i].Last < p[j].Last
}

func (p surname) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func showResults(u []user) {
	for index, value := range u {
		fmt.Println(index, "\tAge:", value.Age, "\tComplete name:", value.First, value.Last, "\n")
		for _, citation := range value.Sayings {
			fmt.Println("\t", citation)
		}
		fmt.Println("\n")
	}
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println("Before sort", users)

	// Sort Sayings of the array of users
	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	fmt.Println("\nAfer sort")

	sort.Sort(age(users))
	fmt.Println("Sort by age:\n")
	showResults(users)

	sort.Sort(surname(users))
	fmt.Println("Sort by surname:\n")
	showResults(users)

}
