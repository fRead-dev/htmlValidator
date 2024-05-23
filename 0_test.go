package htmlValidator

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func printError(t *testing.T, text string) {
	t.Error(fmt.Errorf("%s", text))
	t.Fail()
}
func printFatal(t *testing.T, text string, err error) {
	t.Fatalf(text, err)
}

func randomTag() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, 3)
	rand.Seed(time.Now().UnixNano())

	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
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
		transformObj = TextTransform()
		transformObj.AddDelimiter("{:s:}")
		oldText = "<" + TagDelimiter + ">XX</" + TagDelimiter + ">"
		newText = transformObj.Transform(strings.NewReader(oldText))
		if newText != "{:s:}XX{:s:}" {
			t.Log("NEW", newText)
			t.Log("OLD", oldText)
			printError(t, "Invalid tag ["+TagDelimiter+"]")
		}
	})

	//		Абзац
	t.Run("TagParagraph", func(t *testing.T) {
		transformObj = TextTransform()
		transformObj.AddParagraph("{s:}", "{:s}")
		oldText = "<" + TagParagraph + ">XX</" + TagParagraph + ">"
		newText = transformObj.Transform(strings.NewReader(oldText))
		if newText != "{s:}XX{:s}" {
			t.Log("NEW", newText)
			t.Log("OLD", oldText)
			printError(t, "Invalid tag ["+TagParagraph+"]")
		}
	})

	//	 "left"
	t.Run("AttrLeft", func(t *testing.T) {
		transformObj = TextTransform()
		transformObj.AddParagraphLeft("{s:}", "{:s}")
		oldText = "<" + TagParagraph + " " + AttrLeft + ">XX</" + TagParagraph + ">"
		newText = transformObj.Transform(strings.NewReader(oldText))
		if newText != "{s:}XX{:s}" {
			t.Log("NEW", newText)
			t.Log("OLD", oldText)
			printError(t, "Invalid tag ["+TagParagraph+" "+AttrLeft+"]")
		}
	})

	// "right"
	t.Run("AttrRight", func(t *testing.T) {
		transformObj = TextTransform()
		transformObj.AddParagraphRight("{s:}", "{:s}")
		oldText = "<" + TagParagraph + " " + AttrRight + ">XX</" + TagParagraph + ">"
		newText = transformObj.Transform(strings.NewReader(oldText))
		if newText != "{s:}XX{:s}" {
			t.Log("NEW", newText)
			t.Log("OLD", oldText)
			printError(t, "Invalid tag ["+TagParagraph+" "+AttrRight+"]")
		}
	})

	//	 "center"
	t.Run("AttrCenter", func(t *testing.T) {
		transformObj = TextTransform()
		transformObj.AddParagraphCenter("{s:}", "{:s}")
		oldText = "<" + TagParagraph + " " + AttrCenter + ">XX</" + TagParagraph + ">"
		newText = transformObj.Transform(strings.NewReader(oldText))
		if newText != "{s:}XX{:s}" {
			t.Log("NEW", newText)
			t.Log("OLD", oldText)
			printError(t, "Invalid tag ["+TagParagraph+" "+AttrCenter+"]")
		}
	})

	//		Жирный
	t.Run("TagBold", func(t *testing.T) {
		transformObj = TextTransform()
		transformObj.AddTagBold("{s:}", "{:s}")
		oldText = "<" + TagBold + ">XX</" + TagBold + ">"
		newText = transformObj.Transform(strings.NewReader(oldText))
		if newText != "{s:}XX{:s}" {
			t.Log("NEW", newText)
			t.Log("OLD", oldText)
			printError(t, "Invalid tag ["+TagBold+"]")
		}
	})

	//		Курсив
	t.Run("TagItalic", func(t *testing.T) {
		transformObj = TextTransform()
		transformObj.AddTagItalic("{s:}", "{:s}")
		oldText = "<" + TagItalic + ">XX</" + TagItalic + ">"
		newText = transformObj.Transform(strings.NewReader(oldText))
		if newText != "{s:}XX{:s}" {
			t.Log("NEW", newText)
			t.Log("OLD", oldText)
			printError(t, "Invalid tag ["+TagItalic+"]")
		}
	})

	//		Подчеркнутый текст
	t.Run("TagUnderline", func(t *testing.T) {
		transformObj = TextTransform()
		transformObj.AddTagUnderline("{s:}", "{:s}")
		oldText = "<" + TagUnderline + ">XX</" + TagUnderline + ">"
		newText = transformObj.Transform(strings.NewReader(oldText))
		if newText != "{s:}XX{:s}" {
			t.Log("NEW", newText)
			t.Log("OLD", oldText)
			printError(t, "Invalid tag ["+TagUnderline+"]")
		}
	})

	//		Зачеркнутый текст
	t.Run("TagLineThrough", func(t *testing.T) {
		transformObj = TextTransform()
		transformObj.AddTagLineThrough("{s:}", "{:s}")
		oldText = "<" + TagLineThrough + ">XX</" + TagLineThrough + ">"
		newText = transformObj.Transform(strings.NewReader(oldText))
		if newText != "{s:}XX{:s}" {
			t.Log("NEW", newText)
			t.Log("OLD", oldText)
			printError(t, "Invalid tag ["+TagLineThrough+"]")
		}
	})

	//		Цитата
	t.Run("TagQuote", func(t *testing.T) {
		transformObj = TextTransform()
		transformObj.AddTagQuote("{s:}", "{:s}")
		oldText = "<" + TagQuote + ">XX</" + TagQuote + ">"
		newText = transformObj.Transform(strings.NewReader(oldText))
		if newText != "{s:}XX{:s}" {
			t.Log("NEW", newText)
			t.Log("OLD", oldText)
			printError(t, "Invalid tag ["+TagQuote+"]")
		}
	})

	//	 Мелкий текст внизу
	t.Run("TagSubScript", func(t *testing.T) {
		transformObj = TextTransform()
		transformObj.AddTagSubScript("{s:}", "{:s}")
		oldText = "<" + TagSubScript + ">XX</" + TagSubScript + ">"
		newText = transformObj.Transform(strings.NewReader(oldText))
		if newText != "{s:}XX{:s}" {
			t.Log("NEW", newText)
			t.Log("OLD", oldText)
			printError(t, "Invalid tag ["+TagSubScript+"]")
		}
	})

	//		Мелкий текст вверху (степени)
	t.Run("TagSuperScript", func(t *testing.T) {
		transformObj = TextTransform()
		transformObj.AddTagSuperScript("{s:}", "{:s}")
		oldText = "<" + TagSuperScript + ">XX</" + TagSuperScript + ">"
		newText = transformObj.Transform(strings.NewReader(oldText))
		if newText != "{s:}XX{:s}" {
			t.Log("NEW", newText)
			t.Log("OLD", oldText)
			printError(t, "Invalid tag ["+TagSuperScript+"]")
		}
	})
}
