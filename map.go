package htmlValidator

type TAG string

const (
	TagParagraph TAG = "p"  //	Абзац
	TagDelimiter TAG = "hr" //	Горизонтальная линия разделения

	TagBold        TAG = "b"   //	Жирный
	TagItalic      TAG = "i"   //	Курсив
	TagUnderline   TAG = "u"   //	Подчеркнутый текст
	TagLineThrough TAG = "s"   //	Зачеркнутый текст
	TagQuote       TAG = "q"   //	Цитата
	TagSubScript   TAG = "sub" // Мелкий текст внизу
	TagSuperScript TAG = "sup" //	Мелкий текст вверху (степени)
)

/* Метод что проверяет допустимый ли это для разметки тег */
func IsValidTag(data []byte) (isValid bool, isParagraph bool) {
	switch TAG(data) {
	case TagParagraph:
		return true, true

	case TagDelimiter, TagBold, TagItalic, TagUnderline, TagLineThrough, TagQuote, TagSubScript, TagSuperScript:
		return true, false
	}

	return false, false
}
