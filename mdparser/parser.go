package mdparser

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

type markdownParser struct {
	value     string
	codeblock bool
}

func newMarkdownParser() *markdownParser {
	return &markdownParser{}
}

func (mp *markdownParser) parse(md string) string {
	lines := strings.Split(md, "\n")
	for _, line := range lines {
		if mp.isInBlockElement() {
			line = mp.closeBlockElement(line)
		}
		if len(strings.TrimLeft(line, " ")) != 0 {
			mp.parseBlockElement(line)
		}
	}

	mp.parseInlineElement()
	return mp.value
}

func (mp *markdownParser) parseInlineElement() {
	mp.value = replaceAll(mp.value, "`(.*?)`", "<code>$1</code>")
	mp.value = replaceAll(mp.value, `\*\*\*(.*?)\*\*\*`, "<strong><em>$1</em></strong>")
	mp.value = replaceAll(mp.value, `___(.*?)___`, "<strong><em>$1</em></strong>")
	mp.value = replaceAll(mp.value, `\*\*(.*?)\*\*`, "<strong>$1</strong>")
	mp.value = replaceAll(mp.value, `__(.*?)__`, "<strong>$1</strong>")
	mp.value = replaceAll(mp.value, `\*(.*?)\*`, "<em>$1</em>")
	mp.value = replaceAll(mp.value, `_(.*?)_`, "<em>$1</em>")
	mp.value = replaceAll(mp.value, `!\[([^\]]*)\]\(([^ ]+)(?: "([^"]*)")?\)`, "<img src=\"$2\" alt=\"$1\" title=\"$3\">")
	mp.value = replaceAll(mp.value, `\[(.*?)\]\((\S*?)\)`, "<a href=\"$2\">$1</a>")
	mp.value = replaceAll(mp.value, `\[(.*?)\]\((.*?)(?: "(.*?)")?\)`, "<a href=\"$2\" title=\"$3\">$1</a>")
}

func (mp *markdownParser) parseBlockElement(line string) {
	trimmed := strings.TrimLeft(line, " ")
	switch trimmed[0] {
	case '#':
		mp.checkTitle(line)
	case '_':
		mp.checkLine(line, "_")
	case '-':
		mp.checkLine(line, "-")
	case '*':
		mp.checkLine(line, "*")
	case '>':
		mp.value += "<blockquote>" + line[1:] + "</blockquote>"
	case '`':
		mp.checkCodeBlock(line)
	default:
		if !mp.isInBlockElement() {
			mp.value += "<p>" + line + "</p>"
		} else {
			mp.value += line
		}
	}

	mp.value += "\n"
}

func (mp *markdownParser) closeBlockElement(line string) string {
	if mp.codeblock && (len(line) >= 3 && strings.TrimLeft(line, " ")[:3] == "```") {
		mp.value += "</pre></code>\n"
		mp.codeblock = false
		line = ""
	}
	return line
}

func (mp *markdownParser) checkTitle(line string) {
	splits := strings.Split(strings.TrimLeft(line, " "), " ")
	h := len(splits[0])

	if h > 6 {
		message := fmt.Sprintf("%s%s", strings.Repeat("#", h-6), strings.Join(splits[1:], " "))
		mp.value += fmt.Sprintf("<h6>%s</h6>", message)
	} else {
		mp.value += fmt.Sprintf("<h%d>%s</h%d>", h, strings.Join(splits[1:], " "), h)
	}
}

func (mp *markdownParser) checkLine(line, delimiter string) {
	splits := strings.Fields(line)
	h := len(splits[0])
	if (len(splits) > 1 || len(strings.ReplaceAll(strings.Join(splits, ""), delimiter, "")) > 0) || h < 3 {
		mp.value += "<p>" + line + "</p>"
	} else {
		mp.value += `<div style="width: 100%; border-bottom: 2px solid #333;"></div>`
	}
}

func (mp *markdownParser) checkCodeBlock(line string) {
	if len(line) >= 3 && strings.TrimLeft(line[:3], " ") == "```" {
		if !mp.codeblock {
			mp.value += "<code class=\"language-" + line[3:] + "\"><pre>"
		} else {
			mp.value += "</pre></code>"
		}
		mp.codeblock = !mp.codeblock
	} else {
		mp.value += "<p>" + line + "</p>"
	}
}

func (mp *markdownParser) isInBlockElement() bool {
	return slices.Contains([]bool{mp.codeblock}, true)
}

func replaceAll(s, pattern, repl string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(s, repl)
}

func Parse(md string) string {
	mp := newMarkdownParser()
	return mp.parse(md)
}
