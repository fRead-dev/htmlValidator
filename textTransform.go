package htmlValidator

import (
	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/html"
	"io"
)

type paragraphType byte

const (
	pLeft   paragraphType = 10
	pRight  paragraphType = 11
	pCenter paragraphType = 12

	pDef  paragraphType = 20
	pNone               = 0
	pWait               = 1
)

type tagTransformObj struct {
	begin string // Замена открывающего тега
	end   string // Замена закрывающего тега
}

type replaceTagParagraphObj struct {
	def           tagTransformObj //Абзац по умолчанию
	replaceLeft   tagTransformObj //Абзац по левому краю
	replaceRight  tagTransformObj //Абзац по правому краю
	replaceCenter tagTransformObj //Абзац по центру
}

type TextTransformObj struct {
	replaceTagParagraph replaceTagParagraphObj
	replaceTagDelimiter string

	replaceTagBold        tagTransformObj
	replaceTagItalic      tagTransformObj
	replaceTagUnderline   tagTransformObj
	replaceTagLineThrough tagTransformObj
	replaceTagQuote       tagTransformObj
	replaceTagSubScript   tagTransformObj
	replaceTagSuperScript tagTransformObj
}

/* Конструктор класса транчформации текстового блока */
func TextTransform() TextTransformObj {
	obj := TextTransformObj{}
	return obj
}

//###################################################################//

// AddParagraph Установка замены для параграфа
func (obj *TextTransformObj) AddParagraph(begin string, end string) *TextTransformObj {
	obj.replaceTagParagraph.def.begin = begin
	obj.replaceTagParagraph.def.end = end

	if obj.replaceTagParagraph.replaceLeft.begin == "" {
		obj.AddParagraphLeft(begin, end)
	}
	if obj.replaceTagParagraph.replaceRight.begin == "" {
		obj.AddParagraphRight(begin, end)
	}
	if obj.replaceTagParagraph.replaceCenter.begin == "" {
		obj.AddParagraphCenter(begin, end)
	}

	return obj
}

// AddParagraphLeft Установка замены для параграфа с позиционированием по левому краю
func (obj *TextTransformObj) AddParagraphLeft(begin string, end string) *TextTransformObj {
	obj.replaceTagParagraph.replaceLeft.begin = begin
	obj.replaceTagParagraph.replaceLeft.end = end
	return obj
}

// AddParagraphRight Установка замены для параграфа с позиционированием по правому краю
func (obj *TextTransformObj) AddParagraphRight(begin string, end string) *TextTransformObj {
	obj.replaceTagParagraph.replaceRight.begin = begin
	obj.replaceTagParagraph.replaceRight.end = end
	return obj
}

// AddParagraphCenter Установка замены для параграфа с позиционированием по центру
func (obj *TextTransformObj) AddParagraphCenter(begin string, end string) *TextTransformObj {
	obj.replaceTagParagraph.replaceCenter.begin = begin
	obj.replaceTagParagraph.replaceCenter.end = end
	return obj
}

//####//

// AddDelimiter Установка замены тега разделителя
func (obj *TextTransformObj) AddDelimiter(delimiter string) *TextTransformObj {
	obj.replaceTagDelimiter = delimiter
	return obj
}

//####//

// AddTagBold Установка замены для тега стиля `жирный`
func (obj *TextTransformObj) AddTagBold(begin string, end string) *TextTransformObj {
	obj.replaceTagBold.begin = begin
	obj.replaceTagBold.end = end
	return obj
}

// AddTagItalic Установка замены для тега стиля `курсив`
func (obj *TextTransformObj) AddTagItalic(begin string, end string) *TextTransformObj {
	obj.replaceTagItalic.begin = begin
	obj.replaceTagItalic.end = end
	return obj
}

// AddTagUnderline Установка замены для тега стиля `подчеркнутый`
func (obj *TextTransformObj) AddTagUnderline(begin string, end string) *TextTransformObj {
	obj.replaceTagUnderline.begin = begin
	obj.replaceTagUnderline.end = end
	return obj
}

// AddTagLineThrough Установка замены для тега стиля `зачеркнутый`
func (obj *TextTransformObj) AddTagLineThrough(begin string, end string) *TextTransformObj {
	obj.replaceTagLineThrough.begin = begin
	obj.replaceTagLineThrough.end = end
	return obj
}

// AddTagQuote Установка замены для тега стиля `цитата`
func (obj *TextTransformObj) AddTagQuote(begin string, end string) *TextTransformObj {
	obj.replaceTagQuote.begin = begin
	obj.replaceTagQuote.end = end
	return obj
}

// AddTagSubScript Установка замены для тега стиля `мелкий внизу`
func (obj *TextTransformObj) AddTagSubScript(begin string, end string) *TextTransformObj {
	obj.replaceTagSubScript.begin = begin
	obj.replaceTagSubScript.end = end
	return obj
}

// AddTagSuperScript Установка замены для тега стиля `мелкий вверху`
func (obj *TextTransformObj) AddTagSuperScript(begin string, end string) *TextTransformObj {
	obj.replaceTagSuperScript.begin = begin
	obj.replaceTagSuperScript.end = end
	return obj
}

//###################################################################//

func (obj *TextTransformObj) tagReplace(data []byte, isBegin bool) (tag string, isValid bool, isParagraph paragraphType) {
	nameTag := string(data)

	validTag := func() string {
		switch nameTag {

		case TagDelimiter:
			return obj.replaceTagDelimiter
		case TagBold:
			if isBegin {
				return obj.replaceTagBold.begin
			} else {
				return obj.replaceTagBold.end
			}
		case TagItalic:
			if isBegin {
				return obj.replaceTagItalic.begin
			} else {
				return obj.replaceTagItalic.end
			}
		case TagUnderline:
			if isBegin {
				return obj.replaceTagUnderline.begin
			} else {
				return obj.replaceTagUnderline.end
			}
		case TagLineThrough:
			if isBegin {
				return obj.replaceTagLineThrough.begin
			} else {
				return obj.replaceTagLineThrough.end
			}
		case TagQuote:
			if isBegin {
				return obj.replaceTagQuote.begin
			} else {
				return obj.replaceTagQuote.end
			}
		case TagSubScript:
			if isBegin {
				return obj.replaceTagSubScript.begin
			} else {
				return obj.replaceTagSubScript.end
			}
		case TagSuperScript:
			if isBegin {
				return obj.replaceTagSuperScript.begin
			} else {
				return obj.replaceTagSuperScript.end
			}

		}
		return tag
	}

	switch nameTag {
	case TagParagraph:
		return "", true, pWait

	case TagDelimiter, TagBold, TagItalic, TagUnderline, TagLineThrough, TagQuote, TagSubScript, TagSuperScript:
		return validTag(), true, pNone
	}

	return "", false, pNone
}

func (obj *TextTransformObj) paragraphReplace(key []byte) (isParagraph paragraphType) {
	attr := string(key)

	switch attr {
	case AttrLeft:
		return pLeft
	case AttrRight:
		return pRight
	case AttrCenter:
		return pCenter
	}

	return pDef
}

func (obj *TextTransformObj) paragraphPrint(types paragraphType, isBegin bool) string {
	switch types {
	case pDef, pWait:
		if isBegin {
			return obj.replaceTagParagraph.def.begin
		} else {
			return obj.replaceTagParagraph.def.end
		}

	case pLeft:
		if isBegin {
			return obj.replaceTagParagraph.replaceLeft.begin
		} else {
			return obj.replaceTagParagraph.replaceLeft.end
		}

	case pRight:
		if isBegin {
			return obj.replaceTagParagraph.replaceRight.begin
		} else {
			return obj.replaceTagParagraph.replaceRight.end
		}

	case pCenter:
		if isBegin {
			return obj.replaceTagParagraph.replaceCenter.begin
		} else {
			return obj.replaceTagParagraph.replaceCenter.end
		}
	}

	return ""
}

/* Трансормация входного html-текста согласно параметрам */
func (obj *TextTransformObj) Transform(htmlText io.Reader) (retText string) {
	parser := html.NewLexer(parse.NewInput(htmlText))
	var waitParagraph paragraphType = pNone

	for {
		typeToken, data := parser.Next()

		switch typeToken {
		case html.StartTagCloseToken, html.StartTagVoidToken:
			continue

		case html.AttributeToken:
			if waitParagraph != pNone { //обрабатываем только ожидающие атрибуты
				waitParagraph = obj.paragraphReplace(parser.AttrKey())
				retText += obj.paragraphPrint(waitParagraph, true)
			}

		case html.StartTagToken:
			tag, isValid, isParagraph := obj.tagReplace(parser.AttrKey(), true)
			if isValid { //если тег валиден
				if isParagraph == pNone { //и если это не параграф
					retText += tag

				} else { //если таки параграф
					waitParagraph = pWait
				}

			}

		case html.EndTagToken:
			tag, isValid, isParagraph := obj.tagReplace(parser.AttrKey(), false)
			if isValid { //если тег валиден
				if isParagraph == pNone {
					retText += tag //и если это не параграф
				} else {
					retText += obj.paragraphPrint(waitParagraph, false)
					waitParagraph = pNone
				}

			}

		case html.TextToken:
			if waitParagraph == pWait {
				retText += obj.paragraphPrint(waitParagraph, true)
			}
			retText += string(data)

		case html.ErrorToken:
			return retText

		default:
			continue
		}
	}
}
