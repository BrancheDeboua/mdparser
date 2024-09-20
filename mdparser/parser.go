package mdparser

import (
	"fmt"
	"regexp"
	"strings"
)

type markdownParser struct {
	lines []string
}

func newMarkdownParser() *markdownParser {
	return &markdownParser{}
}

func (m *markdownParser) addLine(line string) {
	m.lines = append(m.lines, line)
}

func (m *markdownParser) parse() string {
	var value string
	for _, v := range m.lines {
		if len(strings.TrimLeft(v, " ")) != 0 {
			switch strings.TrimLeft(v, " ")[0] {
			case '#':
				value += m.checkTitle(v)
			case '_':
				value += m.checkLine(v, "_")
			case '-':
				value += m.checkLine(v, "-")
			case '*':
				value += m.checkLine(v, "*")
			case '>':
				value += "<blockquote>" + v[1:] + "</blockquote>"
			default:
				value += "<p>" + v + "</p>"
			}

			value = replaceAll(value, `\*\*(.*?)\*\*`, "<b>$1</b>")
			value = replaceAll(value, `__(.*?)__`, "<b>$1</b>")
			value = replaceAll(value, `\*(.*?)\*`, "<i>$1</i>")
			value = replaceAll(value, `_(.*?)_`, "<i>$1</i>")
			value = replaceAll(value, `\[(.*?)\]\((.*?)\)`, "<a href=\"$2\">$1</a>")

			value += "\n"
		}
	}
	return value
}

func (m *markdownParser) checkTitle(line string) string {
	splits := strings.Split(strings.TrimLeft(line, " "), " ")
	h := len(splits[0])

	if h > 6 {
		message := fmt.Sprintf("%s%s", strings.Repeat("#", h-6), strings.Join(splits[1:], " "))
		return fmt.Sprintf("<h6>%s</h6>", message)
	} else {
		return fmt.Sprintf("<h%d>%s</h%d>", h, strings.Join(splits[1:], " "), h)
	}
}

func (m *markdownParser) checkLine(line, delimiter string) string {
	splits := strings.Fields(line)
	h := len(splits[0])
	if (len(splits) > 1 || len(strings.ReplaceAll(strings.Join(splits, ""), delimiter, "")) > 0) || h < 3 {
		return "<p>" + line + "</p>"
	}

	return `<div style="width: 100%; border-bottom: 2px solid #333;"></div>`
}

func replaceAll(s, pattern, repl string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(s, repl)
}

func Parse(md string) string {
	mp := newMarkdownParser()

	for _, v := range strings.Split(md, "\n") {
		mp.addLine(v)
	}

	return mp.parse()
}
