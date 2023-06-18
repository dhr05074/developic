package str

import "testing"

func TestExtractCodeBlocksFromMarkdown(t *testing.T) {
	markdown := "" +
		"# Hello, world!\n" +
		"\n" +
		"```go\n" +
		"package main\n" +
		"```"
	codeBlocks := ExtractCodeBlocksFromMarkdown(markdown)
	if len(codeBlocks) != 1 {
		t.Errorf("expected 1 code block, got %d", len(codeBlocks))
	}
	if codeBlocks[0] != "package main" {
		t.Errorf("expected code block to be 'package main', got '%s'", codeBlocks[0])
	}
}
