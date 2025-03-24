package responses

type Either[L any, R any, M any] struct {
	left  *L
	right *R
	meta  M
}

func Left[L any, R any, M any](meta M, l L) Either[L, R, M] {
	return Either[L, R, M]{left: &l, right: nil, meta: meta}
}

func Right[L any, R any, M any](meta M, r R) Either[L, R, M] {
	return Either[L, R, M]{left: nil, right: &r, meta: meta}
}

func (e Either[L, R, M]) IsLeft() bool {
	return e.left != nil
}

func (e Either[L, R, M]) IsRight() bool {
	return e.right != nil
}

func (e Either[L, R, M]) GetLeft() *L {
	return e.left
}

func (e Either[L, R, M]) GetRight() *R {
	return e.right
}

func (e Either[L, R, M]) GetMeta() M {
	return e.meta
}
