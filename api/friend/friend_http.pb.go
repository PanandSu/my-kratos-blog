// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v5.26.0
// source: api/friend/friend.proto

package api

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

const OperationFriendCreateFriend = "/api.Friend/CreateFriend"
const OperationFriendDeleteFriend = "/api.Friend/DeleteFriend"
const OperationFriendGetFriend = "/api.Friend/GetFriend"
const OperationFriendListFriend = "/api.Friend/ListFriend"
const OperationFriendUpdateFriend = "/api.Friend/UpdateFriend"

type FriendHTTPServer interface {
	CreateFriend(context.Context, *CreateFriendRequest) (*CreateFriendReply, error)
	DeleteFriend(context.Context, *DeleteFriendRequest) (*DeleteFriendReply, error)
	GetFriend(context.Context, *GetFriendRequest) (*GetFriendReply, error)
	ListFriend(context.Context, *ListFriendRequest) (*ListFriendReply, error)
	UpdateFriend(context.Context, *UpdateFriendRequest) (*UpdateFriendReply, error)
}

func RegisterFriendHTTPServer(s *http.Server, srv FriendHTTPServer) {
	r := s.Route("/")
	r.POST("/api/addFriend", _Friend_CreateFriend0_HTTP_Handler(srv))
	r.PUT("/api/updateFriend", _Friend_UpdateFriend0_HTTP_Handler(srv))
	r.DELETE("/api/deleteFriend/{id}", _Friend_DeleteFriend0_HTTP_Handler(srv))
	r.GET("/api/getFriend/{id}", _Friend_GetFriend0_HTTP_Handler(srv))
	r.GET("/api/getAllFriend", _Friend_ListFriend0_HTTP_Handler(srv))
}

func _Friend_CreateFriend0_HTTP_Handler(srv FriendHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateFriendRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFriendCreateFriend)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateFriend(ctx, req.(*CreateFriendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateFriendReply)
		return ctx.Result(200, reply)
	}
}

func _Friend_UpdateFriend0_HTTP_Handler(srv FriendHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateFriendRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFriendUpdateFriend)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateFriend(ctx, req.(*UpdateFriendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateFriendReply)
		return ctx.Result(200, reply)
	}
}

func _Friend_DeleteFriend0_HTTP_Handler(srv FriendHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteFriendRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFriendDeleteFriend)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteFriend(ctx, req.(*DeleteFriendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteFriendReply)
		return ctx.Result(200, reply)
	}
}

func _Friend_GetFriend0_HTTP_Handler(srv FriendHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetFriendRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFriendGetFriend)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetFriend(ctx, req.(*GetFriendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetFriendReply)
		return ctx.Result(200, reply)
	}
}

func _Friend_ListFriend0_HTTP_Handler(srv FriendHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListFriendRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFriendListFriend)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListFriend(ctx, req.(*ListFriendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListFriendReply)
		return ctx.Result(200, reply)
	}
}

type FriendHTTPClient interface {
	CreateFriend(ctx context.Context, req *CreateFriendRequest, opts ...http.CallOption) (rsp *CreateFriendReply, err error)
	DeleteFriend(ctx context.Context, req *DeleteFriendRequest, opts ...http.CallOption) (rsp *DeleteFriendReply, err error)
	GetFriend(ctx context.Context, req *GetFriendRequest, opts ...http.CallOption) (rsp *GetFriendReply, err error)
	ListFriend(ctx context.Context, req *ListFriendRequest, opts ...http.CallOption) (rsp *ListFriendReply, err error)
	UpdateFriend(ctx context.Context, req *UpdateFriendRequest, opts ...http.CallOption) (rsp *UpdateFriendReply, err error)
}

type FriendHTTPClientImpl struct {
	cc *http.Client
}

func NewFriendHTTPClient(client *http.Client) FriendHTTPClient {
	return &FriendHTTPClientImpl{client}
}

func (c *FriendHTTPClientImpl) CreateFriend(ctx context.Context, in *CreateFriendRequest, opts ...http.CallOption) (*CreateFriendReply, error) {
	var out CreateFriendReply
	pattern := "/api/addFriend"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFriendCreateFriend))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FriendHTTPClientImpl) DeleteFriend(ctx context.Context, in *DeleteFriendRequest, opts ...http.CallOption) (*DeleteFriendReply, error) {
	var out DeleteFriendReply
	pattern := "/api/deleteFriend/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationFriendDeleteFriend))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FriendHTTPClientImpl) GetFriend(ctx context.Context, in *GetFriendRequest, opts ...http.CallOption) (*GetFriendReply, error) {
	var out GetFriendReply
	pattern := "/api/getFriend/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationFriendGetFriend))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FriendHTTPClientImpl) ListFriend(ctx context.Context, in *ListFriendRequest, opts ...http.CallOption) (*ListFriendReply, error) {
	var out ListFriendReply
	pattern := "/api/getAllFriend"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationFriendListFriend))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FriendHTTPClientImpl) UpdateFriend(ctx context.Context, in *UpdateFriendRequest, opts ...http.CallOption) (*UpdateFriendReply, error) {
	var out UpdateFriendReply
	pattern := "/api/updateFriend"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFriendUpdateFriend))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
