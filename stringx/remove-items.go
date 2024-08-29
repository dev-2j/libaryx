package stringx

func RemoveItems(s *[]string, items ...string) {

	// make a map of items
	m := make(map[string]bool)
	for _, v := range items {
		m[v] = true
	}

	// make a slice of items that are not in the map
	var r []string
	for _, v := range *s {
		if !m[v] {
			r = append(r, v)
		}
	}

	// assign the new slice to the pointer
	*s = r
}
