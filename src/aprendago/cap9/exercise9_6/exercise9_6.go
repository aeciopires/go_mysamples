package main

import (
	"fmt"
)

func main() {
	// slice of int, length, capacity of abjascent array
	slice := make([]string, 26, 27)
	slice = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Println(slice, len(slice), cap(slice))
	for aux := 0; aux < len(slice); aux++ {
		fmt.Println(slice[aux])
	}
}
