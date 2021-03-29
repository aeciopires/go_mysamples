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
*/

package main

import (
	"fmt"
	"encoding/json"
)

type client struct {
	Name string `json:"name"` // Essa terceira coluna é uma tag para converter Name para name no JSON
	Id   string `json:"id"`   // Essa terceira coluna é uma tag para converter ID para id no JSON
	Poc  bool   `json:"poc"`  // Essa terceira coluna é uma tag para converter Poc para poc no JSON
}

func main() {
	sliceBytes := []byte(`{"name":"Aecio","id":"8j4rf","poc":false}`)

	var aecio client
	err := json.Unmarshal(sliceBytes, &aecio)
	if err != nil {
		fmt.Println("[ERROR]: ",err)
	}

	fmt.Println(aecio)
	fmt.Println(aecio.Poc)
}
