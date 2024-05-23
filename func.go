package htmlValidator

import (
	"io"
	"strings"
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

/* Возрашвет только текст и БЫСТРО */
func TextFast(htmlText io.Reader) string {
	buf := new(strings.Builder)
	io.Copy(buf, htmlText)

	input := buf.String()
	output := make([]rune, 0, len(input))
	inTag := false

	for _, char := range input {
		if char == '<' {
			inTag = true
		} else if char == '>' {
			inTag = false
		} else if !inTag {
			output = append(output, char)
		}
	}

	return string(output)
}
