package data

func paginate[T any](s []T, page, perPage int) []T {
	if len(s) == 0 {
		return nil
	}
	inBounds := func(v int) int {
		return max(0, min(v, len(s)-1))
	}
	start := (page - 1) * perPage
	startBounded := inBounds(start)
	if startBounded >= len(s) {
		return nil
	}
	end := page*perPage - 1
	endBounded := inBounds(end) + 1
	return s[startBounded:endBounded]
}
