package main

import "net/rpc"

//Author: Boyn
//Date: 2020/4/2

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface interface {
	Hello(request String, reply *String) error
}

type HelloService struct{}

func (p *HelloService) Hello(request String, reply *String) error {
	reply.Value = "hello:" + request.Value
	return nil
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}
