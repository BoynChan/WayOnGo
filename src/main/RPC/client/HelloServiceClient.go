package main

import "net/rpc"

//Author: Boyn
//Date: 2020/4/2

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

type HelloServiceClient struct {
	*rpc.Client
}

func (h *HelloServiceClient) Hello(request string, reply *string) error {
	return h.Client.Call(HelloServiceName+".Hello", request, reply)
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}
