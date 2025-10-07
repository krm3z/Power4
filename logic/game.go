package logic

import "fmt"

const (
	Rows    = 6
	Columns = 7
)

// Game représente l’état du jeu
type Game struct {
	Grid      [Rows][Columns]string
	Current   string // "R" ou "Y"
	Winner    string
	GameOver  bool
	TurnCount int
}

// NewGame crée une nouvelle partie vide
func NewGame() *Game {
	g := &Game{
		Current: "R", // le joueur rouge commence
	}
	for r := 0; r < Rows; r++ {
		for c := 0; c < Columns; c++ {
			g.Grid[r][c] = "."
		}
	}
	return g
}

// Play ajoute un jeton dans une colonne donnée
func (g *Game) Play(col int) bool {
	if col < 0 || col >= Columns {
		return false
	}

	// Trouver la première case vide depuis le bas
	for row := Rows - 1; row >= 0; row-- {
		if g.Grid[row][col] == "." {
			g.Grid[row][col] = g.Current
			g.TurnCount++
			if g.checkWin(row, col) {
				g.GameOver = true
				g.Winner = g.Current
			}
			g.switchPlayer()
			return true
		}
	}
	return false
}

// switchPlayer change le joueur courant
func (g *Game) switchPlayer() {
	if g.Current == "R" {
		g.Current = "Y"
	} else {
		g.Current = "R"
	}
}

// checkWin vérifie s’il y a un alignement de 4
func (g *Game) checkWin(row, col int) bool {
	directions := [][2]int{
		{0, 1},  // horizontal
		{1, 0},  // vertical
		{1, 1},  // diagonale ↘
		{1, -1}, // diagonale ↙
	}

	for _, dir := range directions {
		count := 1
		for i := 1; i < 4; i++ {
			r, c := row+dir[0]*i, col+dir[1]*i
			if r < 0 || r >= Rows || c < 0 || c >= Columns || g.Grid[r][c] != g.Current {
				break
			}
			count++
		}
		for i := 1; i < 4; i++ {
			r, c := row-dir[0]*i, col-dir[1]*i
			if r < 0 || r >= Rows || c < 0 || c >= Columns || g.Grid[r][c] != g.Current {
				break
			}
			count++
		}
		if count >= 4 {
			return true
		}
	}
	return false
}

// Print affiche la grille dans la console (debug)
func (g *Game) Print() {
	for r := 0; r < Rows; r++ {
		for c := 0; c < Columns; c++ {
			fmt.Print(g.Grid[r][c], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}