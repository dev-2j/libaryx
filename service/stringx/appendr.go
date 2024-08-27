package stringx

func Appendr(s *[]string, cs ...string) {

	// remove items from the slice
	RemoveItems(s, cs...)

	// append the new items
	*s = append(*s, cs...)

}
