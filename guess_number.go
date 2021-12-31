package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	number_guess := rand.Intn(100-0) + 0

	is_loop_active := true
	numero_de_jogadas := 0

	fmt.Println("BEM VINDO AO JOGO DE ACERTAR O NÚMERO!!!!")

	var numero_user int

	for is_loop_active == true {
		fmt.Print("Digite o número: ")
		fmt.Scanln(&numero_user)
		numero_de_jogadas++

		if numero_user == number_guess {
			fmt.Println("Você acertou! O número era: ", numero_user)
			fmt.Println("Número de jogadas: ", numero_de_jogadas)
			break
		} else {
			if numero_user < number_guess {
				fmt.Println("O número digitado é menor.")
			} else {
				fmt.Println("O número digitado é maior.")
			}
		}
	}
}
