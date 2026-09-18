import (
	"strings"
	"unicode/utf8"
)

type TextStats struct {
	Chars     int
	Words     int
	Sentences int
}

func textStats(text string) TextStats {
	return TextStats{
		Chars:     utf8.RuneCountInString(text),
		Words:     len(strings.Fields(text)),
		Sentences: strings.Count(text, ".") + strings.Count(text, "!") + strings.Count(text, "?"),
	}
}