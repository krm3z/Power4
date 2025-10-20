🧩 Power4 Web
🎯 Description

Power4 Web est un projet web développé en Golang (Go) qui permet de jouer au Puissance 4 directement depuis un navigateur.
Le jeu peut se jouer à deux en local ou contre un bot.
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
├── Power4/                    # Dossier du projet (README.md)
│   └── README.md
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
└── main.go                    # Serveur principal (routes + gestion du jeu)

🚀 Installation et exécution
1️⃣ Cloner le projet
git clone https://github.com/ton-utilisateur/power4-web.git
cd power4-web

2️⃣ Initialiser le module Go
go mod init power4-web
go mod tidy

3️⃣ Lancer le serveur
go run main.go

4️⃣ Ouvrir le navigateur

Accède à :
👉 http://localhost:8080

🎮 Modes de jeu

🧍 Local 2 joueurs : deux joueurs s’affrontent sur la même machine.

🤖 Mode Bot : joue contre une intelligence artificielle (niveau basique).

🧠 Règles du jeu

Chaque joueur joue à tour de rôle.

Choisis une colonne pour déposer ton jeton.

Le premier à aligner 4 jetons consécutifs (horizontal, vertical ou diagonal) gagne.

Si la grille est remplie sans vainqueur, la partie se termine sur une égalité.

💡 Fonctionnalités

✅ Interface web intuitive

✅ Mode 2 joueurs local

✅ Gestion complète du jeu côté serveur

🔜 Amélioration de l’IA du bot

🔜 Système de score et rejouabilité

👨‍💻 Auteur

Projet développé par Kerem
Dans le cadre du projet power4 – B1 Informatique (Ynov Campus)
