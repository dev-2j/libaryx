package stringx

// func RemoveDuplicate(strSlice []string) []string {

func Unique(s []string) []string {
	// https://stackoverflow.com/questions/66643946/how-to-remove-duplicates-strings-or-int-from-slice-in-go
	vs := make(map[string]bool)
	for _, v := range s {
		vs[v] = true
	}
	list := make([]string, 0, len(vs))
	for v := range vs {
		list = append(list, v)
	}
	return list
}
