package main

import (
	"fmt"
)

type autoMobile struct {
	ports int
	color string
}

type suv struct {
	autoMobile
	wheelsFourTraction bool
}

type sedan struct {
	autoMobile
	luxeModel bool
}

func main() {
	car1 := suv{
		wheelsFourTraction: false,
		autoMobile: autoMobile{
			ports: 4,
			color: "black",
		},
	}

	car2 := sedan{
		luxeModel: false,
		autoMobile: autoMobile{
			ports: 2,
			color: "white",
		},
	}

	fmt.Println(car1)
	fmt.Println(car1.wheelsFourTraction)
	fmt.Println(car2)
	fmt.Println(car2.color)
}
