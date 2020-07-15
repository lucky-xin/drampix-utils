package collections

type Set struct {
	elements map[interface{}]int8
}

func NewSet() (set *Set) {
	set = &Set{elements: make(map[interface{}]int8)}
	return
}

func (set *Set) Add(e interface{}) (ok bool) {
	if _, ok := set.elements[e]; !ok {
		set.elements[e] = 1
	}
	return
}

func (set *Set) Remove(e interface{}) (ok bool) {
	if _, ok := set.elements[e]; ok {
		delete(set.elements, e)
	}
	return
}

func (set *Set) Contains(e interface{}) (ok bool) {
	_, ok = set.elements[e]
	return
}
