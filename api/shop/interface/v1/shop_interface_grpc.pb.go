// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: shop/interface/v1/shop_interface.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ShopInterfaceClient is the client API for ShopInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopInterfaceClient interface {
	ListCategory(ctx context.Context, in *ListCategoryRequest, opts ...grpc.CallOption) (*ListCategoryReply, error)
	GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsReply, error)
	ListGoods(ctx context.Context, in *ListGoodsRequest, opts ...grpc.CallOption) (*ListGoodsReply, error)
	GetSKUs(ctx context.Context, in *GetSKUsRequest, opts ...grpc.CallOption) (*GetSKUsReply, error)
	WxLogin(ctx context.Context, in *WxLoginRequest, opts ...grpc.CallOption) (*WxLoginReply, error)
	AddCart(ctx context.Context, in *AddCartRequest, opts ...grpc.CallOption) (*AddCartReply, error)
	GetMyCart(ctx context.Context, in *GetMyCartRequest, opts ...grpc.CallOption) (*GetMyCartReply, error)
	UpdateCartNum(ctx context.Context, in *UpdateCartNumRequest, opts ...grpc.CallOption) (*UpdateCartNumReply, error)
}

type shopInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewShopInterfaceClient(cc grpc.ClientConnInterface) ShopInterfaceClient {
	return &shopInterfaceClient{cc}
}

func (c *shopInterfaceClient) ListCategory(ctx context.Context, in *ListCategoryRequest, opts ...grpc.CallOption) (*ListCategoryReply, error) {
	out := new(ListCategoryReply)
	err := c.cc.Invoke(ctx, "/api.shop.interface.v1.ShopInterface/ListCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsReply, error) {
	out := new(GetGoodsReply)
	err := c.cc.Invoke(ctx, "/api.shop.interface.v1.ShopInterface/GetGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) ListGoods(ctx context.Context, in *ListGoodsRequest, opts ...grpc.CallOption) (*ListGoodsReply, error) {
	out := new(ListGoodsReply)
	err := c.cc.Invoke(ctx, "/api.shop.interface.v1.ShopInterface/ListGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) GetSKUs(ctx context.Context, in *GetSKUsRequest, opts ...grpc.CallOption) (*GetSKUsReply, error) {
	out := new(GetSKUsReply)
	err := c.cc.Invoke(ctx, "/api.shop.interface.v1.ShopInterface/GetSKUs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) WxLogin(ctx context.Context, in *WxLoginRequest, opts ...grpc.CallOption) (*WxLoginReply, error) {
	out := new(WxLoginReply)
	err := c.cc.Invoke(ctx, "/api.shop.interface.v1.ShopInterface/WxLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) AddCart(ctx context.Context, in *AddCartRequest, opts ...grpc.CallOption) (*AddCartReply, error) {
	out := new(AddCartReply)
	err := c.cc.Invoke(ctx, "/api.shop.interface.v1.ShopInterface/AddCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) GetMyCart(ctx context.Context, in *GetMyCartRequest, opts ...grpc.CallOption) (*GetMyCartReply, error) {
	out := new(GetMyCartReply)
	err := c.cc.Invoke(ctx, "/api.shop.interface.v1.ShopInterface/GetMyCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) UpdateCartNum(ctx context.Context, in *UpdateCartNumRequest, opts ...grpc.CallOption) (*UpdateCartNumReply, error) {
	out := new(UpdateCartNumReply)
	err := c.cc.Invoke(ctx, "/api.shop.interface.v1.ShopInterface/UpdateCartNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopInterfaceServer is the server API for ShopInterface service.
// All implementations must embed UnimplementedShopInterfaceServer
// for forward compatibility
type ShopInterfaceServer interface {
	ListCategory(context.Context, *ListCategoryRequest) (*ListCategoryReply, error)
	GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsReply, error)
	ListGoods(context.Context, *ListGoodsRequest) (*ListGoodsReply, error)
	GetSKUs(context.Context, *GetSKUsRequest) (*GetSKUsReply, error)
	WxLogin(context.Context, *WxLoginRequest) (*WxLoginReply, error)
	AddCart(context.Context, *AddCartRequest) (*AddCartReply, error)
	GetMyCart(context.Context, *GetMyCartRequest) (*GetMyCartReply, error)
	UpdateCartNum(context.Context, *UpdateCartNumRequest) (*UpdateCartNumReply, error)
	mustEmbedUnimplementedShopInterfaceServer()
}

// UnimplementedShopInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedShopInterfaceServer struct {
}

func (UnimplementedShopInterfaceServer) ListCategory(context.Context, *ListCategoryRequest) (*ListCategoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCategory not implemented")
}
func (UnimplementedShopInterfaceServer) GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoods not implemented")
}
func (UnimplementedShopInterfaceServer) ListGoods(context.Context, *ListGoodsRequest) (*ListGoodsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGoods not implemented")
}
func (UnimplementedShopInterfaceServer) GetSKUs(context.Context, *GetSKUsRequest) (*GetSKUsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSKUs not implemented")
}
func (UnimplementedShopInterfaceServer) WxLogin(context.Context, *WxLoginRequest) (*WxLoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WxLogin not implemented")
}
func (UnimplementedShopInterfaceServer) AddCart(context.Context, *AddCartRequest) (*AddCartReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCart not implemented")
}
func (UnimplementedShopInterfaceServer) GetMyCart(context.Context, *GetMyCartRequest) (*GetMyCartReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyCart not implemented")
}
func (UnimplementedShopInterfaceServer) UpdateCartNum(context.Context, *UpdateCartNumRequest) (*UpdateCartNumReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCartNum not implemented")
}
func (UnimplementedShopInterfaceServer) mustEmbedUnimplementedShopInterfaceServer() {}

// UnsafeShopInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShopInterfaceServer will
// result in compilation errors.
type UnsafeShopInterfaceServer interface {
	mustEmbedUnimplementedShopInterfaceServer()
}

func RegisterShopInterfaceServer(s grpc.ServiceRegistrar, srv ShopInterfaceServer) {
	s.RegisterService(&ShopInterface_ServiceDesc, srv)
}

func _ShopInterface_ListCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).ListCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.shop.interface.v1.ShopInterface/ListCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).ListCategory(ctx, req.(*ListCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_GetGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).GetGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.shop.interface.v1.ShopInterface/GetGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).GetGoods(ctx, req.(*GetGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_ListGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).ListGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.shop.interface.v1.ShopInterface/ListGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).ListGoods(ctx, req.(*ListGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_GetSKUs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSKUsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).GetSKUs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.shop.interface.v1.ShopInterface/GetSKUs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).GetSKUs(ctx, req.(*GetSKUsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_WxLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WxLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).WxLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.shop.interface.v1.ShopInterface/WxLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).WxLogin(ctx, req.(*WxLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_AddCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).AddCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.shop.interface.v1.ShopInterface/AddCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).AddCart(ctx, req.(*AddCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_GetMyCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).GetMyCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.shop.interface.v1.ShopInterface/GetMyCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).GetMyCart(ctx, req.(*GetMyCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_UpdateCartNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCartNumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).UpdateCartNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.shop.interface.v1.ShopInterface/UpdateCartNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).UpdateCartNum(ctx, req.(*UpdateCartNumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShopInterface_ServiceDesc is the grpc.ServiceDesc for ShopInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShopInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.shop.interface.v1.ShopInterface",
	HandlerType: (*ShopInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCategory",
			Handler:    _ShopInterface_ListCategory_Handler,
		},
		{
			MethodName: "GetGoods",
			Handler:    _ShopInterface_GetGoods_Handler,
		},
		{
			MethodName: "ListGoods",
			Handler:    _ShopInterface_ListGoods_Handler,
		},
		{
			MethodName: "GetSKUs",
			Handler:    _ShopInterface_GetSKUs_Handler,
		},
		{
			MethodName: "WxLogin",
			Handler:    _ShopInterface_WxLogin_Handler,
		},
		{
			MethodName: "AddCart",
			Handler:    _ShopInterface_AddCart_Handler,
		},
		{
			MethodName: "GetMyCart",
			Handler:    _ShopInterface_GetMyCart_Handler,
		},
		{
			MethodName: "UpdateCartNum",
			Handler:    _ShopInterface_UpdateCartNum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop/interface/v1/shop_interface.proto",
}
