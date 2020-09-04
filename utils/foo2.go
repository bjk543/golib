package utils

import "strings"

func GetMapKeys(mymap map[string]bool) []string {
	keys := make([]string, 0, len(mymap))
	for k := range mymap {
		keys = append(keys, k)
	}
	return keys
}

func HasKey(s string, ss []string) bool {
	for _, v := range ss {
		if strings.Contains(s, v) {
			return true
		}
	}
	return false
}

func GetKeyWord(ss, l, r string) string {

	a := strings.Index(ss, l)
	if a < 0 {
		return ""
	}
	ss = ss[a:]
	b := strings.Index(ss, r)
	if b < 0 {
		return ""
	}
	return ss[len(l):b]
}

func StringArrayToString(s []string) string {
	ss := ""
	for _, s1 := range s {
		if len(strings.TrimSpace(s1)) == 0 {
			continue
		}
		ss = ss + s1 + "\n"
	}
	return ss
}
