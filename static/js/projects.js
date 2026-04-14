function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

async function typeLine(element, text, speed = 32, keepCursor = false) {
  element.textContent = "";
  element.classList.add("typing");

  for (let i = 0; i < text.length; i++) {
    element.textContent += text[i];
    await sleep(speed);
  }

  await sleep(180);

  if (!keepCursor) {
    element.classList.remove("typing");
  }
}

window.addEventListener("load", async () => {
  const bootLines = document.querySelectorAll(".boot-line");
  const topLinks = document.querySelectorAll(".top-links a");
  const projectPanel = document.querySelector(".projects-panel");
  const projectCards = document.querySelectorAll(".project-card");

  bootLines.forEach((line) => {
    line.textContent = "";
  });

  if (topLinks.length) {
    gsap.set(topLinks, { opacity: 0, y: -10 });
  }

  if (projectPanel) {
    gsap.set(projectPanel, { opacity: 0, y: 24, scale: 0.985 });
  }

  if (projectCards.length) {
    gsap.set(projectCards, { opacity: 0, y: 14 });
  }

  if (bootLines[0]) {
    await typeLine(bootLines[0], bootLines[0].dataset.text || "", 32);
    await sleep(70);
  }

  const tl = gsap.timeline();

  if (topLinks.length) {
    tl.to(topLinks, {
      opacity: 1,
      y: 0,
      duration: 0.35,
      stagger: 0.06,
      ease: "power2.out"
    });
  }

  if (projectPanel) {
    tl.to(projectPanel, {
      opacity: 1,
      y: 0,
      scale: 1,
      duration: 0.7,
      ease: "power3.out"
    }, "-=0.15");
  }

  if (projectCards.length) {
    tl.to(projectCards, {
      opacity: 1,
      y: 0,
      duration: 0.35,
      stagger: 0.05,
      ease: "power2.out"
    }, "<0.08");
  }
});