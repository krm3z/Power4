package main

import (
	"fmt"
	"power4/logic"
)

func main() {
	game := logic.NewGame()

	game.Play(0)
	game.Play(0)
	game.Play(1)
	game.Play(1)
	game.Play(2)
	game.Play(2)
	game.Play(3) // ici le joueur Rouge gagne

	game.Print()
	fmt.Println("Gagnant:", game.Winner)
}
