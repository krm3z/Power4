package logic

type Game struct {
	Grid      [][]string
	Rows      int
	Columns   int
	Current   string
	Winner    string
	GameOver  bool
}

// Création d'une grille personnalisée
func NewGameCustom(rows, cols int) *Game {
	grid := make([][]string, rows)
	for i := range grid {
		grid[i] = make([]string, cols)
		for j := range grid[i] {
			grid[i][j] = "⚪"
		}
	}
	return &Game{
		Grid:     grid,
		Rows:     rows,
		Columns:  cols,
		Current:  "R",
		GameOver: false,
	}
}

// Jouer un pion dans une colonne
func (g *Game) Play(col int) bool {
	if col < 0 || col >= g.Columns || g.GameOver {
		return false
	}

	for i := g.Rows - 1; i >= 0; i-- {
		if g.Grid[i][col] == "⚪" {
			g.Grid[i][col] = g.Current
			g.checkWin()
			if !g.GameOver {
				if g.Current == "R" {
					g.Current = "Y"
				} else {
					g.Current = "R"
				}
			}
			return true
		}
	}
	return false
}

// Vérifie s’il y a un gagnant
func (g *Game) checkWin() {
	dirs := [][2]int{{1, 0}, {0, 1}, {1, 1}, {1, -1}}

	for r := 0; r < g.Rows; r++ {
		for c := 0; c < g.Columns; c++ {
			player := g.Grid[r][c]
			if player == "⚪" {
				continue
			}
			for _, d := range dirs {
				count := 1
				for k := 1; k < 4; k++ {
					nr := r + d[0]*k
					nc := c + d[1]*k
					if nr < 0 || nr >= g.Rows || nc < 0 || nc >= g.Columns {
						break
					}
					if g.Grid[nr][nc] == player {
						count++
					} else {
						break
					}
				}
				if count >= 4 {
					g.Winner = player
					g.GameOver = true
					return
				}
			}
		}
	}

	// Vérifie s’il n’y a plus de place
	full := true
	for c := 0; c < g.Columns; c++ {
		if g.Grid[0][c] == "⚪" {
			full = false
			break
		}
	}
	if full {
		g.GameOver = true
	}
}

// Clone profond du plateau
func (g *Game) Clone() *Game {
	newGrid := make([][]string, g.Rows)
	for i := range g.Grid {
		newGrid[i] = make([]string, g.Columns)
		copy(newGrid[i], g.Grid[i])
	}
	return &Game{
		Grid:     newGrid,
		Rows:     g.Rows,
		Columns:  g.Columns,
		Current:  g.Current,
		Winner:   g.Winner,
		GameOver: g.GameOver,
	}
}

// Évaluation de la position (pour IA)
func (g *Game) Evaluate() int {
	score := 0
	rows, cols := g.Rows, g.Columns

	checkWindow := func(window []string) int {
		countY, countR := 0, 0
		for _, cell := range window {
			if cell == "Y" {
				countY++
			} else if cell == "R" {
				countR++
			}
		}
		if countY == 4 {
			return 10000
		}
		if countY == 3 && countR == 0 {
			return 50
		}
		if countY == 2 && countR == 0 {
			return 10
		}
		if countR == 3 && countY == 0 {
			return -80
		}
		return 0
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols-3; c++ {
			score += checkWindow(g.Grid[r][c : c+4])
		}
	}

	for c := 0; c < cols; c++ {
		for r := 0; r < rows-3; r++ {
			window := []string{g.Grid[r][c], g.Grid[r+1][c], g.Grid[r+2][c], g.Grid[r+3][c]}
			score += checkWindow(window)
		}
	}

	for r := 3; r < rows; r++ {
		for c := 0; c < cols-3; c++ {
			window := []string{g.Grid[r][c], g.Grid[r-1][c+1], g.Grid[r-2][c+2], g.Grid[r-3][c+3]}
			score += checkWindow(window)
		}
	}

	for r := 0; r < rows-3; r++ {
		for c := 0; c < cols-3; c++ {
			window := []string{g.Grid[r][c], g.Grid[r+1][c+1], g.Grid[r+2][c+2], g.Grid[r+3][c+3]}
			score += checkWindow(window)
		}
	}

	return score
}
