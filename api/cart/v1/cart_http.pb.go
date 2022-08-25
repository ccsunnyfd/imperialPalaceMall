// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.4.1
// - protoc             v3.21.5
// source: cart/v1/cart.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationCartServiceAddCart = "/api.cart.v1.CartService/AddCart"

type CartServiceHTTPServer interface {
	AddCart(context.Context, *AddCartRequest) (*AddCartReply, error)
}

func RegisterCartServiceHTTPServer(s *http.Server, srv CartServiceHTTPServer) {
	r := s.Route("/")
	r.PUT("/v1/user/my/cart", _CartService_AddCart0_HTTP_Handler(srv))
}

func _CartService_AddCart0_HTTP_Handler(srv CartServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddCartRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCartServiceAddCart)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddCart(ctx, req.(*AddCartRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddCartReply)
		return ctx.Result(200, reply)
	}
}

type CartServiceHTTPClient interface {
	AddCart(ctx context.Context, req *AddCartRequest, opts ...http.CallOption) (rsp *AddCartReply, err error)
}

type CartServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewCartServiceHTTPClient(client *http.Client) CartServiceHTTPClient {
	return &CartServiceHTTPClientImpl{client}
}

func (c *CartServiceHTTPClientImpl) AddCart(ctx context.Context, in *AddCartRequest, opts ...http.CallOption) (*AddCartReply, error) {
	var out AddCartReply
	pattern := "/v1/user/my/cart"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCartServiceAddCart))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}