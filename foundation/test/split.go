package split

import "strings"

func Split(str, sep string) []string {
	ret := make([]string, 0, 5)

	index := strings.Index(str, sep)
	for index >= 0 {
		if index != 0 {
			ret = append(ret, str[:index])
		}
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}
	ret = append(ret, str)
	return ret
}
