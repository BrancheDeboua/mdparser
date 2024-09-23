package mdparser

import (
	"testing"
)

// Helper function to compare expected vs. actual output
func runTest(t *testing.T, input, expected string) {
	result := Parse(input)
	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s\n", expected, result)
	}
}

func TestParseHeaders(t *testing.T) {
	// Test for headers from H1 to H6
	runTest(t, "# Header 1", "<h1>Header 1</h1>\n")
	runTest(t, "## Header 2", "<h2>Header 2</h2>\n")
	runTest(t, "### Header 3", "<h3>Header 3</h3>\n")
	runTest(t, "###### Header 6", "<h6>Header 6</h6>\n")
}

func TestParseBoldAndItalic(t *testing.T) {
	// Test for bold and italic parsing
	runTest(t, "**Bold Text**", "<p><strong>Bold Text</strong></p>\n")
	runTest(t, "*Italic Text*", "<p><em>Italic Text</em></p>\n")
	runTest(t, "***Bold Italic***", "<p><strong><em>Bold Italic</em></strong></p>\n")
}

func TestParseBlockquote(t *testing.T) {
	// Test for blockquote
	runTest(t, "> This is a blockquote", "<blockquote> This is a blockquote</blockquote>\n")
}

func TestParseLinks(t *testing.T) {
	// Test for parsing links
	runTest(t, "[Go](https://golang.org)", "<p><a href=\"https://golang.org\">Go</a></p>\n")
}

func TestParseHorizontalRules(t *testing.T) {
	// Test for horizontal lines with different delimiters
	runTest(t, "---", "<div style=\"width: 100%; border-bottom: 2px solid #333;\"></div>\n")
	runTest(t, "***", "<div style=\"width: 100%; border-bottom: 2px solid #333;\"></div>\n")
	runTest(t, "___", "<div style=\"width: 100%; border-bottom: 2px solid #333;\"></div>\n")
}

func TestParseParagraph(t *testing.T) {
	// Test for simple paragraphs
	runTest(t, "This is a paragraph.", "<p>This is a paragraph.</p>\n")
}

func TestParseCodeBlock(t *testing.T) {
	// Test for code blocks
	runTest(t, "```go\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n}\n```", "<code class=\"language-go\"><pre>\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n}\n</pre></code>\n")
}
