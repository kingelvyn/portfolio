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
  const menuItems = document.querySelectorAll(".menu a, .menu button");
  const topLinks = document.querySelectorAll(".top-links a");
  const portraitMeta = document.querySelectorAll(".portrait-meta .meta-line");
  const portraitLabel = document.querySelector(".portrait-label");

  const profileOpenBtn = document.querySelector("[data-profile-open]");
  const profileModal = document.querySelector("[data-profile-modal]");
  const profileCloseBtns = document.querySelectorAll("[data-profile-close]");

  bootLines.forEach((line) => {
    line.textContent = "";
  });

  if (menuItems.length) {
    gsap.set(menuItems, { opacity: 0, x: -12 });
  }

  if (topLinks.length) {
    gsap.set(topLinks, { opacity: 0, y: -10 });
  }

  if (heroCard) {
    gsap.set(heroCard, { opacity: 0, y: 24, scale: 0.985 });
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

  if (bootLines[0]) {
    await typeLine(bootLines[0], bootLines[0].dataset.text || "", 32);
    await sleep(70);
  }

  if (bootLines[1]) {
    await typeLine(bootLines[1], bootLines[1].dataset.text || "", 32);
    await sleep(70);
  }

  if (bootLines[2]) {
    await typeLine(bootLines[2], bootLines[2].dataset.text || "", 32);
    await sleep(70);
  }

  if (bootLines[3]) {
    await typeLine(bootLines[3], bootLines[3].dataset.text || "", 32, true);
    await sleep(60);
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

  if (heroCard) {
    tl.to(heroCard, {
      opacity: 1,
      y: 0,
      scale: 1,
      duration: 0.75,
      ease: "power3.out"
    }, "-=0.15");
  }

  if (portraitPanel) {
    tl.to(portraitPanel, {
      opacity: 1,
      x: 0,
      duration: 0.5,
      ease: "power2.out"
    }, "<0.08");
  }

  if (portraitLabel) {
    tl.to(portraitLabel, {
      opacity: 1,
      y: 0,
      duration: 0.3,
      ease: "power2.out"
    }, "<0.08");
  }

  if (portraitMeta.length) {
    tl.to(portraitMeta, {
      opacity: 1,
      y: 0,
      duration: 0.28,
      stagger: 0.07,
      ease: "power2.out"
    }, "<0.05");
  }

  if (menuItems.length) {
    tl.to(menuItems, {
      opacity: 1,
      x: 0,
      duration: 0.32,
      stagger: 0.07,
      ease: "power2.out"
    }, "<0.02");
  }

  if (profileOpenBtn && profileModal) {
    profileOpenBtn.addEventListener("click", () => {
      profileModal.classList.add("is-open");
      profileModal.setAttribute("aria-hidden", "false");
      document.body.style.overflow = "hidden";
    });
  
    profileCloseBtns.forEach((btn) => {
      btn.addEventListener("click", () => {
        profileModal.classList.remove("is-open");
        profileModal.setAttribute("aria-hidden", "true");
        document.body.style.overflow = "";
      });
    });
  
    document.addEventListener("keydown", (event) => {
      if (event.key === "Escape" && profileModal.classList.contains("is-open")) {
        profileModal.classList.remove("is-open");
        profileModal.setAttribute("aria-hidden", "true");
        document.body.style.overflow = "";
      }
    });
  }

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