package entry

type Entry[K comparable, V any] interface {
	Key() K
	Value() V
}

func NewEntry[K comparable, V any](key K, value V) Entry[K, V] {
	return &_Entry[K, V]{
		key:   key,
		value: value,
	}
}

type _Entry[K comparable, V any] struct {
	key   K
	value V
}

func (self *_Entry[K, V]) Key() K {
	return self.key
}

func (self *_Entry[K, V]) Value() V {
	return self.value
}
