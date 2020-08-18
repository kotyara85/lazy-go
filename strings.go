package lazygo

import (
	"regexp"
	"strings"
)

// ToLower lowers the variable in place
func ToLower(str *string) {
	tmp := *str
	*str = strings.ToLower(tmp)
}

// ToUpper appers the variable in place
func ToUpper(str *string) {
	tmp := *str
	*str = strings.ToUpper(tmp)
}

// RightPad fixes issues with fmt.Sprintf("") right padding
func RightPad(s *string, pad string, plength int) {
	tmp := *s
	for i := len(tmp); i < plength; i++ {
		tmp = tmp + pad
	}
	*s = tmp
}

// Replace replaces text
func Replace(s *string, reg string, repl string) error {
	if re, err := regexp.Compile(reg); err == nil {
		*s = re.ReplaceAllLiteralString(*s, repl)
	} else {
		return err
	}
	return nil
}
