package main

import (
	"fmt"
	"github.com/dayan-be/demo-service/global"
	"github.com/dayan-be/demo-service/proto"
	"github.com/micro/go-micro"
	"context"
	"github.com/micro/go-micro/registry"
)

func main() {
	service := micro.NewService(micro.Registry(registry.NewRegistry(registry.Addrs(global.Config().Registry.Addr))))
	service.Init()
	cl := demo.NewSayService("demo-service", service.Client())
	rsp, err := cl.Hello(context.Background(), &demo.HelloReq{
		Name: "EvanPan",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Msg)
}
