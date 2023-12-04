package main

import (
	"my-micro/service/user/handler"
	user "my-micro/service/user/proto"
	"my-micro/web/model"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {

	// 初始化数据库
	model.InitDB()

	registry := consul.NewRegistry()

	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
		micro.Address("localhost:15331"),
		micro.Registry(registry),
	)
	user.RegisterUserHandler(service.Server(), new(handler.User))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
