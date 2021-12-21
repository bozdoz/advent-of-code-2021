package types

type Queue[T any] []*T

func (queue *Queue[T]) Shift() *T {
	if len(*queue) < 0 {
		return nil
	}

	first := (*queue)[0]
	*queue = (*queue)[1:]

	return first
}

func (queue *Queue[T]) Push(item *T) {
	*queue = append(*queue, item)
}
