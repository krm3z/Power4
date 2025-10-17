package logic

type WinPos struct {
	Row int
	Col int
}

type Game struct {
	Grid     [][]string 
	Rows     int      
	Columns  int    
	Current  string 
	Winner   string  
	GameOver bool   
	WinCells []WinPos 
}

func NewGameCustom(rows, cols int) *Game {
	grid := make([][]string, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]string, cols)
		for j := 0; j < cols; j++ {
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

func (g *Game) Play(col int) bool {
	if col < 0 || col >= g.Columns || g.GameOver {
		return false
	}

	for row := g.Rows - 1; row >= 0; row-- {
		if g.Grid[row][col] == "⚪" {
			g.Grid[row][col] = g.Current
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

func (g *Game) checkWin() {
	directions := [][2]int{
		{1, 0},  
		{0, 1},  
		{1, 1}, 
		{1, -1},
	}

	for r := 0; r < g.Rows; r++ {
		for c := 0; c < g.Columns; c++ {
			player := g.Grid[r][c]

			if player == "⚪" {
				continue
			}

			for _, dir := range directions {
				count := 1
				winCells := []WinPos{{r, c}}

				for k := 1; k < 4; k++ {
					nr := r + dir[0]*k
					nc := c + dir[1]*k

					if nr < 0 || nr >= g.Rows || nc < 0 || nc >= g.Columns {
						break
					}

					if g.Grid[nr][nc] == player {
						count++
						winCells = append(winCells, WinPos{nr, nc})
					} else {
						break
					}
				}

				if count >= 4 {
					g.Winner = player
					g.GameOver = true
					g.WinCells = winCells
					return
				}
			}
		}
	}
}

func (g *Game) Clone() *Game {
	newGrid := make([][]string, g.Rows)
	for i := 0; i < g.Rows; i++ {
		newGrid[i] = make([]string, g.Columns)
		copy(newGrid[i], g.Grid[i])
	}

	newGame := &Game{
		Grid:     newGrid,
		Rows:     g.Rows,
		Columns:  g.Columns,
		Current:  g.Current,
		Winner:   g.Winner,
		GameOver: g.GameOver,
		WinCells: make([]WinPos, len(g.WinCells)),
	}

	copy(newGame.WinCells, g.WinCells)
	return newGame
}

func (g *Game) Evaluate() int {
	score := 0

	checkWindow := func(window []string) int {
		countY := 0
		countR := 0

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

	for r := 0; r < g.Rows; r++ {
		for c := 0; c <= g.Columns-4; c++ {
			window := g.Grid[r][c : c+4]
			score += checkWindow(window)
		}
	}

	for c := 0; c < g.Columns; c++ {
		for r := 0; r <= g.Rows-4; r++ {
			window := []string{
				g.Grid[r][c],
				g.Grid[r+1][c],
				g.Grid[r+2][c],
				g.Grid[r+3][c],
			}
			score += checkWindow(window)
		}
	}

	for r := 3; r < g.Rows; r++ {
		for c := 0; c <= g.Columns-4; c++ {
			window := []string{
				g.Grid[r][c],
				g.Grid[r-1][c+1],
				g.Grid[r-2][c+2],
				g.Grid[r-3][c+3],
			}
			score += checkWindow(window)
		}
	}

	for r := 0; r <= g.Rows-4; r++ {
		for c := 0; c <= g.Columns-4; c++ {
			window := []string{
				g.Grid[r][c],
				g.Grid[r+1][c+1],
				g.Grid[r+2][c+2],
				g.Grid[r+3][c+3],
			}
			score += checkWindow(window)
		}
	}

	return score
}
