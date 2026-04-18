package chat

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// LoadKnowledge reads all .md files from dir and returns them stitched
// together as a single context block for use in the system prompt.
func LoadKnowledge(dir string) (string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", fmt.Errorf("reading knowledge dir %q: %w", dir, err)
	}

	var sb strings.Builder
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		path := filepath.Join(dir, e.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			return "", fmt.Errorf("reading %q: %w", path, err)
		}
		sb.WriteString("---\n")
		sb.WriteString(string(data))
		sb.WriteString("\n")
	}

	if sb.Len() == 0 {
		return "", fmt.Errorf("no .md files found in %q", dir)
	}
	return sb.String(), nil
}

// BuildSystemPrompt wraps the knowledge block in an instruction prompt.
func BuildSystemPrompt(knowledge string) string {
	return `You are a helpful AI assistant embedded in Elvyn's personal portfolio website.
Your job is to answer questions about Elvyn — his background, skills, and projects.

Use ONLY the information provided below. Do not invent facts.
If you do not know the answer based on the provided information, say so politely.
Keep answers concise (2-4 sentences unless more detail is clearly needed).
Do not discuss topics unrelated to Elvyn or his work.

=== KNOWLEDGE BASE ===
` + knowledge + `=== END KNOWLEDGE BASE ===`
}