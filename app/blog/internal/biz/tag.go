package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	api "my-kratos-blog/api/tag"
)

type TagRepo interface {
	CreateTag(ctx context.Context, request *api.CreateTagRequest) *api.CreateTagReply
	DeleteTag(ctx context.Context, request *api.DeleteTagRequest) *api.DeleteTagReply
	ListTag(ctx context.Context, request *api.ListTagRequest) *api.ListTagReply
}
type TagUseCase struct {
	repo TagRepo
	log  *log.Helper
}

func NewTagUseCase(repo TagRepo, logger log.Logger) *TagUseCase {
	return &TagUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *TagUseCase) CreateTag(ctx context.Context, request *api.CreateTagRequest) *api.CreateTagReply {
	return uc.repo.CreateTag(ctx, request)
}
func (uc *TagUseCase) DeleteTag(ctx context.Context, request *api.DeleteTagRequest) *api.DeleteTagReply {
	return uc.repo.DeleteTag(ctx, request)
}

func (uc *TagUseCase) ListTag(ctx context.Context, request *api.ListTagRequest) *api.ListTagReply {
	return uc.repo.ListTag(ctx, request)
}
