package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GerarCartas() int {
	return rand.Intn(13) + 1
}

func main() {

	fmt.Println("Bem vindo ao jogo 21")
	fmt.Println("Vence quem tiver uma somatória de 21 pontos")

	var Jogar string
	fmt.Println("Deseja jogar? (S/N)")
	fmt.Scan(&Jogar)

	if Jogar != "S" && Jogar != "s" {
		fmt.Println("Encerrando o Jogo...")
	} else {
		CartaPC := GerarCartas()
		CartasUsuario := GerarCartas()

		fmt.Println("Sua primeira carta é:", CartasUsuario)
		time.Sleep(2 * time.Second)
		fmt.Println("A carta do seu adversário é:", CartaPC)
		time.Sleep(1 * time.Second)
		fmt.Println("Deseja continuar? (S/N)")
		fmt.Scan(&Jogar)

		if Jogar != "s" && Jogar != "S" {
			fmt.Println("Você perdeu")
		} else {
			SomaCartasU := CartasUsuario
			SomaCartasPC := CartaPC

			for SomaCartasU < 21 && Jogar == "s" || Jogar == "S" {
				SomaCartasU += GerarCartas()
				fmt.Println("Você tem:", SomaCartasU)
				time.Sleep(2 * time.Second)

				if SomaCartasU > 21 {
					fmt.Println("Você estourou! O computador venceu.")
					break
				}

				fmt.Println("Deseja pegar mais uma carta? (S/N)")
				time.Sleep(2 * time.Second)
				fmt.Scan(&Jogar)
			}

			for SomaCartasPC < 16 {
				SomaCartasPC += GerarCartas()
			}

			fmt.Println("Sua pontuação:", SomaCartasU)
			fmt.Println("Pontuação do adversário:", SomaCartasPC)

			if SomaCartasU == 21 {
				fmt.Println("Você venceu!")
			} else if SomaCartasU > 21 {
				fmt.Println("Você perdeu!")
			} else if SomaCartasPC == 21 {
				fmt.Println("O computador venceu!")
			} else if SomaCartasPC > 21 {
				fmt.Println("O jogador venceu!")
			} else if SomaCartasU > SomaCartasPC {
				fmt.Println("Você venceu!")
			} else if SomaCartasU < SomaCartasPC {
				fmt.Println("O computador venceu!")
			} else {
				fmt.Println("Empate!")
			}
		}
	}
}
