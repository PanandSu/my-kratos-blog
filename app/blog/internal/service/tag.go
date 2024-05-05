package service

import (
	"context"
	api "my-kratos-blog/api/tag"
	"my-kratos-blog/app/blog/internal/biz"
)

type TagService struct {
	api.UnimplementedTagServer
	uc *biz.TagUseCase
}

func (s *TagService) CreateTag(ctx context.Context, req *api.CreateTagRequest) (*api.CreateTagReply, error) {
	return s.uc.CreateTag(ctx, req), nil
}
func (s *TagService) DeleteTag(ctx context.Context, req *api.DeleteTagRequest) (*api.DeleteTagReply, error) {
	return s.uc.DeleteTag(ctx, req), nil
}
func (s *TagService) ListTag(ctx context.Context, req *api.ListTagRequest) (*api.ListTagReply, error) {
	return s.uc.ListTag(ctx, req), nil
}
