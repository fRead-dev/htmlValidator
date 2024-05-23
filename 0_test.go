package htmlValidator

import (
	"fmt"
	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/html"
	"strings"
	"testing"
)

func printError(t *testing.T, text string) {
	t.Error(fmt.Errorf("%s", text))
}
func printFatal(t *testing.T, text string, err error) {
	t.Fatalf(text, err)
}

//###################################################################//

func TestMap(t *testing.T) {
	type StructObj struct {
		tag         string
		isTag       bool
		isParagraph bool
	}
	dict := map[string]StructObj{
		"ValidParagraph": {TagParagraph, true, true},
		"ValidTag":       {TagDelimiter, true, false},
		"InvalidTag":     {"div", false, false},
	}

	for key, obj := range dict {
		t.Run(key, func(t *testing.T) {
			tag, valid, paragraph := IsValidTag([]byte(obj.tag))
			if tag != obj.tag {
				printError(t, "Invalid Paragraph stringify")
			}
			if valid != obj.isTag {
				printError(t, "Invalid switch tag")
			}
			if paragraph != obj.isParagraph {
				printError(t, "Invalid switch paragraph")
			}
		})
	}

}

func TestTextTransform(t *testing.T) {
	transformObj := TextTransform()
	newText := ""
	oldText := ""

	//		Горизонтальная линия разделения
	t.Run("TagDelimiter", func(t *testing.T) {

	})

	//		Абзац
	t.Run("TagParagraph", func(t *testing.T) {

	})

	//	 "left"
	t.Run("AttrLeft", func(t *testing.T) {

	})

	// "right"
	t.Run("AttrRight", func(t *testing.T) {

	})

	//	 "center"
	t.Run("AttrCenter", func(t *testing.T) {

	})

	//		Жирный
	t.Run("TagBold", func(t *testing.T) {

	})

	//		Курсив
	t.Run("TagItalic", func(t *testing.T) {

	})

	//		Подчеркнутый текст
	t.Run("TagUnderline", func(t *testing.T) {

	})

	//		Зачеркнутый текст
	t.Run("TagLineThrough", func(t *testing.T) {

	})

	//		Цитата
	t.Run("TagQuote", func(t *testing.T) {

	})

	//	 Мелкий текст внизу
	t.Run("TagSubScript", func(t *testing.T) {

	})

	//		Мелкий текст вверху (степени)
	t.Run("TagSuperScript", func(t *testing.T) {
		transformObj = TextTransform()
		transformObj.AddTagSubScript("{s:}", "{:s}")
		oldText = "<" + TagSuperScript + ">XX</" + TagSuperScript + ">"
		newText = transformObj.Transform(strings.NewReader(oldText))
		if newText != "{s:}XX{:s}" {
			printError(t, "Invalid tag ["+TagSuperScript+"]")
		}
	})
}

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
