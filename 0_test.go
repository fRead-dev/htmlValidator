package htmlValidator

import (
	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/html"
	"strings"
	"testing"
)

func TestHtml(t *testing.T) {

	text := "" +
		"<p>простой абзац</p>" +
		"<p left>лево</p>" +
		"<p right>право</p>" +
		"<p center>центр</p>" +
		"<hr/>" +
		"<hr>" +
		"<b>жирный</b>" +
		"<i>наклонный</i>" +
		"<u>подчеркнутый</u>" +
		"<s>зачеркнутый</s>" +
		"<q>цитата</q>" +
		"<sub>в низ мелкий текст</sub>" +
		"<sup>в верх мелкий текст</sup>"

	// Создаем новый парсер для парсинга HTML
	parser := html.NewLexer(parse.NewInput(strings.NewReader(text)))

	textBuf := ""
	//tokenBuf := ""

	for {
		typeToken, data := parser.Next()

		switch typeToken {
		case html.StartTagCloseToken:
			continue

		case html.StartTagToken:
			t.Log("START", string(data), string(parser.AttrKey()), string(parser.AttrVal()))

		case html.EndTagToken:
			t.Log("END", string(data), string(parser.AttrKey()), string(parser.AttrVal()))

		case html.TextToken:
			textBuf += string(data)

		case html.ErrorToken:
			t.Log(textBuf)
			return

		default:
			t.Log(typeToken, string(data))
		}
	}
}
