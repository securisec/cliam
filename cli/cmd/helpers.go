package cmd

// remove duplicates from a slice of strings
func removeDuplicates(slice []string) []string {
	if len(slice) == 0 {
		return slice
	}
	m := make(map[string]bool)
	for _, s := range slice {
		m[s] = true
	}
	var res []string
	for k := range m {
		res = append(res, k)
	}
	return res
}
