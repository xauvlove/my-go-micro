package utils

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/consul"
)

func InitMicro() micro.Service {
	registry := consul.NewRegistry()
	return micro.NewService(micro.Registry(registry))
}
