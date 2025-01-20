package main

import (
	"context"

	api "github.com/North-al/kitex_demo/hello/kitex_gen/api"
)

// HelloImpl implements the last service interface defined in the IDL.
type HelloImpl struct{}

// Echo implements the HelloImpl interface.
func (s *HelloImpl) Echo(ctx context.Context, request *api.Request) (resp *api.Response, err error) {
	// TODO: Your code here...
	return
}

// Say implements the HelloImpl interface.
func (s *HelloImpl) Say(ctx context.Context, request *api.SayRequest) (resp *api.SayResponse, err error) {
	// TODO: Your code here...
	return
}

// Add implements the HelloImpl interface.
func (s *HelloImpl) Add(ctx context.Context, request *api.AddRequest) (resp *api.AddResponse, err error) {
	// TODO: Your code here...
	return
}
