package main

import (
	imooc "cap-imooc/proto/cap"
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
)

type CapServer struct{}

func (c *CapServer) SayHello(ctx context.Context, req *imooc.SayReq, res *imooc.SayResp) error {
	res.Answer = "hello Norman, you are great!" + req.Message + " !!!!"

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("imooc.cap.server"),
	)

	service.Init()

	imooc.RegisterCapHandler(service.Server(), new(CapServer))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
