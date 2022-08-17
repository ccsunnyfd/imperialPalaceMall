// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.4.1
// - protoc             v3.21.5
// source: mall/v1/goods.proto

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

const OperationGoodsServiceGetGoods = "/api.mall.v1.GoodsService/GetGoods"
const OperationGoodsServiceGetSKUs = "/api.mall.v1.GoodsService/GetSKUs"
const OperationGoodsServiceListGoods = "/api.mall.v1.GoodsService/ListGoods"

type GoodsServiceHTTPServer interface {
	GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsReply, error)
	GetSKUs(context.Context, *GetSKUsRequest) (*GetSKUsReply, error)
	ListGoods(context.Context, *ListGoodsRequest) (*ListGoodsReply, error)
}

func RegisterGoodsServiceHTTPServer(s *http.Server, srv GoodsServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/goods/{id}", _GoodsService_GetGoods0_HTTP_Handler(srv))
	r.GET("/goods", _GoodsService_ListGoods0_HTTP_Handler(srv))
	r.GET("/goods/{id}/sku", _GoodsService_GetSKUs0_HTTP_Handler(srv))
}

func _GoodsService_GetGoods0_HTTP_Handler(srv GoodsServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetGoodsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGoodsServiceGetGoods)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetGoods(ctx, req.(*GetGoodsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetGoodsReply)
		return ctx.Result(200, reply)
	}
}

func _GoodsService_ListGoods0_HTTP_Handler(srv GoodsServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListGoodsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGoodsServiceListGoods)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListGoods(ctx, req.(*ListGoodsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListGoodsReply)
		return ctx.Result(200, reply)
	}
}

func _GoodsService_GetSKUs0_HTTP_Handler(srv GoodsServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSKUsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGoodsServiceGetSKUs)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSKUs(ctx, req.(*GetSKUsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSKUsReply)
		return ctx.Result(200, reply)
	}
}

type GoodsServiceHTTPClient interface {
	GetGoods(ctx context.Context, req *GetGoodsRequest, opts ...http.CallOption) (rsp *GetGoodsReply, err error)
	GetSKUs(ctx context.Context, req *GetSKUsRequest, opts ...http.CallOption) (rsp *GetSKUsReply, err error)
	ListGoods(ctx context.Context, req *ListGoodsRequest, opts ...http.CallOption) (rsp *ListGoodsReply, err error)
}

type GoodsServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewGoodsServiceHTTPClient(client *http.Client) GoodsServiceHTTPClient {
	return &GoodsServiceHTTPClientImpl{client}
}

func (c *GoodsServiceHTTPClientImpl) GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...http.CallOption) (*GetGoodsReply, error) {
	var out GetGoodsReply
	pattern := "/goods/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGoodsServiceGetGoods))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GoodsServiceHTTPClientImpl) GetSKUs(ctx context.Context, in *GetSKUsRequest, opts ...http.CallOption) (*GetSKUsReply, error) {
	var out GetSKUsReply
	pattern := "/goods/{id}/sku"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGoodsServiceGetSKUs))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GoodsServiceHTTPClientImpl) ListGoods(ctx context.Context, in *ListGoodsRequest, opts ...http.CallOption) (*ListGoodsReply, error) {
	var out ListGoodsReply
	pattern := "/goods"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGoodsServiceListGoods))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
