package service

import "net/http"

type Result[T any] struct {
	Code    int
	Message string
	Data    T
}

func (r *Result[T]) IsFailed() bool {
	return r.Code != http.StatusOK
}

func ResultOK[T any](data T) *Result[T] {
	return &Result[T]{Code: http.StatusOK, Message: "OK", Data: data}
}

func ResultFailed[T any](code int, message string) *Result[T] {
	return &Result[T]{Code: code, Message: message}
}
