package service

import (
	"context"

	pb "my-kratos-blog/api/blog"
)

type BlogService struct {
	pb.UnimplementedBlogServer
}

func NewBlogService() *BlogService {
	return &BlogService{}
}

func (s *BlogService) CreateBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.CreateBlogReply, error) {
	return &pb.CreateBlogReply{}, nil
}
func (s *BlogService) UpdateBlog(ctx context.Context, req *pb.UpdateBlogRequest) (*pb.UpdateBlogReply, error) {
	return &pb.UpdateBlogReply{}, nil
}
func (s *BlogService) UpdateIndividualFields(ctx context.Context, req *pb.UpdateIndividualFieldsRequest) (*pb.UpdateIndividualFieldsReply, error) {
	return &pb.UpdateIndividualFieldsReply{}, nil
}
func (s *BlogService) DeleteBlog(ctx context.Context, req *pb.DeleteBlogRequest) (*pb.DeleteBlogReply, error) {
	return &pb.DeleteBlogReply{}, nil
}
func (s *BlogService) GetBlogByTag(ctx context.Context, req *pb.GetBlogRequest) (*pb.GetBlogReply, error) {
	return &pb.GetBlogReply{}, nil
}
func (s *BlogService) ListBlog(ctx context.Context, req *pb.ListBlogRequest) (*pb.ListBlogReply, error) {
	return &pb.ListBlogReply{}, nil
}
func (s *BlogService) GetBlogByID(ctx context.Context, req *pb.GetBlogIDRequest) (*pb.GetBlogIDReply, error) {
	return &pb.GetBlogIDReply{}, nil
}
func (s *BlogService) GetBlogByTitle(ctx context.Context, req *pb.GetBlogByTitleRequest) (*pb.GetBlogByTitleReply, error) {
	return &pb.GetBlogByTitleReply{}, nil
}
func (s *BlogService) UpdateOnly(ctx context.Context, req *pb.UpdateOnlyRequest) (*pb.UpdateOnlyReply, error) {
	return &pb.UpdateOnlyReply{}, nil
}
func (s *BlogService) CacheBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.CreateBlogReply, error) {
	return &pb.CreateBlogReply{}, nil
}
func (s *BlogService) GetCacheBlog(ctx context.Context, req *pb.ListBlogRequest) (*pb.ListCacheReply, error) {
	return &pb.ListCacheReply{}, nil
}
func (s *BlogService) DeleteCacheBlog(ctx context.Context, req *pb.DeleteCacheBlogRequest) (*pb.DeleteCacheBlogReply, error) {
	return &pb.DeleteCacheBlogReply{}, nil
}
func (s *BlogService) AddSuggestBlog(ctx context.Context, req *pb.SuggestBlogRequest) (*pb.SuggestBlogReply, error) {
	return &pb.SuggestBlogReply{}, nil
}
func (s *BlogService) DeleteSuggestBlog(ctx context.Context, req *pb.SuggestBlogRequest) (*pb.SuggestBlogReply, error) {
	return &pb.SuggestBlogReply{}, nil
}
func (s *BlogService) GetAllSuggest(ctx context.Context, req *pb.SearchAllSuggest) (*pb.SearchAllReply, error) {
	return &pb.SearchAllReply{}, nil
}
func (s *BlogService) UpdateBlogVisit(ctx context.Context, req *pb.UpdateBlogVisitRq) (*pb.UpdateBlogVisitRq, error) {
	return &pb.UpdateBlogVisitRq{}, nil
}
func (s *BlogService) GetCommentPower(ctx context.Context, req *pb.GetCommentPowerRq) (*pb.GetCommentPowerReply, error) {
	return &pb.GetCommentPowerReply{}, nil
}
