package set

type Set interface {
	Add(value interface{})
	Values() []interface{}
}

type set struct {
	structMap map[interface{}]struct{}
}

func NewSet() Set {
	return &set{
		structMap: make(map[interface{}]struct{}),
	}
}

func (s *set) Add(value interface{}) {
	s.structMap[value] = struct{}{}
}

func (s *set) Values() []interface{} {
	values := make([]interface{}, 0, len(s.structMap))
	for key := range s.structMap {
		values = append(values, key)
	}
	return values
}
