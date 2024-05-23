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
		"<p left=\"\">лево</p>" +
		"<p right=''>право</p>" +
		"<p center>центр</p>" +
		"<hr/>" +
		"<hr>" +
		"<b>жирный</b>" +
		"<i>наклонный</i>" +
		"<u>подчеркнутый</u>" +
		"<s>зачеркнутый</s>" +
		"<q>цитата</q>" +
		"<sub>в низ мелкий текст</sub>" +
		"<sup>в верх мелкий текст</sup>" +
		"<div>косячный блок</div>" +
		"<img src=''/>"

	// Создаем новый парсер для парсинга HTML
	parser := html.NewLexer(parse.NewInput(strings.NewReader(text)))

	textBuf := ""

	errorTags := map[string]uint16{}
	errorTags = make(map[string]uint16)

	waitParagraph := false

	for {
		typeToken, data := parser.Next()

		switch typeToken {
		case html.StartTagCloseToken, html.StartTagVoidToken:
			continue

		case html.AttributeToken:
			if waitParagraph { //обрабатываем только ожидающие атрибуты
				key, isValid := isValidParagraphAttribute(parser.AttrKey())

				if isValid {
					textBuf += "<" + TagParagraph + " " + key + ">"
				} else {
					textBuf += "<" + TagParagraph + ">"
				}

				waitParagraph = false
			}
			t.Log("ATTRIBUTE", string(parser.AttrKey()), string(parser.AttrVal()))

		case html.StartTagToken:
			tag, isValid, isParagraph := IsValidTag(parser.AttrKey())
			if isValid { //если тег валиден
				if !isParagraph { //и если это не параграф
					textBuf += "<" + tag + ">"

				} else { //если таки параграф
					waitParagraph = true
				}

			} else { // Обрабатываем ошибку с неизвестным тегом
				errorTags[tag] += 1
			}

		case html.EndTagToken:
			waitParagraph = false
			tag, isValid, _ := IsValidTag(parser.AttrKey())
			if isValid { //если тег валиден
				textBuf += "</" + tag + ">"
			}

		case html.TextToken:
			waitParagraph = false
			textBuf += string(data)

		case html.ErrorToken:
			t.Log(textBuf)
			t.Log(errorTags)
			return

		default:
			t.Log("DEF", typeToken, string(data))
		}
	}
}
