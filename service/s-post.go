package service

import (
	"context"
	pb "post-servic/genproto/post"
)

func (p *PostService) CreatePost(ctx context.Context, in *pb.Post) (*pb.PostResponse, error) {
	return &pb.PostResponse{}, nil
}
func (p *PostService) UpdatePost(ctx context.Context, in *pb.UpdateAPost) (*pb.PostResponse, error) {
	return &pb.PostResponse{}, nil
}
func (p *PostService) GetPostByID(ctx context.Context, in *pb.PostId) (*pb.PostResponse, error) {
	return &pb.PostResponse{}, nil
}
func (p *PostService) ListPosts(ctx context.Context, in *pb.PostList) (*pb.PostListResponse, error) {
	return &pb.PostListResponse{}, nil
}
func (p *PostService) DeletePost(ctx context.Context, in *pb.PostId) (*pb.Message, error) {
	return &pb.Message{}, nil
}
func (p *PostService) AddImageToPost(ctx context.Context, in *pb.ImageUrl) (*pb.PostResponse, error) {
	return &pb.PostResponse{}, nil
}

func (p *PostService) RemoveImageFromPost(ctx context.Context, in *pb.ImageUrl) (*pb.Message, error) {
	return &pb.Message{}, nil
}
func (p *PostService) GetPostByCountry(ctx context.Context, in *pb.PostCountry) (*pb.PostListResponse, error) {
	return &pb.PostListResponse{}, nil
}
