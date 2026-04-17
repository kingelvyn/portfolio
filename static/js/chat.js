// static/js/chat.js
// Floating AI chat widget — no dependencies required.

(function () {
  'use strict';

  const MAX_INPUT = 500;
  const API_URL = '/api/chat';

  // ── Build DOM ──────────────────────────────────────────────────────────────

  function buildWidget() {
    // Bubble toggle button
    const bubble = el('button', { id: 'ai-chat-bubble', 'aria-label': 'Open AI assistant' });
    bubble.innerHTML = `
      <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
           stroke-linecap="round" stroke-linejoin="round" width="26" height="26">
        <path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/>
      </svg>`;

    // Panel
    const panel = el('div', { id: 'ai-chat-panel', role: 'dialog',
      'aria-label': 'AI assistant', 'aria-hidden': 'true' });

    const header = el('div', { id: 'ai-chat-header' });
    header.innerHTML = `
      <span>Ask me about Elvyn</span>
      <button id="ai-chat-close" aria-label="Close chat">✕</button>`;

    const messages = el('div', { id: 'ai-chat-messages', 'aria-live': 'polite' });

    // Greeting
    appendMessage(messages, 'assistant',
      "Hi! I'm Elvyn's portfolio assistant. Ask me about his background, skills, or projects.");

    const form = el('form', { id: 'ai-chat-form' });
    const input = el('input', {
      id: 'ai-chat-input',
      type: 'text',
      placeholder: 'Ask a question…',
      maxlength: String(MAX_INPUT),
      autocomplete: 'off',
      'aria-label': 'Your message',
    });
    const counter = el('span', { id: 'ai-chat-counter' });
    counter.textContent = `0/${MAX_INPUT}`;

    const sendBtn = el('button', { type: 'submit', id: 'ai-chat-send' });
    sendBtn.textContent = 'Send';

    const inputRow = el('div', { id: 'ai-chat-input-row' });
    inputRow.appendChild(input);
    inputRow.appendChild(sendBtn);

    form.appendChild(inputRow);
    form.appendChild(counter);
    panel.appendChild(header);
    panel.appendChild(messages);
    panel.appendChild(form);

    document.body.appendChild(bubble);
    document.body.appendChild(panel);

    injectStyles();

    // ── Events ──────────────────────────────────────────────────────────────

    bubble.addEventListener('click', () => togglePanel(panel, bubble));
    document.getElementById('ai-chat-close').addEventListener('click', () => closePanel(panel, bubble));

    input.addEventListener('input', () => {
      counter.textContent = `${input.value.length}/${MAX_INPUT}`;
    });

    form.addEventListener('submit', async (e) => {
      e.preventDefault();
      const text = input.value.trim();
      if (!text) return;

      appendMessage(messages, 'user', text);
      input.value = '';
      counter.textContent = `0/${MAX_INPUT}`;
      sendBtn.disabled = true;
      input.disabled = true;

      const typing = appendTyping(messages);

      try {
        const reply = await sendMessage(text);
        typing.remove();
        appendMessage(messages, 'assistant', reply);
      } catch (err) {
        typing.remove();
        appendMessage(messages, 'assistant',
          err.message || 'Something went wrong. Please try again.');
      } finally {
        sendBtn.disabled = false;
        input.disabled = false;
        input.focus();
      }
    });

    // Close on Escape
    document.addEventListener('keydown', (e) => {
      if (e.key === 'Escape' && panel.getAttribute('aria-hidden') === 'false') {
        closePanel(panel, bubble);
      }
    });
  }

  // ── API ───────────────────────────────────────────────────────────────────

  async function sendMessage(message) {
    const res = await fetch(API_URL, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ message }),
    });
    const data = await res.json();
    if (!res.ok) {
      throw new Error(data.error || `Server error ${res.status}`);
    }
    return data.reply;
  }

  // ── Helpers ───────────────────────────────────────────────────────────────

  function el(tag, attrs = {}) {
    const e = document.createElement(tag);
    Object.entries(attrs).forEach(([k, v]) => e.setAttribute(k, v));
    return e;
  }

  function appendMessage(container, role, text) {
    const msg = el('div', { class: `ai-msg ai-msg--${role}` });
    msg.textContent = text;
    container.appendChild(msg);
    container.scrollTop = container.scrollHeight;
    return msg;
  }

  function appendTyping(container) {
    const msg = el('div', { class: 'ai-msg ai-msg--assistant ai-msg--typing' });
    msg.innerHTML = '<span></span><span></span><span></span>';
    container.appendChild(msg);
    container.scrollTop = container.scrollHeight;
    return msg;
  }

  function togglePanel(panel, bubble) {
    const hidden = panel.getAttribute('aria-hidden') !== 'false';
    if (hidden) {
      panel.setAttribute('aria-hidden', 'false');
      panel.classList.add('ai-chat-panel--open');
      bubble.classList.add('ai-chat-bubble--open');
      document.getElementById('ai-chat-input').focus();
    } else {
      closePanel(panel, bubble);
    }
  }

  function closePanel(panel, bubble) {
    panel.setAttribute('aria-hidden', 'true');
    panel.classList.remove('ai-chat-panel--open');
    bubble.classList.remove('ai-chat-bubble--open');
  }

  // ── Styles ────────────────────────────────────────────────────────────────

  function injectStyles() {
    const css = `
      #ai-chat-bubble {
        position: fixed;
        bottom: 1.5rem;
        right: 1.5rem;
        z-index: 9999;
        width: 3.25rem;
        height: 3.25rem;
        border-radius: 50%;
        background: #2563eb;
        color: #fff;
        border: none;
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
        box-shadow: 0 4px 14px rgba(0,0,0,.25);
        transition: background 0.2s, transform 0.2s;
      }
      #ai-chat-bubble:hover { background: #1d4ed8; transform: scale(1.05); }
      #ai-chat-bubble.ai-chat-bubble--open { background: #1e40af; }

      #ai-chat-panel {
        position: fixed;
        bottom: 5.5rem;
        right: 1.5rem;
        z-index: 9998;
        width: 22rem;
        max-width: calc(100vw - 2rem);
        background: #fff;
        border-radius: 1rem;
        box-shadow: 0 8px 30px rgba(0,0,0,.18);
        display: flex;
        flex-direction: column;
        overflow: hidden;
        /* hidden by default */
        opacity: 0;
        pointer-events: none;
        transform: translateY(10px);
        transition: opacity 0.2s, transform 0.2s;
      }
      #ai-chat-panel.ai-chat-panel--open {
        opacity: 1;
        pointer-events: auto;
        transform: translateY(0);
      }

      #ai-chat-header {
        background: #2563eb;
        color: #fff;
        padding: 0.75rem 1rem;
        display: flex;
        align-items: center;
        justify-content: space-between;
        font-weight: 600;
        font-size: 0.95rem;
        font-family: system-ui, sans-serif;
      }
      #ai-chat-close {
        background: transparent;
        border: none;
        color: #fff;
        cursor: pointer;
        font-size: 1rem;
        padding: 0 0.25rem;
        opacity: 0.8;
      }
      #ai-chat-close:hover { opacity: 1; }

      #ai-chat-messages {
        flex: 1;
        overflow-y: auto;
        padding: 0.75rem 1rem;
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
        max-height: 20rem;
        font-family: system-ui, sans-serif;
        font-size: 0.875rem;
        line-height: 1.45;
      }

      .ai-msg {
        padding: 0.5rem 0.75rem;
        border-radius: 0.75rem;
        max-width: 85%;
        word-break: break-word;
      }
      .ai-msg--user {
        background: #2563eb;
        color: #fff;
        align-self: flex-end;
        border-bottom-right-radius: 0.2rem;
      }
      .ai-msg--assistant {
        background: #f1f5f9;
        color: #1e293b;
        align-self: flex-start;
        border-bottom-left-radius: 0.2rem;
      }

      /* Typing indicator */
      .ai-msg--typing {
        display: flex;
        gap: 4px;
        align-items: center;
        padding: 0.6rem 0.85rem;
      }
      .ai-msg--typing span {
        width: 7px;
        height: 7px;
        border-radius: 50%;
        background: #94a3b8;
        animation: ai-bounce 1.2s infinite;
      }
      .ai-msg--typing span:nth-child(2) { animation-delay: 0.2s; }
      .ai-msg--typing span:nth-child(3) { animation-delay: 0.4s; }
      @keyframes ai-bounce {
        0%, 80%, 100% { transform: translateY(0); }
        40%           { transform: translateY(-5px); }
      }

      #ai-chat-form {
        padding: 0.5rem 0.75rem 0.75rem;
        border-top: 1px solid #e2e8f0;
        background: #fff;
      }
      #ai-chat-input-row {
        display: flex;
        gap: 0.4rem;
      }
      #ai-chat-input {
        flex: 1;
        border: 1px solid #cbd5e1;
        border-radius: 0.5rem;
        padding: 0.45rem 0.65rem;
        font-size: 0.875rem;
        font-family: system-ui, sans-serif;
        outline: none;
        transition: border-color 0.15s;
      }
      #ai-chat-input:focus { border-color: #2563eb; }
      #ai-chat-send {
        background: #2563eb;
        color: #fff;
        border: none;
        border-radius: 0.5rem;
        padding: 0.45rem 0.85rem;
        font-size: 0.875rem;
        font-family: system-ui, sans-serif;
        cursor: pointer;
        transition: background 0.15s;
      }
      #ai-chat-send:hover:not(:disabled) { background: #1d4ed8; }
      #ai-chat-send:disabled { opacity: 0.5; cursor: not-allowed; }
      #ai-chat-counter {
        font-size: 0.72rem;
        color: #94a3b8;
        display: block;
        text-align: right;
        margin-top: 0.25rem;
        font-family: system-ui, sans-serif;
      }

      @media (prefers-color-scheme: dark) {
        #ai-chat-panel { background: #1e293b; }
        #ai-chat-form  { background: #1e293b; border-top-color: #334155; }
        .ai-msg--assistant { background: #334155; color: #e2e8f0; }
        #ai-chat-input {
          background: #0f172a; color: #e2e8f0; border-color: #475569;
        }
        #ai-chat-input:focus { border-color: #60a5fa; }
        #ai-chat-counter { color: #64748b; }
      }
    `;
    const style = document.createElement('style');
    style.textContent = css;
    document.head.appendChild(style);
  }

  // ── Init ──────────────────────────────────────────────────────────────────

  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', buildWidget);
  } else {
    buildWidget();
  }
})();