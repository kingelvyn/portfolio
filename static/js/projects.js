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

  gsap.set(topLinks, { opacity: 0, y: -10 });
  gsap.set(projectCards, { opacity: 0, y: 14 });

  bootLines.forEach((line) => {
    line.textContent = "";
  });

  if (bootLines[0]) {
    await typeLine(bootLines[0], bootLines[0].dataset.text || "", 32);
    await sleep(80);
  }

  if (bootLines[1]) {
    await typeLine(bootLines[1], bootLines[1].dataset.text || "", 32);
    await sleep(80);
  }

  if (projectPanel) {
    gsap.fromTo(
      projectPanel,
      { opacity: 0, y: 24, scale: 0.985 },
      {
        opacity: 1,
        y: 0,
        scale: 1,
        duration: 0.8,
        ease: "power3.out"
      }
    );
  }

  if (bootLines[2]) {
    await typeLine(bootLines[2], bootLines[2].dataset.text || "", 32, true);
    await sleep(100);
  }

  gsap.to(topLinks, {
    opacity: 1,
    y: 0,
    duration: 0.35,
    stagger: 0.06,
    ease: "power2.out"
  });

  gsap.to(projectCards, {
    opacity: 1,
    y: 0,
    duration: 0.35,
    stagger: 0.05,
    ease: "power2.out",
    delay: 0.1
  });
});