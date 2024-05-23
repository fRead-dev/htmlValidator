package htmlValidator

const (
	TagParagraph string = "p"  //	Абзац
	TagDelimiter string = "hr" //	Горизонтальная линия разделения

	TagBold        string = "b"   //	Жирный
	TagItalic      string = "i"   //	Курсив
	TagUnderline   string = "u"   //	Подчеркнутый текст
	TagLineThrough string = "s"   //	Зачеркнутый текст
	TagQuote       string = "q"   //	Цитата
	TagSubScript   string = "sub" // Мелкий текст внизу
	TagSuperScript string = "sup" //	Мелкий текст вверху (степени)
)

/* Метод что проверяет допустимый ли это для разметки тег */
func IsValidTag(data []byte) (tag string, isValid bool, isParagraph bool) {
	tag = string(data)

	switch tag {
	case TagParagraph:
		return tag, true, true

	case TagDelimiter, TagBold, TagItalic, TagUnderline, TagLineThrough, TagQuote, TagSubScript, TagSuperScript:
		return tag, true, false
	}

	return tag, false, false
}

/* Метод, что проверяет допустимый ли атрибут параграфа */
func isValidParagraphAttribute(key []byte) (attr string, isValid bool) {
	attr = string(key)

	switch attr {
	case "left", "right", "center":
		return attr, true
	}

	return attr, false
}
