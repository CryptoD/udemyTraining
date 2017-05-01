package main // pelne bledow, zignoruj

import "fmt"

func main() {

	var h string
	fmt.Println("Ustaw haslo:")
	fmt.Scan(&h)

	fmt.Println("Wpisz haslo:")

	for h != true {

		fmt.Println("Fuck you")
		break
	}

	for h == true {

		fmt.Println("Dobre haslo!")
		break
	}

}
