package main

import (
	"fmt"
)

// Aulas sobre Ponteiros:
// https://www.youtube.com/watch?v=mOM0qTB5ppU&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=109
// https://www.youtube.com/watch?v=mOM0qTB5ppU&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=110
// https://www.youtube.com/watch?v=mOM0qTB5ppU&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=111
// https://www.youtube.com/watch?v=mOM0qTB5ppU&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=112

/*
- Dica: a maneira "correta" para fazer dereference de um valor em um struct seria (*valor).campo
- Mas consta uma exceção na documentação. Link: https://golang.org/ref/spec#Selectors​
- "As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method),
- → x.f is shorthand for (*x).f." ←
- Ou seja, podemos usar tanto o atalho p1.nome quanto o tradicional (*p1).nome
*/

type person struct {
	name    string
	surname string
}

func changeMe(aux *person, newSurname string) {
	aux.surname = newSurname
	//(*aux).surname = newSurname
}

func main() {
	son := person{
		name:    "Aecio",
		surname: "Pires",
	}
	fmt.Println("Before...")
	fmt.Println("My name is: ", son)

	changeMe(&son, "dos Santos")

	fmt.Println("After...")
	fmt.Println("My name is: ", son)
}
