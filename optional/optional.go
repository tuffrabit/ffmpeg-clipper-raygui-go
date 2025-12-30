package optional

type Optional[T any] struct {
	isSet bool
	value T
}

func (o *Optional[T]) Set(item T) {
	o.value = item
	o.isSet = true
}

func (o *Optional[T]) IsSet() bool {
	return o.isSet
}

func (o *Optional[T]) Get() T {
	return o.value
}
