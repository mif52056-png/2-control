import (
	"errors"
	"strings"
	"unicode/utf8"
)

func validateUser(name string, age int, email string) error {
	if name == "" || utf8.RuneCountInString(name) >= 50 {
		return errors.New("некорректное имя")
	}
	if age < 18 || age > 120 {
		return errors.New("некорректный возраст")
	}
	if strings.Index(email, "@") == -1 {
		return errors.New("некорректный email")
	}
	return nil
}