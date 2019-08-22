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
