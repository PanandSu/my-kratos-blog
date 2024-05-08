package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "my-kratos-blog/api/blog"
)

type BlogRepo interface {
	CreateBlog()
	UpdateBlog()
}

type BlogUseCase struct {
	repo BlogRepo
	log  *log.Helper
}

func (uc *BlogUseCase) CreateBlog(ctx context.Context, request *pb.CreateBlogRequest) *pb.CreateBlogReply {
	//TODO implement me
	panic("implement me")
}

func (uc *BlogUseCase) UpdateBlog(ctx context.Context, request *pb.UpdateBlogRequest) *pb.UpdateBlogReply {
	//TODO implement me
	panic("implement me")
}

func (uc *BlogUseCase) UpdateIndividualFields(ctx context.Context, request *pb.UpdateIndividualFieldsRequest) *pb.UpdateIndividualFieldsReply {
	//TODO implement me
	panic("implement me")
}

func (uc *BlogUseCase) DeleteBlog(ctx context.Context, request *pb.DeleteBlogRequest) *pb.DeleteBlogReply {
	//TODO implement me
	panic("implement me")
}

func (uc *BlogUseCase) GetBlogByTag(ctx context.Context, request *pb.GetBlogRequest) *pb.GetBlogReply {
	//TODO implement me
	panic("implement me")
}

func (uc *BlogUseCase) ListBlog(ctx context.Context, request *pb.ListBlogRequest) *pb.ListBlogReply {
	//TODO implement me
	panic("implement me")
}

func (uc *BlogUseCase) GetBlogByID(ctx context.Context, request *pb.GetBlogIDRequest) *pb.GetBlogIDReply {
	//TODO implement me
	panic("implement me")
}

func (uc *BlogUseCase) GetBlogByTitle(ctx context.Context, request *pb.GetBlogByTitleRequest) *pb.GetBlogByTitleReply {
	//TODO implement me
	panic("implement me")
}

func (uc *BlogUseCase) UpdateOnly(ctx context.Context, request *pb.UpdateOnlyRequest) *pb.UpdateOnlyReply {
	//TODO implement me
	panic("implement me")
}

func (uc *BlogUseCase) CacheBlog(ctx context.Context, request *pb.CreateBlogRequest) (*pb.CreateBlogReply {
}

func (uc *BlogUseCase) GetCacheBlog(ctx context.Context, request *pb.ListBlogRequest) (*pb.ListCacheReply {

}

func (uc *BlogUseCase) DeleteCacheBlog(ctx context.Context, request *pb.DeleteCacheBlogRequest) (*pb.DeleteCacheBlogReply {

}

func (uc *BlogUseCase) AddSuggestBlog(ctx context.Context, request *pb.SuggestBlogRequest) (*pb.SuggestBlogReply {

}

func (uc *BlogUseCase) DeleteSuggestBlog(ctx context.Context, request *pb.SuggestBlogRequest) (*pb.SuggestBlogReply {

}

func (uc *BlogUseCase) GetAllSuggest(ctx context.Context, suggest *pb.SearchAllSuggest) (*pb.SearchAllReply {

}

func (uc *BlogUseCase) UpdateBlogVisit(ctx context.Context, rq *pb.UpdateBlogVisitRq) (*pb.UpdateBlogVisitRq {

}

func (uc *BlogUseCase) GetCommentPower(ctx context.Context, rq *pb.GetCommentPowerRq) (*pb.GetCommentPowerReply {
	
}

func (uc *BlogUseCase) mustEmbedUnimplementedBlogServer() {

}
