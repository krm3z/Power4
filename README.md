🧩 Power4 Web
🎯 Description

Power4 Web est un projet web développé en Golang (Go) qui permet de jouer au Puissance 4 directement depuis un navigateur.
Le jeu peut se jouer à deux en local ou contre un bot, avec 4 niveaux de difficulté.
Toute la logique est gérée côté serveur avec Go, et l’interface est rendue grâce aux templates HTML et au CSS.

🛠️ Technologies utilisées

Go (Golang) → gestion du serveur, des routes et de la logique du jeu

HTML / CSS → interface utilisateur et design du plateau

Go Templates → génération dynamique des pages

📁 Structure du projet
POWER4/
│
├── logic/                     # Logique du jeu
│   └── game.go
│
├── static/                    # Ressources statiques (images, styles, etc.)
│   ├── img/
│   │   └── background.jpg
│   └── style.css
│
├── templates/                 # Templates HTML
│   ├── bot.html               # Page pour le mode contre IA
│   ├── game.html              # Page principale du plateau de jeu
│   ├── index.html             # Page d’accueil
│   ├── local-game.html        # Page de partie locale
│   ├── local.html             # Menu du mode local
│   └── menu.html              # Menu principal
│
├── go.mod                     # Fichier du module Go
├── main.go                    # Serveur principal (routes, bot et gestion du jeu)
└── README.md

🚀 Installation et exécution
1️⃣ Cloner le projet
git clone https://github.com/krm3z/Power4.git
cd Power4

2️⃣ Lancer le serveur
go run main.go

3️⃣ Ouvrir le navigateur

Accède à :
👉 http://localhost:8080

🎮 Modes de jeu

🧍 Local 2 joueurs : deux joueurs s’affrontent sur la même machine.

🤖 Mode Bot : joue contre le bot, avec 4 niveaux de difficulté. Chaque niveau a sa propre taille de grille.

| Niveau | Grille | Comment joue le bot |
|---|---|---|
| Facile 🟢 | 6 × 7 | Il joue une colonne au hasard. |
| Normal 🟡 | 6 × 9 | Il repère souvent un coup gagnant (6 fois sur 10), sinon il joue au hasard. |
| Difficile 🔴 | 7 × 8 | Il anticipe 3 coups à l'avance avec l'algorithme minimax. |
| Mode Dieu 👑 | 7 × 8 | Il anticipe 5 coups à l'avance avec minimax. |

🧠 Règles du jeu

Chaque joueur joue à tour de rôle.

Choisis une colonne pour déposer ton jeton.

Le premier à aligner 4 jetons consécutifs (horizontal, vertical ou diagonal) gagne.

Si la grille est remplie sans vainqueur, la partie se termine sur une égalité.

💡 Fonctionnalités

✅ Interface web intuitive

✅ Mode 2 joueurs local

✅ Mode contre le bot, avec 4 niveaux de difficulté (dont deux basés sur minimax)

✅ Gestion complète du jeu côté serveur

🔜 Système de score et rejouabilité

👨‍💻 Auteur

Projet développé par Kerem
Dans le cadre du projet power4 – B1 Informatique (Ynov Campus)
