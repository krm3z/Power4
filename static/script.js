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

      // Trouve la colonne
      const col = parseInt(form.querySelector("input[name='col']").value);
      if (Number.isNaN(col)) return;

      // Trouve la première case vide en partant du bas
      let targetRow = null;
      for (let i = rows.length - 1; i >= 0; i--) {
        const btn = rows[i].children[col]?.querySelector(".cell");
        if (btn && btn.textContent.trim() === "⚪") {
          targetRow = i;
          break;
        }
      }
      if (targetRow === null) return;

      // Coordonnées
      const gridRect = grid.getBoundingClientRect();
      const targetCell = rows[targetRow].children[col].querySelector(".cell");
      const rect = targetCell.getBoundingClientRect();

      const disc = document.createElement("div");
      disc.className = "disc red";
      disc.style.left = `${rect.left - gridRect.left}px`;
      disc.style.setProperty("--targetY", `${rect.top - gridRect.top}px`);
      disc.style.animation = "drop 0.45s ease-in forwards";
      layer.appendChild(disc);

      // Quand l’animation se termine → joue le coup
      disc.addEventListener("animationend", () => {
        form.submit();
      }, { once: true });
    });
  });
});
