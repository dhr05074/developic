package str

import (
	"regexp"
	"strings"
)

func ExtractCodeBlocksFromMarkdown(markdown string) []string {
	// Regular expression to match code blocks in Markdown
	regex := regexp.MustCompile("`{3}[\n\r]*([^\n\r]*)[\n\r]*([^`]*)`{3}")

	// Find all matches of code blocks
	matches := regex.FindAllStringSubmatch(markdown, -1)

	// Extract code blocks from the matches
	codeBlocks := make([]string, len(matches))
	for i, match := range matches {
		codeBlocks[i] = strings.TrimSpace(match[2])
	}

	return codeBlocks
}
