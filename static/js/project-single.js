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

window.addEventListener("DOMContentLoaded", async () => {
  const bootLines = document.querySelectorAll(".boot-line");
  const topLinks = document.querySelectorAll(".top-links a");
  const detailPanel = document.querySelector(".project-detail-panel");
  const detailHeader = document.querySelector(".detail-header");
  const markdownNodes = document.querySelectorAll(".project-markdown > *");
  const archiveButton = document.querySelector(".detail-btn");
  const panelTag = document.querySelector(".panel-tag");
  const statusBadge = document.querySelector(".detail-status-badge");

  bootLines.forEach((line) => {
    line.textContent = "";
  });

  if (topLinks.length) {
    gsap.set(topLinks, { autoAlpha: 0, y: -10 });
  }

  if (detailPanel) {
    gsap.set(detailPanel, { autoAlpha: 0, y: 24, scale: 0.985 });
  }

  if (detailHeader) {
    gsap.set(detailHeader, { autoAlpha: 0, y: 12 });
  }

  if (markdownNodes.length) {
    gsap.set(markdownNodes, { autoAlpha: 0, y: 14 });
  }

  if (panelTag) {
    gsap.set(panelTag, { autoAlpha: 0, x: 12 });
  }

  if (statusBadge) {
    gsap.set(statusBadge, { autoAlpha: 0, scale: 0.85 });
  }

  if (bootLines[0]) {
    await typeLine(bootLines[0], bootLines[0].dataset.text || "", 32, true);
    await sleep(60);
  }

  const tl = gsap.timeline();

  if (topLinks.length) {
    tl.to(topLinks, {
      autoAlpha: 1,
      y: 0,
      duration: 0.35,
      stagger: 0.06,
      ease: "power2.out"
    });
  }

  if (detailPanel) {
    tl.to(detailPanel, {
      autoAlpha: 1,
      y: 0,
      scale: 1,
      duration: 0.7,
      ease: "power3.out"
    }, "-=0.15");
  }

  if (detailHeader) {
    tl.to(detailHeader, {
      autoAlpha: 1,
      y: 0,
      duration: 0.4,
      ease: "power2.out"
    }, "<0.08");
  }

  if (panelTag) {
    tl.to(panelTag, {
      autoAlpha: 1,
      x: 0,
      duration: 0.3,
      ease: "power2.out"
    }, "<0.02");
  }

  if (statusBadge) {
    tl.to(statusBadge, {
      autoAlpha: 1,
      scale: 1,
      duration: 0.25,
      ease: "back.out(1.6)"
    }, "<0.02");
  }

  if (markdownNodes.length) {
    tl.to(markdownNodes, {
      autoAlpha: 1,
      y: 0,
      duration: 0.28,
      stagger: 0.03,
      ease: "power2.out"
    }, "<0.08");
  }

  if (archiveButton) {
    archiveButton.addEventListener("mouseenter", () => {
      gsap.to(archiveButton, { x: 4, duration: 0.15 });
    });

    archiveButton.addEventListener("mouseleave", () => {
      gsap.to(archiveButton, { x: 0, duration: 0.15 });
    });
  }
});