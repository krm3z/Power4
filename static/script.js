// --- Animation de chute des pions ---
document.addEventListener("DOMContentLoaded", () => {
  const grid = document.querySelector(".grid");
  const rows = document.querySelectorAll(".row");
  const layer = document.querySelector(".animation-layer");

  if (!grid || !rows.length || !layer) return;

  let isAnimating = false;

  document.querySelectorAll(".cell").forEach((cell) => {
    cell.addEventListener("click", (e) => {
      const form = e.target.closest("form");
      if (!form || isAnimating) {
        e.preventDefault();
        return;
      }

      e.preventDefault();
      isAnimating = true;

      const col = parseInt(form.querySelector("input[name='col']").value);
      if (Number.isNaN(col)) {
        isAnimating = false;
        return;
      }

      // Cherche la première case vide depuis le bas
      let targetRow = null;
      for (let i = rows.length - 1; i >= 0; i--) {
        const disc = rows[i].children[col]?.querySelector(".disc");
        if (disc && disc.classList.contains("empty")) {
          targetRow = i;
          break;
        }
      }

      if (targetRow === null) {
        isAnimating = false;
        return;
      }

      // Coordonnées
      const gridRect = grid.getBoundingClientRect();
      const targetCell = rows[targetRow].children[col].querySelector(".cell");
      const rect = targetCell.getBoundingClientRect();

      const subtitle = document.querySelector(".subtitle");
      const color = subtitle && subtitle.textContent.includes("🔴") ? "red" : "yellow";

      // Crée un pion animé
      const disc = document.createElement("div");
      disc.className = `disc ${color}`;
      disc.style.position = "absolute";
      disc.style.left = `${rect.left - gridRect.left + 10}px`;
      disc.style.top = `-100px`;
      disc.style.width = "70px";
      disc.style.height = "70px";
      disc.style.borderRadius = "50%";
      disc.style.animation = "drop 0.5s ease-in forwards";
      disc.style.setProperty("--targetY", `${rect.top - gridRect.top}px`);
      layer.appendChild(disc);

      // Quand la chute est finie → soumission du formulaire
      disc.addEventListener(
        "animationend",
        () => {
          form.submit();
        },
        { once: true }
      );
    });
  });
});

// --- Gestion de la pop-up de victoire ---
window.addEventListener("load", () => {
  const victoryScreen = document.querySelector(".victory-screen");
  if (!victoryScreen) return;

  // Autorise les clics
  victoryScreen.style.pointerEvents = "auto";
  const victoryCard = victoryScreen.querySelector(".victory-card");
  if (victoryCard) victoryCard.style.pointerEvents = "auto";


  });
  