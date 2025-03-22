package model

type Result[T any] struct {
	Data         T      `json:"data,omitempty"`
	ErrorMessage string `json:"error,omitempty"`
	IsSuccessful bool   `json:"is_success,omitempty"`
	ErrorCode    string `json:"error_code,omitempty"`
}

func Success[T any](data T) *Result[T] {
	return &Result[T]{
		IsSuccessful: true,
		Data:         data,
	}
}

func Failure[T any](errorMessage string, erorCode string) *Result[T] {
	return &Result[T]{
		IsSuccessful: false,
		ErrorMessage: errorMessage,
		ErrorCode:    erorCode,
	}
}

func (r *Result[T]) ChainWithResult(f func(r *Result[T]) *Result[T]) *Result[T] {
	if !r.IsSuccessful {
		return r
	}
	return f(r)
}

func (r *Result[T]) Map(f func(T) T) *Result[T] {
	if !r.IsSuccessful {
		return r
	}
	return &Result[T]{
		IsSuccessful: true,
		Data:         f(r.Data),
	}
}
