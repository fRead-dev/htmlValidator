package htmlValidator

import (
	"io"
)

/* Приведение html-текста к "стандартному" виду */
func Standardization(htmlText io.Reader) string {
	transformObj := TextTransform()

	transformObj.AddParagraph("<"+TagParagraph+">", "</"+TagParagraph+">")
	transformObj.AddParagraphLeft("<"+TagParagraph+" "+AttrLeft+">", "</"+TagParagraph+">")
	transformObj.AddParagraphRight("<"+TagParagraph+" "+AttrRight+">", "</"+TagParagraph+">")
	transformObj.AddParagraphCenter("<"+TagParagraph+" "+AttrCenter+">", "</"+TagParagraph+">")

	transformObj.AddDelimiter("<" + TagDelimiter + ">")

	transformObj.AddTagBold("<"+TagBold+">", "</"+TagBold+">")
	transformObj.AddTagItalic("<"+TagItalic+">", "</"+TagItalic+">")
	transformObj.AddTagUnderline("<"+TagUnderline+">", "</"+TagUnderline+">")
	transformObj.AddTagLineThrough("<"+TagLineThrough+">", "</"+TagLineThrough+">")
	transformObj.AddTagQuote("<"+TagQuote+">", "</"+TagQuote+">")
	transformObj.AddTagSubScript("<"+TagSubScript+">", "</"+TagSubScript+">")
	transformObj.AddTagSuperScript("<"+TagSuperScript+">", "</"+TagSuperScript+">")

	return transformObj.Transform(htmlText)
}

/* Возращает только текст */
func Text(htmlText io.Reader) string {
	transformObj := TextTransform()
	return transformObj.Transform(htmlText)
}
