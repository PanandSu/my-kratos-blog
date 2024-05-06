package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	api "my-kratos-blog/api/tag"
	"my-kratos-blog/app/blog/internal/biz"
)

type tagRepo struct {
	data *Data
	log  *log.Helper
}

func NewTagRepo(data *Data) biz.TagRepo {
	return &tagRepo{
		data: data,
		log:  log.NewHelper(log.With()),
	}
}
func (r *tagRepo) CreateTag(ctx context.Context, request *api.CreateTagRequest) *api.CreateTagReply {}

func (r *tagRepo) DeleteTag(ctx context.Context, request *api.DeleteTagRequest) *api.DeleteTagReply {}

func (r *tagRepo) ListTag(ctx context.Context, request *api.ListTagRequest) *api.ListTagReply {}
