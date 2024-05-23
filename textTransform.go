package htmlValidator

import (
	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/html"
	"io"
)

type TagTransformObj struct {
	Begin string // Замена открывающего тега
	End   string // Замена закрывающего тега
}

type ReplaceTagParagraphObj struct {
	Default       TagTransformObj //Абзац по умолчанию
	ReplaceLeft   TagTransformObj //Абзац по левому краю
	ReplaceRight  TagTransformObj //Абзац по правому краю
	ReplaceCenter TagTransformObj //Абзац по центру
}

type TextTransformObj struct {
	ReplaceTagParagraph ReplaceTagParagraphObj
	ReplaceTagDelimiter string

	ReplaceTagBold        TagTransformObj
	ReplaceTagItalic      TagTransformObj
	ReplaceTagUnderline   TagTransformObj
	ReplaceTagLineThrough TagTransformObj
	ReplaceTagQuote       TagTransformObj
	ReplaceTagSubScript   TagTransformObj
	ReplaceTagSuperScript TagTransformObj
}

/* Конструктор класса транчформации текстового блока */
func TextTransform() TextTransformObj {
	obj := TextTransformObj{}
	return obj
}

//###################################################################//

// AddParagraph Установка замены для параграфа
func (obj *TextTransformObj) AddParagraph(begin string, end string) *TextTransformObj {
	obj.ReplaceTagParagraph.Default.Begin = begin
	obj.ReplaceTagParagraph.Default.End = end

	if obj.ReplaceTagParagraph.ReplaceLeft.Begin == "" {
		obj.AddParagraphLeft(begin, end)
	}
	if obj.ReplaceTagParagraph.ReplaceRight.Begin == "" {
		obj.AddParagraphRight(begin, end)
	}
	if obj.ReplaceTagParagraph.ReplaceCenter.Begin == "" {
		obj.AddParagraphCenter(begin, end)
	}

	return obj
}

// AddParagraphLeft Установка замены для параграфа с позиционированием по левому краю
func (obj *TextTransformObj) AddParagraphLeft(begin string, end string) *TextTransformObj {
	obj.ReplaceTagParagraph.ReplaceLeft.Begin = begin
	obj.ReplaceTagParagraph.ReplaceLeft.End = end
	return obj
}

// AddParagraphRight Установка замены для параграфа с позиционированием по правому краю
func (obj *TextTransformObj) AddParagraphRight(begin string, end string) *TextTransformObj {
	obj.ReplaceTagParagraph.ReplaceRight.Begin = begin
	obj.ReplaceTagParagraph.ReplaceRight.End = end
	return obj
}

// AddParagraphCenter Установка замены для параграфа с позиционированием по центру
func (obj *TextTransformObj) AddParagraphCenter(begin string, end string) *TextTransformObj {
	obj.ReplaceTagParagraph.ReplaceCenter.Begin = begin
	obj.ReplaceTagParagraph.ReplaceCenter.End = end
	return obj
}

//####//

// AddDelimiter Установка замены тега разделителя
func (obj *TextTransformObj) AddDelimiter(delimiter string) *TextTransformObj {
	obj.ReplaceTagDelimiter = delimiter
	return obj
}

//####//

// AddTagBold Установка замены для тега стиля `жирный`
func (obj *TextTransformObj) AddTagBold(begin string, end string) *TextTransformObj {
	obj.ReplaceTagBold.Begin = begin
	obj.ReplaceTagBold.End = end
	return obj
}

// AddTagItalic Установка замены для тега стиля `курсив`
func (obj *TextTransformObj) AddTagItalic(begin string, end string) *TextTransformObj {
	obj.ReplaceTagItalic.Begin = begin
	obj.ReplaceTagItalic.End = end
	return obj
}

// AddTagUnderline Установка замены для тега стиля `подчеркнутый`
func (obj *TextTransformObj) AddTagUnderline(begin string, end string) *TextTransformObj {
	obj.ReplaceTagUnderline.Begin = begin
	obj.ReplaceTagUnderline.End = end
	return obj
}

// AddTagLineThrough Установка замены для тега стиля `зачеркнутый`
func (obj *TextTransformObj) AddTagLineThrough(begin string, end string) *TextTransformObj {
	obj.ReplaceTagLineThrough.Begin = begin
	obj.ReplaceTagLineThrough.End = end
	return obj
}

// AddTagQuote Установка замены для тега стиля `цитата`
func (obj *TextTransformObj) AddTagQuote(begin string, end string) *TextTransformObj {
	obj.ReplaceTagQuote.Begin = begin
	obj.ReplaceTagQuote.End = end
	return obj
}

// AddTagSubScript Установка замены для тега стиля `мелкий внизу`
func (obj *TextTransformObj) AddTagSubScript(begin string, end string) *TextTransformObj {
	obj.ReplaceTagSubScript.Begin = begin
	obj.ReplaceTagSubScript.End = end
	return obj
}

// AddTagSuperScript Установка замены для тега стиля `мелкий вверху`
func (obj *TextTransformObj) AddTagSuperScript(begin string, end string) *TextTransformObj {
	obj.ReplaceTagSuperScript.Begin = begin
	obj.ReplaceTagSuperScript.End = end
	return obj
}

//###################################################################//

/* Трансормация входного html-текста согласно параметрам */
func (obj *TextTransformObj) Transform(htmlText io.Reader) (retText string) {
	parser := html.NewLexer(parse.NewInput(htmlText))

	for {
		typeToken, data := parser.Next()

		switch typeToken {
		case html.StartTagCloseToken, html.StartTagVoidToken:
			continue

		case html.TextToken:
			retText += string(data)

		case html.ErrorToken:
			return retText

		default:
			continue
		}
	}
}
