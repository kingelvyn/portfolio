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
  const heroCard = document.querySelector(".hero-content");
  const portraitPanel = document.querySelector(".portrait-panel");
  const menuItems = document.querySelectorAll(".menu a");
  const topLinks = document.querySelectorAll(".top-links a");
  const portraitMeta = document.querySelectorAll(".portrait-meta .meta-line");
  const portraitLabel = document.querySelector(".portrait-label");

  // Safety checks so the script doesn't break if one section is missing
  bootLines.forEach((line) => {
    line.textContent = "";
  });

  if (menuItems.length) {
    gsap.set(menuItems, { opacity: 0, x: -12 });
  }

  if (topLinks.length) {
    gsap.set(topLinks, { opacity: 0, y: -10 });
  }

  if (portraitPanel) {
    gsap.set(portraitPanel, { opacity: 0, x: 20 });
  }

  if (portraitLabel) {
    gsap.set(portraitLabel, { opacity: 0, y: 6 });
  }

  if (portraitMeta.length) {
    gsap.set(portraitMeta, { opacity: 0, y: 8 });
  }

  // Type first two lines
  if (bootLines[0]) {
    await typeLine(bootLines[0], bootLines[0].dataset.text || "", 32);
    await sleep(80);
  }

  if (bootLines[1]) {
    await typeLine(bootLines[1], bootLines[1].dataset.text || "", 32);
    await sleep(80);
  }

  // Bring in card while final boot line types
  if (heroCard) {
    gsap.fromTo(
      heroCard,
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
    await typeLine(bootLines[2], bootLines[2].dataset.text || "", 32);
    await sleep(80);
  }

  if (bootLines[3]) {
    await typeLine(bootLines[3], bootLines[3].dataset.text || "", 32, true);
    await sleep(120);
  }

  // Animate top links
  if (topLinks.length) {
    gsap.to(topLinks, {
      opacity: 1,
      y: 0,
      duration: 0.35,
      stagger: 0.06,
      ease: "power2.out"
    });
  }

  // Animate portrait panel slightly after the card
  if (portraitPanel) {
    gsap.to(portraitPanel, {
      opacity: 1,
      x: 0,
      duration: 0.55,
      ease: "power2.out",
      delay: 0.08
    });
  }

  if (portraitLabel) {
    gsap.to(portraitLabel, {
      opacity: 1,
      y: 0,
      duration: 0.35,
      ease: "power2.out",
      delay: 0.18
    });
  }

  if (portraitMeta.length) {
    gsap.to(portraitMeta, {
      opacity: 1,
      y: 0,
      duration: 0.3,
      stagger: 0.08,
      ease: "power2.out",
      delay: 0.24
    });
  }

  // Animate menu buttons last
  if (menuItems.length) {
    gsap.to(menuItems, {
      opacity: 1,
      x: 0,
      duration: 0.35,
      stagger: 0.08,
      ease: "power2.out",
      delay: 0.12
    });
  }

  // Main menu hover effects only
  menuItems.forEach((item) => {
    item.addEventListener("mouseenter", () => {
      gsap.to(item, {
        x: 6,
        duration: 0.15
      });

      gsap.fromTo(
        item,
        { skewX: 0 },
        {
          skewX: -4,
          duration: 0.08,
          yoyo: true,
          repeat: 1
        }
      );
    });

    item.addEventListener("mouseleave", () => {
      gsap.to(item, {
        x: 0,
        skewX: 0,
        duration: 0.15
      });
    });
  });
});