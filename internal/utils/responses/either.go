package response

type Either[L any, R any] struct {
	left  *L
	right *R
}

func Left[L any, R any](l L) Either[L, R] {
	return Either[L, R]{left: &l, right: nil}
}

func Right[L any, R any](r R) Either[L, R] {
	return Either[L, R]{left: nil, right: &r}
}

func (e Either[L, R]) IsLeft() bool {
	return e.left != nil
}

func (e Either[L, R]) IsRight() bool {
	return e.right != nil
}

func (e Either[L, R]) GetLeft() *L {
	return e.left
}

func (e Either[L, R]) GetRight() *R {
	return e.right
}
