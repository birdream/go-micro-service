package main

import (
	"fmt"

	"github.com/birdream/go-micro-service/imooc-stu/user/domain/repository"
	service2 "github.com/birdream/go-micro-service/imooc-stu/user/domain/service"
	"github.com/birdream/go-micro-service/imooc-stu/user/handler"
	user "github.com/birdream/go-micro-service/imooc-stu/user/proto"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//服务参数设置
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)
	//初始化服务
	srv.Init()

	//创建数据库连接
	db, err := gorm.Open("mysql", "root:123456@/micro?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.SingularTable(true)

	//只执行一次，数据表初始化
	//rp:=repository.NewUserRepository(db)
	//rp.InitTable()

	//创建服务实例
	userDataService := service2.NewUserDataService(repository.NewUserRepository(db))
	//注册Handler
	err = user.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService})
	if err != nil {
		fmt.Println(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
