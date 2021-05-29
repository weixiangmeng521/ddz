package set

var Exists = struct{}{}

type Set struct {
	m map[interface{}]struct{}
}
