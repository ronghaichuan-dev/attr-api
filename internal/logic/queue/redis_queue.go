package queue

import (
	"god-help-service/internal/service"
)

type sQueue struct {
}

func init() {
	service.RegisterQueue(NewQueue())
}

func NewQueue() *sQueue {
	return &sQueue{}
}
