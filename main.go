package main

import (
	"fmt"
	"html/template"
	"math"
	"math/rand"
	"net/http"
	"power4/logic"
	"strings"
	"time"
)

var currentGame *logic.Game
var currentLevel string

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	path := fmt.Sprintf("templates/%s.html", tmpl)
	t, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, "Erreur de template : "+err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

func indexHandler(w http.ResponseWriter, r *http.Request) { renderTemplate(w, "index", nil) }
func menuHandler(w http.ResponseWriter, r *http.Request)  { renderTemplate(w, "menu", nil) }
func botHandler(w http.ResponseWriter, r *http.Request)   { renderTemplate(w, "bot", nil) }
func localHandler(w http.ResponseWriter, r *http.Request) { renderTemplate(w, "local", nil) }

func initGame(level string) {
	level = strings.ToLower(level)
	switch level {
	case "easy":
		currentGame = logic.NewGameCustom(6, 7)
	case "normal":
		currentGame = logic.NewGameCustom(6, 9)
	case "hard":
		currentGame = logic.NewGameCustom(7, 8)
	case "god":
		currentGame = logic.NewGameCustom(7, 8)
	default:
		level = "easy"
		currentGame = logic.NewGameCustom(6, 7)
	}
	currentLevel = level
	fmt.Printf("🎮 Nouvelle partie : %s (%dx%d)\n", strings.Title(level), currentGame.Rows, currentGame.Columns)
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	level := strings.ToLower(r.URL.Query().Get("level"))
	if currentGame == nil {
		if level == "" {
			level = "easy"
		}
		initGame(level)
	} else if level != "" && level != currentLevel {
		initGame(level)
	}

	if r.Method == "POST" && !currentGame.GameOver && currentGame.Current == "R" {
		col := 0
		fmt.Sscanf(r.FormValue("col"), "%d", &col)
		if currentGame.Play(col) && !currentGame.GameOver {
			time.Sleep(300 * time.Millisecond)
			botMove(currentLevel)
		}
	}

	data := struct {
		*logic.Game
		Level    string
		LevelRaw string
	}{
		Game:     currentGame,
		Level:    strings.Title(currentLevel),
		LevelRaw: currentLevel,
	}

	renderTemplate(w, "game", data)
}

func botMove(level string) {
	if currentGame == nil || currentGame.GameOver {
		return
	}

	rand.Seed(time.Now().UnixNano())
	cols := currentGame.Columns
	level = strings.ToLower(level)

	switch level {
	case "easy":
		for {
			c := rand.Intn(cols)
			if currentGame.Play(c) {
				break
			}
		}
	case "normal":
		if rand.Float64() < 0.6 {
			for c := 0; c < cols; c++ {
				tmp := currentGame.Clone()
				if tmp.Play(c) && tmp.Winner == "Y" {
					currentGame.Play(c)
					currentGame.Current = "R"
					return
				}
			}
		}
		for {
			c := rand.Intn(cols)
			if currentGame.Play(c) {
				break
			}
		}
	case "hard":
		bestScore := math.Inf(-1)
		bestCol := -1
		depth := 3
		for c := 0; c < cols; c++ {
			tmp := currentGame.Clone()
			if tmp.Play(c) {
				score := minimax(tmp, depth, false)
				if score > bestScore {
					bestScore = score
					bestCol = c
				}
			}
		}
		if bestCol == -1 {
			bestCol = rand.Intn(cols)
		}
		currentGame.Play(bestCol)
	case "god":
		bestScore := math.Inf(-1)
		bestCol := -1
		depth := 5
		for c := 0; c < cols; c++ {
			tmp := currentGame.Clone()
			if tmp.Play(c) {
				score := minimax(tmp, depth, false)
				if score > bestScore {
					bestScore = score
					bestCol = c
				}
			}
		}
		if bestCol == -1 {
			bestCol = rand.Intn(cols)
		}
		currentGame.Play(bestCol)
	}

	currentGame.Current = "R"
}

func minimax(g *logic.Game, depth int, isMax bool) float64 {
	if g.GameOver || depth == 0 {
		return float64(g.Evaluate())
	}
	if isMax {
		best := math.Inf(-1)
		for c := 0; c < g.Columns; c++ {
			tmp := g.Clone()
			if tmp.Play(c) {
				score := minimax(tmp, depth-1, false)
				if score > best {
					best = score
				}
			}
		}
		return best
	} else {
		best := math.Inf(1)
		for c := 0; c < g.Columns; c++ {
			tmp := g.Clone()
			if tmp.Play(c) {
				score := minimax(tmp, depth-1, true)
				if score < best {
					best = score
				}
			}
		}
		return best
	}
}

func localGameHandler(w http.ResponseWriter, r *http.Request) {
	player1 := r.URL.Query().Get("player1")
	player2 := r.URL.Query().Get("player2")
	level := strings.ToLower(r.URL.Query().Get("level"))

	if player1 == "" || player2 == "" {
		http.Redirect(w, r, "/local", http.StatusSeeOther)
		return
	}

	if currentGame == nil || currentGame.GameOver {
		switch level {
		case "easy":
			currentGame = logic.NewGameCustom(6, 7)
		case "normal":
			currentGame = logic.NewGameCustom(6, 9)
		case "hard":
			currentGame = logic.NewGameCustom(7, 8)
		default:
			currentGame = logic.NewGameCustom(6, 7)
		}
	}

	if r.Method == "POST" && !currentGame.GameOver {
		col := 0
		fmt.Sscanf(r.FormValue("col"), "%d", &col)
		currentGame.Play(col)
	}

	data := struct {
		*logic.Game
		Player1 string
		Player2 string
		Level   string
	}{
		Game:    currentGame,
		Player1: player1,
		Player2: player2,
		Level:   strings.Title(level),
	}

	renderTemplate(w, "local-game", data)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/menu", menuHandler)
	http.HandleFunc("/bot", botHandler)
	http.HandleFunc("/play", gameHandler)
	http.HandleFunc("/local", localHandler)
	http.HandleFunc("/local-game", localGameHandler)
	fmt.Println("✅ Serveur Power4 sur :8080")
	http.ListenAndServe(":8080", nil)
}

