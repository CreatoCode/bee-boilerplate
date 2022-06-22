package extension

type SliceMap[E, V any] []E
type Slice[T comparable] []T

func (s *SliceMap[E, V]) Map(fn func(E) V) *[]V {
	result := []V{}
	for _, item := range *s {
		result = append(result, fn(item))
	}

	return &result
}

func (s Slice[T]) Exists(e T) bool {
	for _, x := range s {
		if x == e {
			return true
		}
	}
	return false
}

func (s Slice[T]) Filter(fn func(elem T) bool) Slice[T] {
	var output Slice[T]
	for _, x := range s {
		if fn(x) {
			output = append(output, x)
		}
	}
	return output
}

func (xs Slice[T]) Deduplicate() Slice[T] {
	dict := make(map[T]struct{}, len(xs))
	for _, x := range xs {
		dict[x] = struct{}{}
	}
	var output Slice[T]
	for k, _ := range dict {
		output = append(output, k)
	}
	return output
}
