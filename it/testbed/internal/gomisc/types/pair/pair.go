package pair

type Pair[L any, R any] interface {
	Left() L
	Right() R
	Unpack() (L, R)
}

func NewPair[L any, R any](left L, right R) Pair[L, R] {
	return &_Pair[L, R]{
		left:  left,
		right: right,
	}
}

type _Pair[L any, R any] struct {
	left  L
	right R
}

func (self *_Pair[L, R]) Left() L {
	return self.left
}

func (self *_Pair[L, R]) Right() R {
	return self.right
}

func (self *_Pair[L, R]) Unpack() (L, R) {
	return self.left, self.right
}
