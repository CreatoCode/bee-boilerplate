package extension

//**把用于构造函数[柯里化](https://www.zhihu.com/search?q=%E6%9F%AF%E9%87%8C%E5%8C%96&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra=%7B%22sourceType%22%3A%22answer%22%2C%22sourceId%22%3A1611857930%7D)的ap functor抽出来**

// A Option sets options.
type Option[T any] interface {
	apply(T)
}

// Option wraps a function that modifies options into an
// implementation of the Option interface.
type applyOption[T any] struct {
	f func(T)
}

func (a *applyOption[T]) apply(do T) {
	a.f(do)
}

//NewApplyOption  create an ap functor
func NewApplyOption[T any](f func(T)) *applyOption[T] {
	return &applyOption[T]{
		f: f,
	}
}
