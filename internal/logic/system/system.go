package system

import "god-help-service/internal/service"

type sSystem struct {
}

func init() {
	service.RegisterSystem(NewSystem())
}

func NewSystem() *sSystem {
	return &sSystem{}
}
