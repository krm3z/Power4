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
var currentLevel string // always lowercase: easy|normal|hard|god

// ---------- TEMPLATES ----------
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	path := fmt.Sprintf("templates/%s.html", tmpl)
	t, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, "Erreur de template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	_ = t.Execute(w, data)
}

// ---------- PAGES ----------
func indexHandler(w http.ResponseWriter, r *http.Request) { renderTemplate(w, "index", nil) }
func menuHandler(w http.ResponseWriter, r *http.Request)  { renderTemplate(w, "menu", nil) }
func botHandler(w http.ResponseWriter, r *http.Request)   { renderTemplate(w, "bot", nil) }

// ---------- GAME INIT ----------
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
		currentGame = logic.NewGameCustom(7, 8) // grille demandée
	default:
		level = "easy"
		currentGame = logic.NewGameCustom(6, 7)
	}
	currentLevel = level
	fmt.Printf("🎮 Nouvelle partie: %s (%dx%d)\n", strings.Title(level), currentGame.Rows, currentGame.Columns)
}

// ---------- GAME LOOP ----------
func gameHandler(w http.ResponseWriter, r *http.Request) {
	// 1) Choix/maintien du niveau
	level := strings.ToLower(r.URL.Query().Get("level"))
	if currentGame == nil {
		if level == "" {
			level = "easy"
		}
		initGame(level)
	} else if level != "" && level != currentLevel {
		initGame(level)
	}

	// 2) Coup joueur (toujours R)
	if r.Method == "POST" && !currentGame.GameOver && currentGame.Current == "R" {
		col := 0
		fmt.Sscanf(r.FormValue("col"), "%d", &col)
		if currentGame.Play(col) && !currentGame.GameOver {
			// Laisse une micro pause (ressenti)
			time.Sleep(350 * time.Millisecond)
			botMove(currentLevel)
		}
	}

	// 3) Données pour le template
	data := struct {
		*logic.Game
		Level    string // Pour affichage
		LevelRaw string // Pour URLs
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
		// 100% aléatoire
		for {
			c := rand.Intn(cols)
			if currentGame.Play(c) {
				break
			}
		}

	case "normal":
		// 50% logique, 50% aléatoire
		if rand.Float64() < 0.5 {
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
		// Minimax profondeur 3 avec évaluation heuristique
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
		// IA max : profondeur 5 avec agressivité
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


// ---------- MAIN ----------
func main() {
	// fichiers statiques
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/menu", menuHandler)
	http.HandleFunc("/bot", botHandler)
	http.HandleFunc("/play", gameHandler)

	fmt.Println("✅ Serveur Power4 sur :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
