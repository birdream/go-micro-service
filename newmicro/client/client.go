package main

import (
	imooc "cap-imooc/proto/cap"
	"fmt"

	"context"

	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(
		micro.Name("cap.imooc.client"),
	)

	service.Init()

	capImooc := imooc.NewCapService("imooc.cap.server", service.Client())

	res, err := capImooc.SayHello(context.TODO(), &imooc.SayReq{
		Message: "Follow ME",
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
