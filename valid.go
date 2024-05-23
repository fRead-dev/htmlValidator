package htmlValidator

import (
	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/html"
	"io"
	"unicode/utf8"
)

type ValidateObj struct {
	Size struct {
		Bytes   uint64
		Symbols uint64
	}
	Tags struct {
		Global uint64
		Errors map[string]uint32

		Delimiter   uint32
		Bold        uint32
		Italic      uint32
		Underline   uint32
		LineThrough uint32
		Quote       uint32
		SubScript   uint32
		SuperScript uint32

		Paragraphs struct {
			Global uint32
			Def    uint32
			Left   uint32
			Right  uint32
			Center uint32
		}
	}
}

/* Универсальный метод проверки (собирается статистика по всем меткам) */
func Validate(htmlText io.Reader) ValidateObj {
	obj := ValidateObj{}
	obj.Tags.Errors = make(map[string]uint32)

	waitParagraph := false
	parser := html.NewLexer(parse.NewInput(htmlText))
	for {
		typeToken, data := parser.Next()

		switch typeToken {
		case html.StartTagCloseToken, html.StartTagVoidToken:
			continue

		case html.AttributeToken:
			if waitParagraph { //обрабатываем только ожидающие атрибуты
				key := string(parser.AttrKey())
				switch key {
				case AttrLeft:
					obj.Tags.Paragraphs.Left += 1
				case AttrRight:
					obj.Tags.Paragraphs.Right += 1
				case AttrCenter:
					obj.Tags.Paragraphs.Center += 1
				default:
					obj.Tags.Paragraphs.Def += 1
				}

				waitParagraph = false
			}

		case html.StartTagToken:
			tag := string(parser.AttrKey())
			switch tag {

			case TagParagraph:
				obj.Tags.Paragraphs.Global += 1
				waitParagraph = true

			case TagDelimiter:
				obj.Tags.Delimiter += 1
			case TagBold:
				obj.Tags.Bold += 1
			case TagItalic:
				obj.Tags.Italic += 1
			case TagUnderline:
				obj.Tags.Underline += 1
			case TagLineThrough:
				obj.Tags.LineThrough += 1
			case TagQuote:
				obj.Tags.Quote += 1
			case TagSubScript:
				obj.Tags.SubScript += 1
			case TagSuperScript:
				obj.Tags.SuperScript += 1

			default:
				obj.Tags.Errors[tag] += 1
			}

		case html.EndTagToken:
			waitParagraph = false
			continue

		case html.TextToken:
			if waitParagraph {
				obj.Tags.Paragraphs.Def += 1
			}
			waitParagraph = false
			obj.Size.Bytes += uint64(len(data))
			obj.Size.Symbols += uint64(utf8.RuneCountInString(string(data)))

		case html.ErrorToken:
			return obj

		default:
			continue
		}
	}
}

//###################################################################//
