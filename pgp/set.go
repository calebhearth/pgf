package pgp

// Set implements a collection of values with no duplicates. It can be iterated
// as a map[string]struct{} where the value is always an empty struct.
type set map[string]struct{}

// Add appends v to s.
func (s set) Add(v string) {
	s[v] = struct{}{}
}

// Reduce collects the results of calling transformer for each value of s and
// returns the collection.
func (s set) Reduce(transformer func(string) string) []string {
	results := []string{}
	for key, _ := range s {
		results = append(results, transformer(key))
	}
	return results
}
