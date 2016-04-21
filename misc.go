package MiscLib

func InArray(s string, arr []string) int {
	for i, v := range arr {
		if v == s {
			return i
		}
	}
	return -1
}
