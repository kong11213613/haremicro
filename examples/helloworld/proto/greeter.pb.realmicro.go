// Code generated by protoc-gen-haremicro. DO NOT EDIT.
// source: greeter.proto

package greeter

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/kong11213613/haremicro/client"
	server "github.com/kong11213613/haremicro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Greeter service

// Client API for Greeter service

type GreeterService interface {
	Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type greeterService struct {
	c    client.Client
	name string
}

func NewGreeterService(name string, c client.Client) GreeterService {
	return &greeterService{
		c:    c,
		name: name,
	}
}

func (c *greeterService) Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Greeter.Hello", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterHandler interface {
	Hello(context.Context, *Request, *Response) error
}

func RegisterGreeterHandler(s server.Server, hdlr GreeterHandler, opts ...server.HandlerOption) error {
	type greeter interface {
		Hello(ctx context.Context, in *Request, out *Response) error
	}
	type Greeter struct {
		greeter
	}
	h := &greeterHandler{hdlr}
	return s.Handle(s.NewHandler(&Greeter{h}, opts...))
}

type greeterHandler struct {
	GreeterHandler
}

func (h *greeterHandler) Hello(ctx context.Context, in *Request, out *Response) error {
	return h.GreeterHandler.Hello(ctx, in, out)
}
