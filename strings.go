package lazygo

import (
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
