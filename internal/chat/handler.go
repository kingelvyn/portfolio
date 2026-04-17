package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	maxInputLength = 500
	apiTimeout     = 30 * time.Second
	groqAPIURL     = "https://api.groq.com/openai/v1/chat/completions"
)

// Handler holds shared state for the chat endpoint.
type Handler struct {
	systemPrompt string
	apiKey       string
	model        string
	limiter      *rateLimiter
}

// NewHandler creates a Handler, loading knowledge from dir.
func NewHandler(knowledgeDir string) (*Handler, error) {
	knowledge, err := LoadKnowledge(knowledgeDir)
	if err != nil {
		return nil, fmt.Errorf("loading knowledge: %w", err)
	}

	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("GROQ_API_KEY environment variable is not set")
	}

	model := os.Getenv("GROQ_MODEL")
	if model == "" {
		model = "llama3-8b-8192" // free, fast, good quality
	}

	return &Handler{
		systemPrompt: BuildSystemPrompt(knowledge),
		apiKey:       apiKey,
		model:        model,
		limiter:      newRateLimiter(10, time.Minute),
	}, nil
}

// ServeHTTP implements http.Handler.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.limiter.Middleware(http.HandlerFunc(h.handle)).ServeHTTP(w, r)
}

func (h *Handler) handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var req struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, "invalid request body", http.StatusBadRequest)
		return
	}
	req.Message = strings.TrimSpace(req.Message)
	if req.Message == "" {
		jsonError(w, "message is required", http.StatusBadRequest)
		return
	}
	if len(req.Message) > maxInputLength {
		jsonError(w, fmt.Sprintf("message too long (max %d characters)", maxInputLength), http.StatusBadRequest)
		return
	}

	log.Printf("[chat] ip=%s msg_len=%d", realIP(r), len(req.Message))

	reply, err := h.callGroq(req.Message)
	if err != nil {
		log.Printf("[chat] groq error: %v", err)
		jsonError(w, "AI service unavailable, please try again later", http.StatusServiceUnavailable)
		return
	}

	resp := struct {
		Reply string `json:"reply"`
	}{Reply: reply}
	json.NewEncoder(w).Encode(resp)
}

// ── Groq / OpenAI-compatible API ─────────────────────────────────────────────

type chatRequest struct {
	Model    string        `json:"model"`
	Messages []chatMessage `json:"messages"`
}

type chatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatResponse struct {
	Choices []struct {
		Message chatMessage `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

func (h *Handler) callGroq(userMessage string) (string, error) {
	payload := chatRequest{
		Model: h.model,
		Messages: []chatMessage{
			{Role: "system", Content: h.systemPrompt},
			{Role: "user", Content: userMessage},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("marshaling request: %w", err)
	}

	httpReq, err := http.NewRequest(http.MethodPost, groqAPIURL, bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("building request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+h.apiKey)

	client := &http.Client{Timeout: apiTimeout}
	resp, err := client.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("calling Groq: %w", err)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading response: %w", err)
	}

	var groqResp chatResponse
	if err := json.Unmarshal(raw, &groqResp); err != nil {
		return "", fmt.Errorf("decoding response: %w", err)
	}
	if groqResp.Error != nil {
		return "", fmt.Errorf("groq API error: %s", groqResp.Error.Message)
	}
	if len(groqResp.Choices) == 0 {
		return "", fmt.Errorf("no choices in groq response")
	}

	return strings.TrimSpace(groqResp.Choices[0].Message.Content), nil
}

func jsonError(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	fmt.Fprintf(w, `{"error":%q}`, msg)
}