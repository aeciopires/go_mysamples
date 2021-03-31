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

func main() {
	myArrayInt := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	myArrayStr := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println("before sort: ", myArrayInt)
	sort.Ints(myArrayInt)
	fmt.Println("After sort: ", myArrayInt)

	fmt.Println("")

	fmt.Println("Before sort: ", myArrayStr)
	sort.Strings(myArrayStr)
	fmt.Println("After sort: ", myArrayStr)
}
