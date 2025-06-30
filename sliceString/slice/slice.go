package slice

type slice struct {
	s []string
}

func New(s []string) Slice {
	return &slice{s: s}
}

func (s slice) Set() []string {
	mapp := make(map[string]struct{}, len(s.s))
	result := make([]string, 0, len(s.s))

	for _, v := range s.s {
		if _, ok := mapp[v]; !ok {
			mapp[v] = struct{}{}
			result = append(result, v)
		}

	}

	return result
}
