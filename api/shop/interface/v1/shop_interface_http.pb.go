// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.4.1
// - protoc             v3.21.5
// source: shop/interface/v1/shop_interface.proto

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

const OperationShopInterfaceAddCart = "/api.shop.interface.v1.ShopInterface/AddCart"
const OperationShopInterfaceGetAddress = "/api.shop.interface.v1.ShopInterface/GetAddress"
const OperationShopInterfaceGetGoods = "/api.shop.interface.v1.ShopInterface/GetGoods"
const OperationShopInterfaceGetMyCart = "/api.shop.interface.v1.ShopInterface/GetMyCart"
const OperationShopInterfaceGetSKUs = "/api.shop.interface.v1.ShopInterface/GetSKUs"
const OperationShopInterfaceListCategory = "/api.shop.interface.v1.ShopInterface/ListCategory"
const OperationShopInterfaceListGoods = "/api.shop.interface.v1.ShopInterface/ListGoods"
const OperationShopInterfaceRemoveCartItems = "/api.shop.interface.v1.ShopInterface/RemoveCartItems"
const OperationShopInterfaceSaveAddress = "/api.shop.interface.v1.ShopInterface/SaveAddress"
const OperationShopInterfaceUpdateAddress = "/api.shop.interface.v1.ShopInterface/UpdateAddress"
const OperationShopInterfaceUpdateCartNum = "/api.shop.interface.v1.ShopInterface/UpdateCartNum"
const OperationShopInterfaceWxLogin = "/api.shop.interface.v1.ShopInterface/WxLogin"

type ShopInterfaceHTTPServer interface {
	AddCart(context.Context, *AddCartRequest) (*AddCartReply, error)
	GetAddress(context.Context, *GetAddressRequest) (*GetAddressReply, error)
	GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsReply, error)
	GetMyCart(context.Context, *GetMyCartRequest) (*GetMyCartReply, error)
	GetSKUs(context.Context, *GetSKUsRequest) (*GetSKUsReply, error)
	ListCategory(context.Context, *ListCategoryRequest) (*ListCategoryReply, error)
	ListGoods(context.Context, *ListGoodsRequest) (*ListGoodsReply, error)
	RemoveCartItems(context.Context, *RemoveCartItemsRequest) (*RemoveCartItemsReply, error)
	SaveAddress(context.Context, *SaveAddressRequest) (*SaveAddressReply, error)
	UpdateAddress(context.Context, *UpdateAddressRequest) (*UpdateAddressReply, error)
	UpdateCartNum(context.Context, *UpdateCartNumRequest) (*UpdateCartNumReply, error)
	WxLogin(context.Context, *WxLoginRequest) (*WxLoginReply, error)
}

func RegisterShopInterfaceHTTPServer(s *http.Server, srv ShopInterfaceHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/categories", _ShopInterface_ListCategory0_HTTP_Handler(srv))
	r.GET("/v1/goods/{id}", _ShopInterface_GetGoods0_HTTP_Handler(srv))
	r.GET("/v1/goods", _ShopInterface_ListGoods0_HTTP_Handler(srv))
	r.GET("/v1/goods/{id}/sku", _ShopInterface_GetSKUs0_HTTP_Handler(srv))
	r.POST("/v1/wx/login", _ShopInterface_WxLogin0_HTTP_Handler(srv))
	r.PUT("/v1/user/my/cart", _ShopInterface_AddCart0_HTTP_Handler(srv))
	r.GET("/v1/user/my/cart", _ShopInterface_GetMyCart0_HTTP_Handler(srv))
	r.PUT("/v1/user/my/cart/{id}", _ShopInterface_UpdateCartNum0_HTTP_Handler(srv))
	r.DELETE("/v1/user/my/cart/{ids}", _ShopInterface_RemoveCartItems0_HTTP_Handler(srv))
	r.GET("/v1/user/my/address", _ShopInterface_GetAddress0_HTTP_Handler(srv))
	r.POST("/v1/user/my/address", _ShopInterface_SaveAddress0_HTTP_Handler(srv))
	r.PUT("/v1/user/my/address", _ShopInterface_UpdateAddress0_HTTP_Handler(srv))
}

func _ShopInterface_ListCategory0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListCategoryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShopInterfaceListCategory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCategory(ctx, req.(*ListCategoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListCategoryReply)
		return ctx.Result(200, reply)
	}
}

func _ShopInterface_GetGoods0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetGoodsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShopInterfaceGetGoods)
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

func _ShopInterface_ListGoods0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListGoodsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShopInterfaceListGoods)
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

func _ShopInterface_GetSKUs0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSKUsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShopInterfaceGetSKUs)
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

func _ShopInterface_WxLogin0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in WxLoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShopInterfaceWxLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.WxLogin(ctx, req.(*WxLoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*WxLoginReply)
		return ctx.Result(200, reply)
	}
}

func _ShopInterface_AddCart0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddCartRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShopInterfaceAddCart)
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

func _ShopInterface_GetMyCart0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetMyCartRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShopInterfaceGetMyCart)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMyCart(ctx, req.(*GetMyCartRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetMyCartReply)
		return ctx.Result(200, reply)
	}
}

func _ShopInterface_UpdateCartNum0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateCartNumRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShopInterfaceUpdateCartNum)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateCartNum(ctx, req.(*UpdateCartNumRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateCartNumReply)
		return ctx.Result(200, reply)
	}
}

func _ShopInterface_RemoveCartItems0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RemoveCartItemsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShopInterfaceRemoveCartItems)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RemoveCartItems(ctx, req.(*RemoveCartItemsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RemoveCartItemsReply)
		return ctx.Result(200, reply)
	}
}

func _ShopInterface_GetAddress0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetAddressRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShopInterfaceGetAddress)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAddress(ctx, req.(*GetAddressRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetAddressReply)
		return ctx.Result(200, reply)
	}
}

func _ShopInterface_SaveAddress0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SaveAddressRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShopInterfaceSaveAddress)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveAddress(ctx, req.(*SaveAddressRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SaveAddressReply)
		return ctx.Result(200, reply)
	}
}

func _ShopInterface_UpdateAddress0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateAddressRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShopInterfaceUpdateAddress)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateAddress(ctx, req.(*UpdateAddressRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateAddressReply)
		return ctx.Result(200, reply)
	}
}

type ShopInterfaceHTTPClient interface {
	AddCart(ctx context.Context, req *AddCartRequest, opts ...http.CallOption) (rsp *AddCartReply, err error)
	GetAddress(ctx context.Context, req *GetAddressRequest, opts ...http.CallOption) (rsp *GetAddressReply, err error)
	GetGoods(ctx context.Context, req *GetGoodsRequest, opts ...http.CallOption) (rsp *GetGoodsReply, err error)
	GetMyCart(ctx context.Context, req *GetMyCartRequest, opts ...http.CallOption) (rsp *GetMyCartReply, err error)
	GetSKUs(ctx context.Context, req *GetSKUsRequest, opts ...http.CallOption) (rsp *GetSKUsReply, err error)
	ListCategory(ctx context.Context, req *ListCategoryRequest, opts ...http.CallOption) (rsp *ListCategoryReply, err error)
	ListGoods(ctx context.Context, req *ListGoodsRequest, opts ...http.CallOption) (rsp *ListGoodsReply, err error)
	RemoveCartItems(ctx context.Context, req *RemoveCartItemsRequest, opts ...http.CallOption) (rsp *RemoveCartItemsReply, err error)
	SaveAddress(ctx context.Context, req *SaveAddressRequest, opts ...http.CallOption) (rsp *SaveAddressReply, err error)
	UpdateAddress(ctx context.Context, req *UpdateAddressRequest, opts ...http.CallOption) (rsp *UpdateAddressReply, err error)
	UpdateCartNum(ctx context.Context, req *UpdateCartNumRequest, opts ...http.CallOption) (rsp *UpdateCartNumReply, err error)
	WxLogin(ctx context.Context, req *WxLoginRequest, opts ...http.CallOption) (rsp *WxLoginReply, err error)
}

type ShopInterfaceHTTPClientImpl struct {
	cc *http.Client
}

func NewShopInterfaceHTTPClient(client *http.Client) ShopInterfaceHTTPClient {
	return &ShopInterfaceHTTPClientImpl{client}
}

func (c *ShopInterfaceHTTPClientImpl) AddCart(ctx context.Context, in *AddCartRequest, opts ...http.CallOption) (*AddCartReply, error) {
	var out AddCartReply
	pattern := "/v1/user/my/cart"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationShopInterfaceAddCart))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) GetAddress(ctx context.Context, in *GetAddressRequest, opts ...http.CallOption) (*GetAddressReply, error) {
	var out GetAddressReply
	pattern := "/v1/user/my/address"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationShopInterfaceGetAddress))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...http.CallOption) (*GetGoodsReply, error) {
	var out GetGoodsReply
	pattern := "/v1/goods/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationShopInterfaceGetGoods))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) GetMyCart(ctx context.Context, in *GetMyCartRequest, opts ...http.CallOption) (*GetMyCartReply, error) {
	var out GetMyCartReply
	pattern := "/v1/user/my/cart"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationShopInterfaceGetMyCart))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) GetSKUs(ctx context.Context, in *GetSKUsRequest, opts ...http.CallOption) (*GetSKUsReply, error) {
	var out GetSKUsReply
	pattern := "/v1/goods/{id}/sku"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationShopInterfaceGetSKUs))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) ListCategory(ctx context.Context, in *ListCategoryRequest, opts ...http.CallOption) (*ListCategoryReply, error) {
	var out ListCategoryReply
	pattern := "/v1/categories"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationShopInterfaceListCategory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) ListGoods(ctx context.Context, in *ListGoodsRequest, opts ...http.CallOption) (*ListGoodsReply, error) {
	var out ListGoodsReply
	pattern := "/v1/goods"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationShopInterfaceListGoods))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) RemoveCartItems(ctx context.Context, in *RemoveCartItemsRequest, opts ...http.CallOption) (*RemoveCartItemsReply, error) {
	var out RemoveCartItemsReply
	pattern := "/v1/user/my/cart/{ids}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationShopInterfaceRemoveCartItems))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) SaveAddress(ctx context.Context, in *SaveAddressRequest, opts ...http.CallOption) (*SaveAddressReply, error) {
	var out SaveAddressReply
	pattern := "/v1/user/my/address"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationShopInterfaceSaveAddress))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) UpdateAddress(ctx context.Context, in *UpdateAddressRequest, opts ...http.CallOption) (*UpdateAddressReply, error) {
	var out UpdateAddressReply
	pattern := "/v1/user/my/address"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationShopInterfaceUpdateAddress))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) UpdateCartNum(ctx context.Context, in *UpdateCartNumRequest, opts ...http.CallOption) (*UpdateCartNumReply, error) {
	var out UpdateCartNumReply
	pattern := "/v1/user/my/cart/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationShopInterfaceUpdateCartNum))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) WxLogin(ctx context.Context, in *WxLoginRequest, opts ...http.CallOption) (*WxLoginReply, error) {
	var out WxLoginReply
	pattern := "/v1/wx/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationShopInterfaceWxLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
