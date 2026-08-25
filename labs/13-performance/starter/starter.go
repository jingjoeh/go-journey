package starter

func Deduplicate(values []string) []string {
	seen := make(map[string]struct{})
	results := make([]string, 0, len(values))
	for _, v := range values {
		_, ok := seen[v]
		if !ok {
			seen[v] = struct{}{}
			results = append(results, v)
		}
	}

	return results
}
