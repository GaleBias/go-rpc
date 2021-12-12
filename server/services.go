package main

import "sync"

// 服务名到服务的映射
var services sync.Map

func AddService(service Service) {
	services.Store(service.ServiceName(),service)
}

type Service interface {  // 因为 AddService 需要维护一个serviceName到Service的映射，所以要定义一个共同行为
	ServiceName() string
	//ShutDown()
}
type HelloService interface {
	Service
	SayHello(input * Input)(*Output,error)
}
type UserService interface {
	Service
	GetUser(req *GetUserReq) (*GetUserResp,error)
}

type helloService struct {

}
type userService struct {

}

