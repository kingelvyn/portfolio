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
)

// Handler holds shared state for the chat endpoint.
type Handler struct {
	systemPrompt string
	apiKey       string
	limiter      *rateLimiter
}

// NewHandler creates a Handler, loading knowledge from dir.
func NewHandler(knowledgeDir string) (*Handler, error) {
	knowledge, err := LoadKnowledge(knowledgeDir)
	if err != nil {
		return nil, fmt.Errorf("loading knowledge: %w", err)
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY environment variable is not set")
	}

	return &Handler{
		systemPrompt: BuildSystemPrompt(knowledge),
		apiKey:       apiKey,
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

	reply, err := h.callGemini(req.Message)
	if err != nil {
		log.Printf("[chat] gemini error: %v", err)
		jsonError(w, "AI service unavailable, please try again later", http.StatusServiceUnavailable)
		return
	}

	resp := struct {
		Reply string `json:"reply"`
	}{Reply: reply}
	json.NewEncoder(w).Encode(resp)
}

// ── Gemini API ────────────────────────────────────────────────────────────────

type geminiRequest struct {
	SystemInstruction geminiSystem    `json:"system_instruction"`
	Contents          []geminiContent `json:"contents"`
}

type geminiSystem struct {
	Parts []geminiPart `json:"parts"`
}

type geminiContent struct {
	Role  string       `json:"role"`
	Parts []geminiPart `json:"parts"`
}

type geminiPart struct {
	Text string `json:"text"`
}

type geminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

func (h *Handler) callGemini(userMessage string) (string, error) {
	model := os.Getenv("GEMINI_MODEL")
	if model == "" {
		model = "gemini-2.5-flash"
	}

	url := fmt.Sprintf(
		"https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s",
		model, h.apiKey,
	)

	payload := geminiRequest{
		SystemInstruction: geminiSystem{
			Parts: []geminiPart{{Text: h.systemPrompt}},
		},
		Contents: []geminiContent{
			{
				Role:  "user",
				Parts: []geminiPart{{Text: userMessage}},
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("marshaling request: %w", err)
	}

	client := &http.Client{Timeout: apiTimeout}
	resp, err := client.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("calling Gemini: %w", err)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading response: %w", err)
	}

	var geminiResp geminiResponse
	if err := json.Unmarshal(raw, &geminiResp); err != nil {
		return "", fmt.Errorf("decoding response: %w", err)
	}
	if geminiResp.Error != nil {
		return "", fmt.Errorf("gemini API error: %s", geminiResp.Error.Message)
	}
	if len(geminiResp.Candidates) == 0 || len(geminiResp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("empty response from Gemini")
	}

	return strings.TrimSpace(geminiResp.Candidates[0].Content.Parts[0].Text), nil
}

func jsonError(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	fmt.Fprintf(w, `{"error":%q}`, msg)
}
