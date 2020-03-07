package xstr

import (
	"strings"
)

// const ossHost = "http://127.0.0.1:6006"

const ossHost = "https://oss.likecho.com"

// JoinString format string slice like:n1n2n3.
func StrJoin(strs ...string) string {
	return strings.Join(strs, "")
}

func GetFullImageUrl(str string) string {
	strs := []string{ossHost, str}
	return strings.Join(strs, "")
}
