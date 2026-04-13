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
  const detailPanel = document.querySelector(".project-detail-panel");
  const detailHeader = document.querySelector(".detail-header");
  const markdownNodes = document.querySelectorAll(".project-markdown > *");
  const archiveButton = document.querySelector(".detail-btn");

  bootLines.forEach((line) => {
    line.textContent = "";
  });

  if (topLinks.length) {
    gsap.set(topLinks, { opacity: 0, y: -10 });
  }

  if (detailHeader) {
    gsap.set(detailHeader, { opacity: 0, y: 12 });
  }

  if (markdownNodes.length) {
    gsap.set(markdownNodes, { opacity: 0, y: 14 });
  }

  if (bootLines[0]) {
    await typeLine(bootLines[0], bootLines[0].dataset.text || "", 32);
    await sleep(80);
  }

  if (bootLines[1]) {
    await typeLine(bootLines[1], bootLines[1].dataset.text || "", 32);
    await sleep(80);
  }

  if (detailPanel) {
    gsap.fromTo(
      detailPanel,
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

  if (topLinks.length) {
    gsap.to(topLinks, {
      opacity: 1,
      y: 0,
      duration: 0.35,
      stagger: 0.06,
      ease: "power2.out"
    });
  }

  if (detailHeader) {
    gsap.to(detailHeader, {
      opacity: 1,
      y: 0,
      duration: 0.4,
      ease: "power2.out"
    });
  }

  if (markdownNodes.length) {
    gsap.to(markdownNodes, {
      opacity: 1,
      y: 0,
      duration: 0.3,
      stagger: 0.03,
      ease: "power2.out",
      delay: 0.1
    });
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