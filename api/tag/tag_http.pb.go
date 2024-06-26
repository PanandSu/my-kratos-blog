// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v5.26.0
// source: api/tag/tag.proto

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

const OperationTagCreateTag = "/api.Tag/CreateTag"
const OperationTagDeleteTag = "/api.Tag/DeleteTag"
const OperationTagListTag = "/api.Tag/ListTag"

type TagHTTPServer interface {
	CreateTag(context.Context, *CreateTagRequest) (*CreateTagReply, error)
	DeleteTag(context.Context, *DeleteTagRequest) (*DeleteTagReply, error)
	ListTag(context.Context, *ListTagRequest) (*ListTagReply, error)
}

func RegisterTagHTTPServer(s *http.Server, srv TagHTTPServer) {
	r := s.Route("/")
	r.POST("/api/addTag", _Tag_CreateTag0_HTTP_Handler(srv))
	r.DELETE("/api/deleteTag/{id}", _Tag_DeleteTag0_HTTP_Handler(srv))
	r.GET("/api/getAllTag", _Tag_ListTag0_HTTP_Handler(srv))
}

func _Tag_CreateTag0_HTTP_Handler(srv TagHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateTagRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTagCreateTag)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateTag(ctx, req.(*CreateTagRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateTagReply)
		return ctx.Result(200, reply)
	}
}

func _Tag_DeleteTag0_HTTP_Handler(srv TagHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteTagRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTagDeleteTag)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteTag(ctx, req.(*DeleteTagRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteTagReply)
		return ctx.Result(200, reply)
	}
}

func _Tag_ListTag0_HTTP_Handler(srv TagHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTagRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTagListTag)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTag(ctx, req.(*ListTagRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTagReply)
		return ctx.Result(200, reply)
	}
}

type TagHTTPClient interface {
	CreateTag(ctx context.Context, req *CreateTagRequest, opts ...http.CallOption) (rsp *CreateTagReply, err error)
	DeleteTag(ctx context.Context, req *DeleteTagRequest, opts ...http.CallOption) (rsp *DeleteTagReply, err error)
	ListTag(ctx context.Context, req *ListTagRequest, opts ...http.CallOption) (rsp *ListTagReply, err error)
}

type TagHTTPClientImpl struct {
	cc *http.Client
}

func NewTagHTTPClient(client *http.Client) TagHTTPClient {
	return &TagHTTPClientImpl{client}
}

func (c *TagHTTPClientImpl) CreateTag(ctx context.Context, in *CreateTagRequest, opts ...http.CallOption) (*CreateTagReply, error) {
	var out CreateTagReply
	pattern := "/api/addTag"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTagCreateTag))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TagHTTPClientImpl) DeleteTag(ctx context.Context, in *DeleteTagRequest, opts ...http.CallOption) (*DeleteTagReply, error) {
	var out DeleteTagReply
	pattern := "/api/deleteTag/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTagDeleteTag))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *TagHTTPClientImpl) ListTag(ctx context.Context, in *ListTagRequest, opts ...http.CallOption) (*ListTagReply, error) {
	var out ListTagReply
	pattern := "/api/getAllTag"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationTagListTag))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
