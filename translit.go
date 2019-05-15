package gotranslit

import (
	"strings"
	"unicode"
)

var cyrillic = map[int32]string{
	'а': "a",
	'б': "b",
	'в': "v",
	'г': "g",
	'д': "d",
	'е': "e",
	'ё': "yo",
	'ж': "zh",
	'з': "z",
	'и': "i",
	'й': "j",
	'к': "k",
	'л': "l",
	'м': "m",
	'н': "n",
	'о': "o",
	'п': "p",
	'р': "r",
	'с': "s",
	'т': "t",
	'у': "u",
	'ф': "f",
	'х': "h",
	'ц': "c",
	'ч': "ch",
	'ш': "sh",
	'щ': "sch",
	'ъ': "",
	'ы': "y",
	'ь': "",
	'э': "e",
	'ю': "ju",
	'я': "ja",
	' ': "_",
}

// ToCyrillic translits cyrillic to latin
func ToCyrillic(text string) string {
	out := make([]string, len(text))
	for _, t := range text {
		if c, ok := cyrillic[t]; ok {
			out = append(out, c)
		} else {
			if unicode.IsUpper(t) {
				if c, ok := cyrillic[unicode.ToLower(t)]; ok {
					out = append(out, string(strings.ToUpper(c)))
					continue
				}
			}
			out = append(out, string(t))
		}
	}
	return strings.Join(out, "")
}
