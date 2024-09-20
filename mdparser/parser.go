package mdparser

import (
	"fmt"
	"regexp"
	"strings"
)

func Parse(md string) string {
		var value string
	for _, v := range strings.Split(md, "\n") {
		if len(strings.TrimLeft(v, " ")) != 0 {
			switch strings.TrimLeft(v, " ")[0] {
			case '#':
				value += checkTitle(v)
			case '_':
				value += checkLine(v, "_")
			case '-':
				value += checkLine(v, "-")
			case '*':
				value += checkLine(v, "*")
			case '>':
				value += "<blockquote>" + v[1:] + "</blockquote>"
			default:
				value += "<p>" + v + "</p>"
			}

			value = replaceAll(value, `\*\*\*(.*?)\*\*\*`, "<strong><em>$1</em></strong>")
			value = replaceAll(value, `___(.*?)___`, "<strong><em>$1</em></strong>")
			value = replaceAll(value, `\*\*(.*?)\*\*`, "<strong>$1</strong>")
			value = replaceAll(value, `__(.*?)__`, "<strong>$1</strong>")
			value = replaceAll(value, `\*(.*?)\*`, "<em>$1</em>")
			value = replaceAll(value, `_(.*?)_`, "<em>$1</em>")
			value = replaceAll(value, `\[(.*?)\]\((.*?)\)`, "<a href=\"$2\">$1</a>")

			value += "\n"
		}
	}
	return value
}

func checkTitle(line string) string {
	splits := strings.Split(strings.TrimLeft(line, " "), " ")
	h := len(splits[0])

	if h > 6 {
		message := fmt.Sprintf("%s%s", strings.Repeat("#", h-6), strings.Join(splits[1:], " "))
		return fmt.Sprintf("<h6>%s</h6>", message)
	} else {
		return fmt.Sprintf("<h%d>%s</h%d>", h, strings.Join(splits[1:], " "), h)
	}
}

func checkLine(line, delimiter string) string {
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
