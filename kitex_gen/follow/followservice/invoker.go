// Code generated by Kitex v0.4.4. DO NOT EDIT.

package followservice

import (
	server "github.com/cloudwego/kitex/server"
	follow "tiktok/kitex_gen/follow"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler follow.FollowService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
